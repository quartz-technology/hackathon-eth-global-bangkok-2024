import { useMemo, useState, type FC } from "react";

import EulerIcon from "../assets/euler.svg?react";

import { Page } from "../components/Page";
import { Address } from "../components/Address";
import { useMarkets } from "../hooks/useMarkets";
import {
	Button,
	Headline,
	Subheadline,
	Text,
} from "@telegram-apps/telegram-ui";
import { formatNumber } from "../utils";
import { useClient } from "../components/WalletProvider";
import { erc20Abi, erc4626Abi, formatUnits } from "viem";
import { useAddress } from "../hooks/useAddress";
import { useBalance } from "../hooks/useBalance";
import { mainnet } from "viem/chains";
import { getTransactionReceipt } from "viem/actions";
import { useParams } from "react-router-dom";

const USDC_ADDRESS = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";

export const DepositPage: FC = () => {
	const { client } = useClient();
	const { address } = useAddress();
	const { markets = [] } = useMarkets(address);
	const { balance } = useBalance(address);

	const { market } = useParams<{ market: string }>();

	const formattedBalance = formatNumber(formatUnits(balance ?? 0n, 6));

	console.log(market);

	const [txState, setTxState] = useState<
		"initial" | "signature" | "mining" | "success" | "error"
	>("initial");

	const selectedMarket = useMemo(
		() => markets.find((m) => m.address === market),
		[market, markets],
	);

	const onEarn = async () => {
		if (!client || !address || !balance || !selectedMarket) return;

		setTxState("signature");

		const txHashApprove = await client.writeContract({
			abi: erc20Abi,
			account: address,
			address: USDC_ADDRESS,
			functionName: "approve",
			args: [selectedMarket.address, balance],
			chain: mainnet,
		});

		setTxState("mining");

		const receiptApprove = await getTransactionReceipt(client, {
			hash: txHashApprove,
		});

		if (receiptApprove.status === "reverted") {
			setTxState("error");
			return;
		}

		const txHash = await client.writeContract({
			abi: erc4626Abi,
			account: address,
			address: selectedMarket.address,
			functionName: "deposit",
			args: [balance, address],
			chain: mainnet,
		});

		const receipt = await getTransactionReceipt(client, { hash: txHash });

		if (receipt.status === "success") {
			setTxState("success");
		} else {
			setTxState("error");
		}
	};

	return (
		<Page back={true}>
			<Address />

			<div
				style={{
					display: "flex",
					flexDirection: "column",
					gap: 4,
					justifyContent: "center",
					alignItems: "center",
					marginTop: 46,
					marginBottom: 32,
				}}
			>
				<EulerIcon
					width={96}
					height={96}
					style={{
						borderRadius: 12,
						marginBottom: 24,
						backgroundColor: "#0C2129",
					}}
				/>

				<Headline
					style={{ color: "var(--tg-theme-text-color)" }}
					weight="1"
					Component={"h1"}
				>
					{formatNumber((selectedMarket?.apy ?? 0) * 100)}%
				</Headline>
				<Subheadline level="2">
					Earn the best rewards with Euler Vault
				</Subheadline>
			</div>

			<div
				style={{
					display: "flex",
					gap: 10,
					justifyContent: "center",
					alignItems: "center",
					margin: 20,
					textAlign: "center",
				}}
			>
				<div
					style={{
						width: 10,
						height: 10,
						borderRadius: 1000,
						backgroundColor:
							txState === "success"
								? "var(--tgui--green)"
								: txState === "error"
									? "var(--tgui--destructive_text_color)"
									: txState === "signature"
										? "var(--tgui--link_color)"
										: txState === "mining"
											? "var(--tgui--link_color)"
											: "",
					}}
				/>
				{txState === "signature" && <Text>Signing transaction...</Text>}
				{txState === "mining" && <Text>Transaction is being mined...</Text>}
				{txState === "success" && <Text>Transaction successful!</Text>}
				{txState === "error" && <Text>Transaction failed!</Text>}
			</div>
			<div style={{ padding: "0px 18px" }}>
				<Button
					stretched
					mode="filled"
					size="l"
					style={{ borderRadius: 16, backgroundColor: "#2990FF" }}
					onClick={() => onEarn()}
				>
					Deposit {formattedBalance} USDC
				</Button>
			</div>
		</Page>
	);
};

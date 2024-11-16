import { useState, type FC } from "react";

import MagicIcon from "../assets/magic.svg?react";

import { Page } from "../components/Page";
import { Address } from "../components/Address";
import {
	Button,
	Headline,
	Subheadline,
	Text,
} from "@telegram-apps/telegram-ui";
import { useOperator } from "../hooks/useOperator";
import { useClient } from "../components/WalletProvider";
import { parseAbi } from "viem";
import { useAddress } from "../hooks/useAddress";
import { useBalance } from "../hooks/useBalance";
import { mainnet } from "viem/chains";
import { getTransactionReceipt } from "viem/actions";
import { EVC_ADDRESS, OPERATOR_ADDRESS } from "../utils";
import { apiUrl } from "../config";

export const MagicRebalance: FC = () => {
	const { client } = useClient();
	const { address } = useAddress();
	const { balance } = useBalance(address);
	const { isAuthorized, refetch } = useOperator(address);

	const [txState, setTxState] = useState<
		"initial" | "signature" | "mining" | "success" | "error"
	>("initial");

	const onMagicRebalance = async () => {
		if (
			!client ||
			!address ||
			balance === undefined ||
			isAuthorized === undefined
		)
			return;

		setTxState("signature");

		const txHash = await client.writeContract({
			abi: parseAbi([
				"function setAccountOperator(address account, address operator, bool authorized) external",
			]),
			account: address,
			address: EVC_ADDRESS,
			functionName: "setAccountOperator",
			args: [address, OPERATOR_ADDRESS, !isAuthorized],
			chain: mainnet,
		});

		setTxState("mining");

		const receipt = await getTransactionReceipt(client, {
			hash: txHash,
		});

		if (receipt.status === "success") {
			setTxState("success");
		} else {
			setTxState("error");
		}

		await fetch(`${apiUrl}/rebalancing?wallet=${address}&run=${!isAuthorized}`);

		await refetch();

		setTimeout(() => setTxState("initial"), 2500);
	};

	return (
		<Page back={true}>
			<Address />

			<div
				style={{
					display: "flex",
					flexDirection: "column",
					padding: "0px 32px",
					textAlign: "center",
					gap: 10,
					justifyContent: "center",
					alignItems: "center",
					marginTop: 46,
					marginBottom: 32,
				}}
			>
				<MagicIcon
					width={76}
					height={76}
					style={{
						borderRadius: 12,
						marginBottom: 20,
						padding: 10,
						backgroundColor: "#2990FF",
					}}
				/>

				<Headline
					style={{ color: "var(--tg-theme-text-color)" }}
					weight="1"
					Component={"h1"}
				>
					Magic Euler
				</Headline>
				<Subheadline level="2">
					Assign your Euler position to an Operator, which will automatically
					rebalance your Euler Vaults to ensure you receive the optimal rewards.
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
					mode="filled"
					stretched
					size="l"
					disabled={isAuthorized === undefined}
					style={{
						borderRadius: 16,
						backgroundColor: isAuthorized
							? "var(--tgui--destructive_text_color)"
							: "#2990FF",
					}}
					onClick={() => onMagicRebalance()}
				>
					{isAuthorized ? "Disable" : "Enable"}
				</Button>
			</div>
		</Page>
	);
};

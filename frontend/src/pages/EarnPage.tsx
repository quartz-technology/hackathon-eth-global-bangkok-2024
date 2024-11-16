import { useMemo, type FC } from "react";

import { Page } from "../components/Page";
import { Address } from "../components/Address";
import { useMarkets } from "../hooks/useMarkets";
import {
	Cell,
	LargeTitle,
	List,
	Skeleton,
	Text,
} from "@telegram-apps/telegram-ui";
import { formatAddress, formatNumber } from "../utils";
import { formatUnits } from "viem";
import { useAddress } from "../hooks/useAddress";
import { useBalance } from "../hooks/useBalance";
import { Link } from "react-router-dom";

const AvailableBalance = () => {
	const { address } = useAddress();
	const {
		// for skeleton
		balance = 100000000000n,
		isLoading,
	} = useBalance(address);

	const formattedBalance = formatNumber(formatUnits(balance, 6));

	return (
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
			<Text>Available Balance</Text>
			<Skeleton visible={isLoading} className="rounded-skeleton">
				<LargeTitle
					style={{ color: "var(--tg-theme-text-color)" }}
					weight="1"
					Component={"h1"}
				>
					<span style={{ color: "var(--tg-theme-hint-color)" }}>$</span>
					{formattedBalance}
				</LargeTitle>
			</Skeleton>
		</div>
	);
};

export const EarnPage: FC = () => {
	const { address } = useAddress();
	const {
		markets = [
			{ address: "0x001", apy: 0, balance: 0n, symbol: "eUSDC" },
			{ address: "0x002", apy: 0, balance: 0n, symbol: "eUSDC" },
			{ address: "0x003", apy: 0, balance: 0n, symbol: "eUSDC" },
			{ address: "0x004", apy: 0, balance: 0n, symbol: "eUSDC" },
		],
		isLoading,
	} = useMarkets(address);

	const sortedMarkets = useMemo(
		() => markets.sort((a, b) => b.apy - a.apy),
		[markets],
	);

	return (
		<Page back={true}>
			<Address />

			<AvailableBalance />

			<List
				style={{
					maxHeight: "calc(100vh - 220px)",
					overflowY: "auto",
				}}
				className="hidden-scrollbar"
			>
				{sortedMarkets.map((market, idx) => (
					<Skeleton
						key={market.address}
						visible={isLoading}
						className="rounded-skeleton"
					>
						<Link
							to={`/earn/${market.address}`}
							style={{
								textDecoration: "none",
								color: "var(--tg-theme-text-primary)",
								marginBottom: 12,
								display: "block",
							}}
						>
							<Cell
								style={{
									backgroundColor: "var(--tgui--secondary_fill)",
									borderRadius: 16,
								}}
								interactiveAnimation="background"
								before={
									idx < 3 ? (
										<Text weight="1">{idx + 1}.</Text>
									) : (
										<div style={{ visibility: "hidden" }}>0.</div>
									)
								}
								subtitle={market.symbol}
								after={
									<Text weight="1">{formatNumber(market.apy * 100)}%</Text>
								}
							>
								<Text weight="2">{formatAddress(market.address)}</Text>
							</Cell>
						</Link>
					</Skeleton>
				))}
			</List>
			<div
				style={{
					position: "relative",
				}}
			>
				<div
					style={{
						position: "absolute",
						width: "100%",
						height: 40,
						bottom: 0,
						background:
							"linear-gradient(180deg, rgba(255, 255, 255, 0) 0%, var(--tg-theme-secondary-bg-color) 100%)",
					}}
				/>
			</div>
		</Page>
	);
};

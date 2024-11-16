import {
	Cell,
	List,
	Text,
	Skeleton,
	LargeTitle,
} from "@telegram-apps/telegram-ui";
import { useMemo, type FC } from "react";

import UsdcIcon from "../assets/usdc.svg?react";
import MagicIcon from "../assets/magic.svg?react";
import EulerIcon from "../assets/euler.svg?react";

import { Page } from "../components/Page";
import { Address } from "../components/Address";
import { formatNumber } from "../utils";
import { useAddress } from "../hooks/useAddress";
import { useOperator } from "../hooks/useOperator";
import { useBalance } from "../hooks/useBalance";
import { useMarkets } from "../hooks/useMarkets";
import { formatUnits } from "viem";
import { Link } from "react-router-dom";

const TotalBalance = ({
	totalBalance,
}: { totalBalance: bigint | undefined }) => {
	const formattedBalance = formatNumber(
		formatUnits(totalBalance ?? 100000000n, 6),
	);

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
			<Text>Total Balance</Text>
			<Skeleton
				visible={totalBalance === undefined}
				className="rounded-skeleton"
			>
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

const AvailableBalance = () => {
	const { address } = useAddress();
	const {
		// for skeleton
		balance = 100000000000n,
		isLoading,
	} = useBalance(address);

	const formattedBalance = formatNumber(formatUnits(balance, 6));

	return (
		<Cell
			style={{
				backgroundColor: "var(--tg-theme-bg-color)",
				borderRadius: 16,
			}}
			interactiveAnimation="background"
			before={<UsdcIcon width={40} height={40} />}
			subtitle="Available Balance"
			after={
				<Skeleton visible={isLoading} className="rounded-skeleton">
					<Text>${formattedBalance}</Text>
				</Skeleton>
			}
		>
			<Text weight="2">USDC</Text>
		</Cell>
	);
};

export const IndexPage: FC = () => {
	const { address } = useAddress();
	const { markets } = useMarkets(address);

	const { isAuthorized } = useOperator(address);

	const totalBalance = useMemo(() => {
		return markets?.reduce((acc, market) => {
			return acc + BigInt(market.balance);
		}, 0n);
	}, [markets]);

	return (
		<Page back={false}>
			<Address />

			<TotalBalance totalBalance={totalBalance} />

			<List>
				<AvailableBalance />

				<Link
					to="/earn"
					style={{
						textDecoration: "none",
						color: "var(--tg-theme-text-primary)",
						marginBottom: 12,
						display: "block",
					}}
				>
					<Cell
						style={{
							backgroundColor: "var(--tg-theme-bg-color)",
							borderRadius: 16,
						}}
						interactiveAnimation="background"
						before={<EulerIcon width={40} height={40} />}
						subtitle="Start earning rewards on your USDC"
					>
						<Text weight="2">Earn with Euler</Text>
					</Cell>
				</Link>

				{!isAuthorized && (
					<Link
						to="/magic-rebalance"
						style={{
							textDecoration: "none",
							display: "block",
							color: "var(--tg-theme-text-primary)",
						}}
					>
						<Cell
							style={{
								backgroundColor: "var(--tg-theme-bg-color)",
								borderRadius: 16,
							}}
							interactiveAnimation="background"
							before={
								<MagicIcon
									width={26}
									height={26}
									style={{
										borderRadius: 8,
										background: "#2990FF",
										padding: 6,
									}}
								/>
							}
							subtitle="Automatically maximize your APY"
						>
							<Text weight="2">Enable Magic Euler</Text>
						</Cell>
					</Link>
				)}
			</List>
		</Page>
	);
};

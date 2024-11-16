import {
	Cell,
	List,
	Text,
	Skeleton,
	LargeTitle,
} from "@telegram-apps/telegram-ui";
import type { FC } from "react";

import UsdcIcon from "../assets/usdc.svg?react";
import EulerIcon from "../assets/euler.svg?react";

import { Page } from "../components/Page";
import { Address } from "../components/Address";
import { formatNumber } from "../utils";
import { useAddress } from "../hooks/useAddress";
import { useBalance } from "../hooks/useBalance";
import { formatUnits } from "viem";

const TotalBalance = () => {
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
			<Text>Total Balance</Text>
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
			USDC
		</Cell>
	);
};

export const IndexPage: FC = () => {
	return (
		<Page back={false}>
			<Address />

			<TotalBalance />

			<List>
				<AvailableBalance />

				<Cell
					style={{
						backgroundColor: "var(--tgui--secondary_fill)",
						borderRadius: 16,
					}}
					interactiveAnimation="background"
					before={<EulerIcon width={40} height={40} />}
					subtitle="Start earning rewards on your USDC"
				>
					Earn with Euler
				</Cell>
			</List>
		</Page>
	);
};

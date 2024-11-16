import { useAddress } from "../hooks/useAddress";
import { useHistory } from "../hooks/useHistory";
import { formatAddress, formatNumber } from "../utils";
import { Card, Cell, List, Text } from "@telegram-apps/telegram-ui";
import { useState } from "react";
import { formatUnits } from "viem";
import ChevronIcon from "../assets/chevron.svg?react";

export const History = () => {
	const { address } = useAddress();
	const { history } = useHistory(address);
	const [showTxHistory, setShowTxHistory] = useState(false);

	if (!history) return null;

	return (
		<Card
			style={{
				position: "absolute",
				bottom: 0,
				left: 0,
				margin: "0px 0px",
				borderRadius: "0x",
				width: "100%",
			}}
		>
			<Cell
				onClick={() => setShowTxHistory(!showTxHistory)}
				after={
					<ChevronIcon
						style={{
							transform: showTxHistory ? "rotate(180deg)" : "",
							marginLeft: "auto",
							width: 18,
							height: 18,
						}}
					/>
				}
			>
				<Text style={{ padding: "8px 0px 0px 8px" }}>Transaction history</Text>
			</Cell>

			<List
				style={{
					maxHeight: showTxHistory ? "500px" : "20px",
					overflowY: "auto",
					transition: "max-height 0.2s ease-in-out",
				}}
				className="hidden-scrollbar"
			>
				{history?.map((hist) => (
					<Cell
						key={hist.id}
						style={{
							backgroundColor: "var(--tgui--secondary_fill)",
							borderRadius: 16,
							pointerEvents: "none",
						}}
						subtitle={<>from: {formatAddress(hist.from)}</>}
						interactiveAnimation="background"
						after={
							<Text>
								${formatNumber(formatUnits(BigInt(hist.amount), 6), true)}
							</Text>
						}
					>
						to: {formatAddress(hist.to)}
					</Cell>
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
		</Card>
	);
};

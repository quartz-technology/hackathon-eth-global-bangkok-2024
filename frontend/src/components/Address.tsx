import { useAddress } from "../hooks/useAddress";
import { Cell, Skeleton, Text } from "@telegram-apps/telegram-ui";
import EthIcon from "../assets/eth.svg?react";
import CopyIcon from "../assets/copy.svg?react";
import { formatAddress } from "../utils";

export const Address = () => {
	const { address = "", isLoading } = useAddress();

	return (
		<div style={{ margin: "6px 18px" }}>
			<Skeleton visible={isLoading} className="rounded-skeleton">
				<Cell
					before={<EthIcon width={24} height={24} />}
					style={{
						display: "flex",
						alignItems: "center",
						backgroundColor: "var(--tg-theme-bg-color)",
						gap: 10,
						borderRadius: 16,
					}}
					after={
						<CopyIcon style={{ marginLeft: "auto" }} width={22} height={22} />
					}
					interactiveAnimation="background"
				>
					<Text weight="2">{formatAddress(address)}</Text>
				</Cell>
			</Skeleton>
		</div>
	);
};

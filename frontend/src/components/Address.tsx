import { useAddress } from "../hooks/useAddress";
import { Cell, Skeleton, Text } from "@telegram-apps/telegram-ui";
import EthIcon from "../assets/eth.svg?react";
import CopyIcon from "../assets/copy.svg?react";

export const Address = () => {
	const { address = "0x0000...", isLoading } = useAddress();

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
					<Text weight="2" style={{ textOverflow: "ellipsis" }}>
						{address}
					</Text>
				</Cell>
			</Skeleton>
		</div>
	);
};

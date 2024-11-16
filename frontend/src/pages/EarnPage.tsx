import type { FC } from "react";

import { Page } from "../components/Page";
import { Address } from "../components/Address";

export const EarnPage: FC = () => {
	return (
		<Page back={false}>
			<Address />

			{/* <TotalBalance />

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
			</List> */}
		</Page>
	);
};

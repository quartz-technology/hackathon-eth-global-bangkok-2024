import type { ComponentType, JSX } from "react";

import { IndexPage } from "../pages/IndexPage";
import { EarnPage } from "../pages/EarnPage";
import { DepositPage } from "../pages/DepositPage";
import { MagicRebalance } from "../pages/MagicRebalance";

interface Route {
	path: string;
	Component: ComponentType;
	title?: string;
	icon?: JSX.Element;
}

export const routes: Route[] = [
	{ path: "/", Component: IndexPage },
	{ path: "/earn", Component: EarnPage },
	{ path: "/earn/:market", Component: DepositPage },
	{ path: "/magic-rebalance", Component: MagicRebalance },
];

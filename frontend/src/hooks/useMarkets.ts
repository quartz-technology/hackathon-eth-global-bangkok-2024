import { skipToken, useQuery } from "@tanstack/react-query";
import { apiUrl } from "../config";

type Market = {
	address: `0x${string}`;
	balance: string;
	symbol: string;
	apy: number;
};

export const useMarkets = (walletAddress: `0x${string}` | undefined) => {
	const {
		data: markets,
		isLoading,
		isPending,
	} = useQuery({
		queryKey: ["getMarkets", walletAddress],
		queryFn: walletAddress
			? async () => {
					const req = await fetch(`${apiUrl}/markets?wallet=${walletAddress}`);
					const data = (await req.json()) as Market[];

					return data;
				}
			: skipToken,
	});

	return { markets, isLoading: isLoading || isPending };
};

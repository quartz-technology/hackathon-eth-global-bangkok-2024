import { subgraphApiUrl } from "../config";
import { skipToken, useQuery } from "@tanstack/react-query";

export const useHistory = (walletAddress: `0x${string}` | undefined) => {
	const { data: history, isLoading } = useQuery({
		queryKey: ["getHistory", walletAddress],
		queryFn: walletAddress
			? async () => {
					const req = await fetch(subgraphApiUrl, {
						method: "POST",
						headers: {
							"Content-Type": "application/json",
						},
						body: JSON.stringify({
							query: `query GetRebalances($wallet: Bytes!) {
					    rebalances(where: {
					      onBehalfOfAccount: $wallet
					    }) {
					        id
					        from
					        to
					        onBehalfOfAccount
					        amount
					        transactionHash
					        blockTimestamp
					        blockNumber
					    }
					}`,
							variables: {
								wallet: walletAddress,
							},
						}),
					});

					const { data } = (await req.json()) as {
						data: {
							rebalances: {
								id: `0x${string}`;
								from: `0x${string}`;
								to: `0x${string}`;
								onBehalfOfAccount: `0x${string}`;
								amount: string;
								transactionHash: `0x${string}`;
								blockTimestamp: string;
								blockNumber: string;
							}[];
						};
					};

					return data.rebalances;
				}
			: skipToken,
	});

	return { history, isLoading };
};

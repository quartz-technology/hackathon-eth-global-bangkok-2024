import { skipToken, useQuery } from "@tanstack/react-query";
import { edenFetch } from "@elysiajs/eden";
import type { Server } from "../server";
import { apiUrl } from "../config";

export const useBalance = (walletAddress: `0x${string}` | undefined) => {
	const {
		data: balance,
		isLoading,
		isPending,
	} = useQuery({
		queryKey: ["getBalance", walletAddress],
		queryFn: walletAddress
			? async () => {
					const fetch = edenFetch<Server>(apiUrl);

					const data = await fetch("/balance", {
						method: "POST",
						body: {
							walletAddress: walletAddress,
						},
					});

					return BigInt(data.data?.balance ?? 0);
				}
			: skipToken,
	});

	return { balance, isLoading: isLoading || isPending };
};

import { web3AuthApiUrl } from "../config";
import type { Server } from "../server";
import { edenFetch } from "@elysiajs/eden";
import { skipToken, useQuery } from "@tanstack/react-query";

export const useHistory = (walletAddress: `0x${string}` | undefined) => {
	const { data: history, isLoading } = useQuery({
		queryKey: ["getHistory", walletAddress],
		queryFn: walletAddress
			? async () => {
					const fetch = edenFetch<Server>(web3AuthApiUrl);

					const data = await fetch("/history", {
						method: "POST",
						body: {
							walletAddress: walletAddress,
						},
					});

					return data.data;
				}
			: skipToken,
	});

	return { history, isLoading };
};

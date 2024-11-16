import { useClient } from "../components/WalletProvider";
import { skipToken, useQuery } from "@tanstack/react-query";

export const useAddress = () => {
	const { client } = useClient();

	const {
		data: address,
		isLoading,
		isPending,
	} = useQuery({
		queryKey: ["getAccount", client],
		queryFn: client
			? async () => {
					return (await client.getAddresses())?.[0];
				}
			: skipToken,
	});

	return { address, isLoading: isLoading || isPending };
};

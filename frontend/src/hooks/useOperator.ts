import { readContract } from "viem/actions";
import { useClient } from "../components/WalletProvider";
import { skipToken, useQuery } from "@tanstack/react-query";
import { parseAbi } from "viem";
import { EVC_ADDRESS, OPERATOR_ADDRESS } from "../utils";

export const useOperator = (walletAddress: `0x${string}` | undefined) => {
	const { client } = useClient();

	const { data: isAuthorized, isLoading } = useQuery({
		queryKey: ["getOperator", walletAddress, client],
		queryFn:
			walletAddress && client
				? async () => {
						const isAuthorized = await readContract(client, {
							abi: parseAbi([
								"function isAccountOperatorAuthorized(address account, address operator) external view returns (bool)",
							]),
							address: EVC_ADDRESS,
							functionName: "isAccountOperatorAuthorized",
							args: [walletAddress, OPERATOR_ADDRESS],
						});

						return isAuthorized;
					}
				: skipToken,
	});

	return { isAuthorized, isLoading };
};

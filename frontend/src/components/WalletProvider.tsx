import { useLaunchParams } from "@telegram-apps/sdk-react";
import { EthereumPrivateKeyProvider } from "@web3auth/ethereum-provider";
import {
	createContext,
	type PropsWithChildren,
	useContext,
	useEffect,
	useState,
} from "react";
import { createClient, custom, walletActions, type WalletClient } from "viem";
import { CHAIN_NAMESPACES, WEB3AUTH_NETWORK } from "@web3auth/base";
import { decodeToken, Web3Auth } from "@web3auth/single-factor-auth";
import { apiUrl, clientId, rpcUrl, verifier } from "../config";
import { edenFetch } from "@elysiajs/eden";
import type { JwtPayload, Server } from "../server";

const WalletContext = createContext<{ client: WalletClient | undefined }>({
	client: undefined,
});

export const useClient = () => useContext(WalletContext);

export const WalletProvider = ({ children }: PropsWithChildren) => {
	const [client, setClient] = useState<WalletClient>();
	const { initDataRaw } = useLaunchParams();

	useEffect(() => {
		(async () => {
			try {
				const privateKeyProvider = new EthereumPrivateKeyProvider({
					config: {
						chainConfig: {
							chainNamespace: CHAIN_NAMESPACES.EIP155,
							rpcTarget: rpcUrl,
							blockExplorerUrl: "https://etherscan.io",
							logo: "https://cryptologos.cc/logos/ethereum-eth-logo.png",
							chainId: "0x1",
							ticker: "ETH",
							tickerName: "Ethereum",
						},
					},
				});

				const web3authInstance = new Web3Auth({
					clientId: clientId,
					web3AuthNetwork: WEB3AUTH_NETWORK.SAPPHIRE_DEVNET,
					usePnPKey: false,
					privateKeyProvider,
				});

				console.log("Initializing Web3Auth...");

				await web3authInstance.init();

				console.log("Web3Auth initialized.");

				if (web3authInstance.status === "connected") {
					await web3authInstance.logout();
				}

				const fetch = edenFetch<Server>(apiUrl);

				if (!initDataRaw) {
					console.error("No initDataRaw received.");
					return;
				}

				const data = await fetch("/auth/telegram", {
					method: "POST",
					body: {
						initDataRaw: initDataRaw,
					},
				});

				const token = data.data?.token;

				if (!token) {
					console.error("No token received from the server.");
					return;
				}

				const { payload } = decodeToken<JwtPayload>(token);

				const provider = await web3authInstance.connect({
					verifier: verifier,
					verifierId: payload.sub,
					idToken: token,
				});

				if (!provider) {
					console.error("No provider received from Web3Auth.");
					return;
				}

				console.log("Connected to Web3Auth.");

				const client = createClient({
					transport: custom({
						async request({ method, params }) {
							const response = await provider.request({ method, params });
							return response;
						},
					}),
				}).extend(walletActions);

				setClient(client);

				console.log("Viem client created.");
			} catch (error) {
				console.error(error);
			}
		})();
	}, [initDataRaw]);

	return (
		<WalletContext.Provider value={{ client }}>
			{children}
		</WalletContext.Provider>
	);
};

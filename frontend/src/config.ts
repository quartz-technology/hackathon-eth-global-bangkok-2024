const web3AuthVerifier = import.meta.env.VITE_WEB3_AUTH_VERIFIER;
const web3AuthClientId = import.meta.env.VITE_WEB3_AUTH_CLIENTID;
const web3AuthApiUrl = import.meta.env.VITE_WEB3_AUTH_API_URL;
const rpcUrl = import.meta.env.VITE_ETH_RPC_URL;

if (!web3AuthVerifier) throw new Error("WEB3_AUTH_VERIFIER is required");
if (!web3AuthClientId) throw new Error("WEB3_AUTH_CLIENTID is required");
if (!web3AuthApiUrl) throw new Error("WEB3_AUTH_API_URL is required");
if (!rpcUrl) throw new Error("ETH_RPC_URL is required");

export { web3AuthVerifier, web3AuthClientId, web3AuthApiUrl, rpcUrl };

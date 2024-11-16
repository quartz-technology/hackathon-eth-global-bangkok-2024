const verifier = import.meta.env.VITE_WEB3_AUTH_VERIFIER;
const clientId = import.meta.env.VITE_WEB3_AUTH_CLIENTID;
const apiUrl = import.meta.env.VITE_API_URL;
const rpcUrl = import.meta.env.VITE_ETH_RPC_URL;

if (!verifier) throw new Error("WEB3_AUTH_VERIFIER is required");
if (!clientId) throw new Error("WEB3_AUTH_CLIENTID is required");
if (!apiUrl) throw new Error("API_URL is required");
if (!rpcUrl) throw new Error("ETH_RPC_URL is required");

export { verifier, clientId, apiUrl, rpcUrl };

import { Elysia, t } from "elysia";
import { cors } from "@elysiajs/cors";
import { sign } from "jsonwebtoken";
import fs from "node:fs";
import path from "node:path";
import crypto from "node:crypto";

import { validate, type User } from "@telegram-apps/init-data-node";
import { createClient, erc20Abi, http } from "viem";
import { readContract, getBalance } from "viem/actions";
import { mainnet } from "viem/chains";

const {
	TELEGRAM_BOT_TOKEN,
	APP_URL,
	NODE_ENV,
	JWT_KEY_ID,
	ONEINCH_API_KEY,
	ETH_RPC_URL,
	PRIVATE_KEY,
} = Bun.env;

console.log(PRIVATE_KEY);

const USDC_ADDRESS = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";

if (!PRIVATE_KEY) throw new Error("PRIVATE_KEY not found");
if (!TELEGRAM_BOT_TOKEN) throw new Error("TELEGRAM_BOT_TOKEN is required");
if (!APP_URL) throw new Error("APP_URL is required");
if (!JWT_KEY_ID) throw new Error("JWT_KEY_ID is required");
if (!ONEINCH_API_KEY) throw new Error("ONEINCH_API_KEY is required");
if (!ETH_RPC_URL) throw new Error("ETH_RPC_URL is required");

const privateKey = crypto.createPrivateKey(PRIVATE_KEY);

export type JwtPayload = {
	telegram_id: string;
	username: string;
	avatar_url: string;
	sub: string;
	name: string;
	iss: string;
	iat: number;
	exp: number;
};

// biome-ignore lint/suspicious/noExplicitAny: type is missing from sdk
const signJwt = (userData: any) => {
	return sign(
		{
			telegram_id: userData.id,
			username: userData.username ?? "<empty-username>",
			avatar_url: userData.photo_url ?? "<empty-avatar>",
			sub: userData.id.toString(),
			name: userData.first_name,
			iss: "https://api.telegram.org",
			iat: Math.floor(Date.now() / 1000),
			exp: Math.floor(Date.now() / 1000) + 60 * 60, // Token valid for 1 hour
		} satisfies JwtPayload,
		privateKey,
		{ algorithm: "RS256", keyid: JWT_KEY_ID },
	);
};

const app = new Elysia({
	serve: {
		tls:
			NODE_ENV === "development"
				? {
						cert: Bun.file(
							"/Users/martinsaldinger/.vite-plugin-mkcert/cert.pem",
						),
						key: Bun.file("/Users/martinsaldinger/.vite-plugin-mkcert/dev.pem"),
					}
				: undefined,
	},
})

	.use(
		cors({
			allowedHeaders: [
				"Content-Type",
				"Authorization",
				"X-Requested-With",
				"Accept",
			],
			origin: [APP_URL],
			credentials: true,
			methods: ["GET", "POST", "OPTIONS"],
		}),
	)
	.get("/test", () => ({
		message: "Connection successful. Server is running!",
	}))
	.post(
		"/auth/telegram",
		async (ctx) => {
			try {
				// Validate the real initDataRaw using @telegram-apps/init-data-node
				validate(ctx.body.initDataRaw, TELEGRAM_BOT_TOKEN);

				// If validation is successful, parse the data
				const data = new URLSearchParams(ctx.body.initDataRaw);
				// biome-ignore lint/style/noNonNullAssertion: data is validated
				const user = JSON.parse(decodeURIComponent(data.get("user")!));

				// Generate the JWT token
				const JWTtoken = signJwt(user);

				return { token: JWTtoken };
			} catch (error) {
				console.error("Error validating Telegram data:", error);
				throw new Error("Invalid Telegram data");
			}
		},
		{ body: t.Object({ initDataRaw: t.String() }) },
	)
	.post(
		"balance",
		async (ctx) => {
			// const req = await fetch(
			// 	`https://api.1inch.dev/balance/v1.2/1/balances/${ctx.body.walletAddress}`,
			// 	{
			// 		headers: {
			// 			Authorization: `Bearer ${ONEINCH_API_KEY}`,
			// 		},
			// 	},
			// );
			// const json = await req.json();

			const client = createClient({
				chain: mainnet,
				transport: http(ETH_RPC_URL),
			});

			const balance = await readContract(client, {
				abi: erc20Abi,
				functionName: "balanceOf",
				args: [ctx.body.walletAddress as `0x${string}`],
				address: USDC_ADDRESS,
			});

			return { balance: balance.toString() };
		},
		{ body: t.Object({ walletAddress: t.String() }) },
	)
	.listen(3000);

console.log(
	`🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`,
);

export type Server = typeof app;

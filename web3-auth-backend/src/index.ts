import { Elysia, t } from "elysia";
import { cors } from "@elysiajs/cors";
import { edenFetch } from "@elysiajs/eden";
import { sign } from "jsonwebtoken";
import fs from "node:fs";
import path from "node:path";
import crypto from "node:crypto";

import { validate, type User } from "@telegram-apps/init-data-node";

const { TELEGRAM_BOT_TOKEN, APP_URL, NODE_ENV, JWT_KEY_ID } = Bun.env;

const privateKeyPath = path.resolve(__dirname, "../privateKey.pem");

if (!fs.existsSync(privateKeyPath)) throw new Error("privateKey.pem not found");
if (!TELEGRAM_BOT_TOKEN) throw new Error("TELEGRAM_BOT_TOKEN is required");
if (!APP_URL) throw new Error("APP_URL is required");
if (!JWT_KEY_ID) throw new Error("JWT_KEY_ID is required");

const privateKey = crypto.createPrivateKey(
	fs.readFileSync(privateKeyPath, "utf8"),
);

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
	.listen(3000);

console.log(
	`🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`,
);

export type Server = typeof app;

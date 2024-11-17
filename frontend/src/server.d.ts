import { Elysia } from "elysia";
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
declare const app: Elysia<
	"",
	false,
	{
		decorator: {};
		store: {};
		derive: {};
		resolve: {};
	},
	{
		type: {};
		error: {};
	},
	{
		schema: {};
		macro: {};
		macroFn: {};
	},
	{
		test: {
			get: {
				body: unknown;
				params: {};
				query: unknown;
				headers: unknown;
				response: {
					200: {
						message: string;
					};
				};
			};
		};
	} & {
		auth: {
			telegram: {
				post: {
					body: {
						initDataRaw: string;
					};
					params: {};
					query: unknown;
					headers: unknown;
					response: {
						200: {
							token: string;
						};
					};
				};
			};
		};
	} & {
		balance: {
			post: {
				body: {
					walletAddress: string;
				};
				params: {};
				query: unknown;
				headers: unknown;
				response: {
					200: {
						balance: string;
					};
				};
			};
		};
	} & {
		history: {
			post: {
				body: {
					walletAddress: string;
				};
				params: {};
				query: unknown;
				headers: unknown;
				response: {
					200: {
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
		};
	},
	{
		derive: {};
		resolve: {};
		schema: {};
	},
	{
		derive: {};
		resolve: {};
		schema: {};
	}
>;
export type Server = typeof app;
export {};

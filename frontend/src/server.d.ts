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

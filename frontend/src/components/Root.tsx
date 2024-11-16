import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { App } from "../components/App";
import { ErrorBoundary } from "../components/ErrorBoundary";
import { WalletProvider } from "./WalletProvider";

function ErrorBoundaryError({ error }: { error: unknown }) {
	return (
		<div>
			<p>An unhandled error occurred:</p>
			<blockquote>
				<code>
					{error instanceof Error
						? error.message
						: typeof error === "string"
							? error
							: JSON.stringify(error)}
				</code>
			</blockquote>
		</div>
	);
}

const queryClient = new QueryClient();

export function Root() {
	return (
		<ErrorBoundary fallback={ErrorBoundaryError}>
			<QueryClientProvider client={queryClient}>
				<WalletProvider>
					<App />
				</WalletProvider>
			</QueryClientProvider>
		</ErrorBoundary>
	);
}

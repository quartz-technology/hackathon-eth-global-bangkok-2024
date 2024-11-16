import ReactDOM from "react-dom/client";
import { retrieveLaunchParams } from "@telegram-apps/sdk-react";

import { Root } from "./components/Root";
import { EnvUnsupported } from "./components/EnvUnsupported";
import { init } from "./init";

import "@telegram-apps/telegram-ui/dist/styles.css";
import "./index.css";

// Mock the environment in case, we are outside Telegram.
import "./mockEnv.ts";

const root = ReactDOM.createRoot(document.getElementById("root")!);

try {
	// Configure all application dependencies.
	init(retrieveLaunchParams().startParam === "debug" || import.meta.env.DEV);

	root.render(
		<Root />,
	);
} catch (e) {
	root.render(<EnvUnsupported />);
}

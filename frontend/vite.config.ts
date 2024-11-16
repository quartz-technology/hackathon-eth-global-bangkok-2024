import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import svgr from "vite-plugin-svgr";
import mkcert from "vite-plugin-mkcert";
import { nodePolyfills } from "vite-plugin-node-polyfills";

// https://vitejs.dev/config/
export default defineConfig({
	base: "/reactjs-template",
	plugins: [
		// Allows using React dev server along with building a React application with Vite.
		// https://npmjs.com/package/@vitejs/plugin-react-swc
		react(),
		// Create a custom SSL certificate valid for the local machine.
		// https://www.npmjs.com/package/vite-plugin-mkcert
		process.env.NODE_ENV === "developpment" ? mkcert() : undefined,
		nodePolyfills(),
		svgr(),
	],
	publicDir: "./public",
	server: {
		host: true,
	},
	preview: {
		host: "0.0.0.0",
	},
	define: {
		"process.env": {},
	},
});

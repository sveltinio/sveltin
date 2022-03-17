import path from 'path';
import { readFileSync } from 'fs';
import { fileURLToPath } from 'url';

const file = fileURLToPath(new URL('package.json', import.meta.url));
const json = readFileSync(file, 'utf8');
const pkg = JSON.parse(json);

import { mdsvex } from 'mdsvex';
import mdsvexConfig from './mdsvex.config.js';

import preprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte', ...mdsvexConfig.extensions],
	// Learn more at https://github.com/sveltejs/svelte-preprocess
	preprocess: [
		mdsvex(mdsvexConfig),
		preprocess({
			postcss: true,
			preserve: ['ld+json'],
		}),
	],
	kit: {
		trailingSlash: 'always',
		adapter: adapter({
			// default options are shown
			pages: 'build',
			assets: 'build',
			fallback: null,
		}),
		prerender: {
			default: true,
			entries: ['*'],
		},
		vite: {
			define: {
				'process.env.VITE_SVELTEKIT_VERSION': JSON.stringify(
					String(pkg.devDependencies['@sveltejs/kit'])
				),
				'process.env.VITE_BUILD_TIME': JSON.stringify(
					new Date().toISOString()
				),
			},
			server: {
				fs: {
					// Allow serving files from one level up to the project root
					// Alternatevaly set server.fs.strict to false
					allow: ['..'],
				},
			},
			resolve: {
				alias: {
					$config: path.resolve('config'),
					$content: path.resolve('content'),
					$lib: path.resolve('src/lib'),
					$themes: path.resolve('themes'),
				},
			},
			optimizeDeps: {
				include: ['@indaco/svelte-iconoir'],
			},
		},
	},
};

export default config;

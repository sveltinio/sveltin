import path from 'path';

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
			postcss: false,
			preserve: ['ld+json'],
		}),
	],
	kit: {
		adapter: adapter({
			// default options are shown
			pages: 'build',
			assets: 'build',
			fallback: null,
		}),
		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',
		vite: {
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

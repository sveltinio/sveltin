import path from 'path';

import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

import { mdsvex } from 'mdsvex';
import mdsvexConfig from './mdsvex.config.js';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte', ...mdsvexConfig.extensions],
	// Learn more at https://github.com/sveltejs/svelte-preprocess
	preprocess: [
		mdsvex(mdsvexConfig),
		preprocess({
			preserve: ['ld+json'],
			scss: {
				prependData: '@use "src/variables.scss" as *;',
			},
		}),
	],
	kit: {
		adapter: adapter({
			// default options are shown
			pages: 'build',
			assets: 'build',
			fallback: null,
		}),
		vite: {
			server: {
				fs: {
					// Allow serving files from one level up to the project root
					// Alternatevaly set server.fs.strict to false
					allow: ['..'],
				},
			},
			css: {
				preprocessorOptions: {
					scss: {
						additionalData: '@use "src/variables.scss" as *;',
					},
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

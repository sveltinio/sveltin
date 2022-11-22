import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

import { mdsvex } from 'mdsvex';
import mdsvexConfig from './mdsvex.config.js';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte', ...mdsvexConfig.extensions],
	preprocess: [
		mdsvex(mdsvexConfig),
		preprocess({
			preserve: ['ld+json'],
			scss: {
				prependData: '@use "src/_variables.scss" as *;',
			},
		}),
	],
	kit: {
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			fallback: '200.html',
			precompress: true,
		}),
		prerender: {
			crawl: true,
			enabled: true,
			entries: ['*'],
		},
	},
};

export default config;

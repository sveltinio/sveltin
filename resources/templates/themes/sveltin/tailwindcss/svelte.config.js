import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { mdsvex } from 'mdsvex';
import mdsvexConfig from './mdsvex.config.js';
import preprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-static';

const sveltinJsonFile = fileURLToPath(new URL('sveltin.json', import.meta.url));
const sveltinJson = readFileSync(sveltinJsonFile, 'utf8');
const { sveltekit } = JSON.parse(sveltinJson);

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte', ...mdsvexConfig.extensions],
	preprocess: [
		mdsvex(mdsvexConfig),
		preprocess({
			postcss: true,
			preserve: ['ld+json'],
		}),
	],
	kit: {
		adapter: adapter({
			pages: sveltekit.adapter.pages,
			assets: sveltekit.adapter.assets,
			fallback: sveltekit.adapter.fallback,
			precompress: sveltekit.adapter.precompress,
			strict: sveltekit.adapter.strict,
		}),
		prerender: {
			crawl: true,
			entries: ['*'],
			/**
			 * While prerendering, silently ignore the failure "no element with id exists on..."
			 * This is due to the mdsvex component import on the client-side with onMount, so
			 * before that prerender cannot find ids.
			 */
			handleMissingId: 'ignore',
		},
	},
};

export default config;

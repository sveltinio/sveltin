import { defineMDSveXConfig as defineConfig } from 'mdsvex';
import relativeImages from 'mdsvex-relative-images';

import emoji from 'remark-emoji';
import remarkSlug from 'remark-slug';
import readingTime from 'remark-reading-time';
import rehypeAutoLinkHeadings from 'rehype-autolink-headings';
import rehypeExternalLinks from 'rehype-external-links';
import rehypeSlug from 'rehype-slug';
import headings from './src/lib/utils/headings.js';

const mdsvexConfig = defineConfig({
	extensions: ['.svelte.md', '.md', '.svx'],
	smartypants: {
		dashes: 'oldschool'
	},
	remarkPlugins: [remarkSlug, headings, emoji, readingTime(), relativeImages],
	rehypePlugins: [
		[rehypeExternalLinks, { target: '_blank', rel: ['noopener', 'noreferrer'] }],
		rehypeSlug[
			(rehypeAutoLinkHeadings,
			{
				behavior: 'wrap'
			})
		]
	]
});

export default mdsvexConfig;

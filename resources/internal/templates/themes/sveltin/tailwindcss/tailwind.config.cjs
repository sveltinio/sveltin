const plugin = require('tailwindcss/plugin');

const config = {
	content: [
		'./src/routes/*.{html,svelte,js,ts}',
		'./themes/**/*.{html,svelte,js,ts}',
	],
	theme: {
		extend: {
			colors: {
				base: 'var(--base)',
				muted: 'var(--muted)',
				accent: 'var(--cyan)',
				dark: 'var(--dark)',
				deepDark: 'var(--deep-dark)',
				river: 'var(--river)',
				haiti: 'var(--haiti)',
				pearl: 'var(--pearl)',
				sant: 'var(--santa)',
				cege: 'var(--cege)',
			},
			fontFamily: {
				mono: ['"IBM Plex Mono"', 'ui-monospace'],
				sans: [
					'"IBM Plex Sans"',
					'system-ui',
					'-apple-system',
					'BlinkMacSystemFont',
					'"Segoe UI"',
					'Roboto',
					'"Helvetica Neue"',
					'Arial',
					'"Noto Sans"',
					'sans-serif',
					'"Apple Color Emoji"',
					'"Segoe UI Emoji"',
					'"Segoe UI Symbol"',
				],
			},
		},
	},
	plugins: [
		require('@tailwindcss/typography'),
		require('@tailwindcss/line-clamp'),
		require('@tailwindcss/aspect-ratio'),
		plugin(function ({ addVariant, e, postcss }) {
			addVariant('firefox', ({ container, separator }) => {
				const isFirefoxRule = postcss.atRule({
					name: '-moz-document',
					params: 'url-prefix()',
				});
				isFirefoxRule.append(container.nodes);
				container.append(isFirefoxRule);
				isFirefoxRule.walkRules((rule) => {
					rule.selector = `.${e(
						`firefox${separator}${rule.selector.slice(1)}`
					)}`;
				});
			});
		}),
	],
};

module.exports = config;

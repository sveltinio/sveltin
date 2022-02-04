const plugin = require('tailwindcss/plugin');

const config = {
	content: [
		'./src/routes/*.{html,svelte,js,ts}',
		'./themes/**/*.{html,svelte,js,ts}',
	],
	theme: {
		extend: {
			textColor: {
				skin: {
					white: 'var(--pure-white)',
					dark: 'var(--dark)',
					base: 'var(--color-text-base)',
					muted: 'var(--color-text-muted)',
					accent: 'var(--color-text-accent)',
				},
			},
			backgroundColor: {
				skin: {
					white: 'var(--pure-white)',
					dark: 'var(--dark)',
					'deep-dark': 'var(--deep-dark)',
				},
			},
			colors: {
				haiti: '#2c2c35',
				pearl: '#1e2028',
				river: '#464a5d',
				santa: '#a0a1b2',
				cege: '#0B7599',
				auburn: '#9e2a2a',
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

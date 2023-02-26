import type { Theme } from '@unocss/preset-uno';
import { defineConfig, presetUno, presetTypography } from 'unocss';
import transformerDirectives from '@unocss/transformer-directives';

export default defineConfig({
	theme: <Theme>{
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
			mono: ['"IBM Plex Mono"', 'ui-monospace'].join(','),
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
			].join(','),
		},
	},
	shortcuts: [
		{ 'markdown-body': 'mx-auto prose prose-gray text-justify' },
		{ 'max-width-none': 'max-w-none' },
	],
	transformers: [transformerDirectives()],
	presets: [
		presetUno(),
		presetTypography({
			cssExtend: {
				h2: {
					'font-weight': '500',
					'font-size': '1.5em',
					'margin-top': '2em',
					'margin-bottom': '1em',
					'line-height': '1.3333333',
				},
				a: {
					'text-decoration': 'none',
				},
			},
		}),
	],
});

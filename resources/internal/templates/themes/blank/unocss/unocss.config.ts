import { defineConfig, presetUno, presetTypography } from 'unocss';
import transformerDirectives from '@unocss/transformer-directives';

export default defineConfig({
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

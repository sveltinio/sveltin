const cssnano = require('cssnano');
const autoprefixer = require('autoprefixer');
const postcssImport = require('postcss-import');
const nesting = require('tailwindcss/nesting');
const tailwindcss = require('tailwindcss');

const mode = process.env.NODE_ENV;
const dev = mode === 'development';

const config = {
	plugins: [
		postcssImport,
		nesting,
		tailwindcss,
		autoprefixer,
		!dev && cssnano,
	],
};

module.exports = config;

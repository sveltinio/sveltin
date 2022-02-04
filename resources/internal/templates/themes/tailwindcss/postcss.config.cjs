const cssnano = require('cssnano');

const mode = process.env.NODE_ENV;
const dev = mode === 'development';

const config = {
	plugins: [
		require('postcss-import'),
		require('tailwindcss/nesting')(require('postcss-nesting')),
		require('tailwindcss'),
		require('autoprefixer'),
		!dev &&
			cssnano({
				preset: 'default',
			}),
	],
};

module.exports = config;

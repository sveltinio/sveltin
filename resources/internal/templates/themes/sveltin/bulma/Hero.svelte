<script>
	import { Link } from '@sveltinio/essentials';
</script>

<section>
	<h1>
		Your Sveltin project is <span class="accent">up & running!</span>
	</h1>

	<div class="buttons-wrapper">
		<Link
			id="quickStartLink"
			external
			icon={false}
			href="https://docs.sveltin.io/quick-start"
			class="hero-btn hero-btn--accent">Quick Start</Link
		>

		<Link
			id="documentationLink"
			external
			icon={false}
			href="https://docs.sveltin.io"
			class="hero-btn hero-btn--base">Read the Docs</Link
		>
	</div>
</section>

<style lang="scss">
	@import 'bulma/sass/utilities/_all';

	section {
		max-width: 80rem /* 1280px */;
		margin-left: auto;
		margin-right: auto;
		padding-top: 11rem /* 176px */;
		padding-bottom: 11rem /* 176px */;
		padding-left: 1.5rem /* 24px */;
		padding-right: 1.5rem /* 24px */;
	}

	h1 {
		font-size: 1.875rem /* 30px */;
		font-weight: 700;
		text-align: center;
		letter-spacing: -0.05em;
		padding-bottom: 1rem /* 16px */;
		line-height: 1.25;

		@include from($sm) {
			font-size: 2.25rem /* 36px */;
		}

		@include from($lg) {
			font-size: 3rem /* 48px */;
		}
	}

	.accent {
		color: $color-text-accent;
		font-style: italic;
		font-weight: 400;
		display: block;
		text-transform: capitalize;
	}

	.buttons-wrapper {
		position: relative;
		display: inline-flex;
		justify-content: center;
		width: 100%;
		margin-left: auto;
		margin-right: auto;
		margin-top: 1rem /* 16px */;
	}

	:global {
		.hero-btn {
			position: relative;
			display: inline-flex;
			align-items: center;
			padding-top: 0.5rem /* 8px */;
			padding-right: 0.5rem /* 8px */;
			padding-bottom: 0.5rem /* 8px */;
			padding-left: 0.5rem /* 8px */;
			font-size: 0.875rem /* 14px */;
			font-weight: 500;
			border-width: 2px;
			border-style: solid;

			@include from($lg) {
				padding-left: 0.75rem /* 8px */;
				padding-right: 0.75rem /* 8px */;
			}
		}

		.hero-btn--base {
			margin-left: -1px;
			color: $dark;
			border-top-right-radius: 0.5rem /* 8px */;
			border-bottom-right-radius: 0.5rem /* 8px */;
			border-color: $color-text-base;

			&:focus {
				z-index: 10;
				outline: 2px solid transparent;
				outline-offset: 2px;
			}
		}

		.hero-btn--accent {
			color: $white;
			background-color: $color-text-base;
			border-top-left-radius: 0.5rem /* 8px */;
			border-bottom-left-radius: 0.5rem /* 8px */;
			border-color: $color-text-base;

			&:hover {
				color: $white;
				background-color: $dark;
			}

			&:focus {
				z-index: 10;
				outline: 2px solid transparent;
				outline-offset: 2px;
				border-color: $white;
			}
		}
	}
</style>

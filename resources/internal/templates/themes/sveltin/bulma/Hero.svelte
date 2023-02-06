<script>
	import { Link } from '@sveltinio/essentials';
</script>

<section class="mx-auto px-6 py-44 container is-max-widescreen">
	<h1
		class="pb-4 has-text-centered has-text-weight-bold is-size-3-mobile is-size-2-tablet is-size-1-desktop"
	>
		Your Sveltin project is <span
			class="is-block accent is-italic has-text-weight-normal is-capitalized"
			>up & running!</span
		>
	</h1>

	<div
		class="is-relative mx-auto is-inline-flex is-flex-direction-row is-justify-content-center mt-4"
		style="width: 100%;"
	>
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

	h1 {
		line-height: 1.25;
		letter-spacing: -0.025em;

		@include from($sm) {
			font-size: 2.25rem /* 36px */;
		}

		@include from($lg) {
			font-size: 3rem /* 48px */;
		}
	}

	.accent {
		color: $cyan;
	}

	:global {
		.hero-btn {
			position: relative;
			display: inline-flex;
			align-items: center;
			padding-top: 0.5rem;
			padding-right: 0.5rem;
			padding-bottom: 0.5rem;
			padding-left: 0.5rem;
			font-size: 0.875rem;
			font-weight: 500;
			line-height: 1.25rem;
			border: 2px solid $base;

			&:focus {
				outline: 2px solid transparent;
				outline-offset: 2px;
			}

			@include from($lg) {
				padding-right: 2rem;
				padding-left: 2rem;
			}
		}

		.hero-btn--base {
			margin-left: -1px;
			color: $dark;
			border-top-right-radius: 0.5rem;
			border-bottom-right-radius: 0.5rem;
		}

		.hero-btn--accent {
			color: $white;
			font-weight: 500;
			background-color: $base;
			border-top-left-radius: 0.5rem;
			border-bottom-left-radius: 0.5rem;

			&:hover {
				color: $white;
				background-color: $dark;
			}
		}
	}
</style>

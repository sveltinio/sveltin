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
			class="cta-btn cta-btn--accent">Quick Start</Link
		>

		<Link
			id="documentationLink"
			external
			icon={false}
			href="https://docs.sveltin.io"
			class="cta-btn cta-btn--base">Read the Docs</Link
		>
	</div>
</section>

<style>
	section {
		max-width: 80rem;
		margin-left: auto;
		margin-right: auto;
		padding-top: 11rem;
		padding-bottom: 11rem;
		padding-left: 1.5rem;
		padding-right: 1.5rem;
	}

	h1 {
		font-size: 1.875rem;
		font-weight: 700;
		text-align: center;
		letter-spacing: -0.025em;
		line-height: 1.25;
		padding-bottom: 1rem;
	}

	.accent {
		color: var(--accent);
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
		margin-top: 1rem;
	}

	:global(.cta-btn) {
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
		border: 2px solid var(--base);
	}

	:global(.cta-btn:focus) {
		outline: 2px solid transparent;
		outline-offset: 2px;
	}

	:global(.cta-btn--base) {
		margin-left: -1px;
		color: var(--dark);
		border-top-right-radius: 0.5rem;
		border-bottom-right-radius: 0.5rem;
	}

	:global(.cta-btn--accent) {
		color: var(--white);
		background-color: var(--base);
		border-top-left-radius: 0.5rem;
		border-bottom-left-radius: 0.5rem;
	}

	:global(.cta-btn--accent:hover) {
		color: var(--white);
		background-color: var(--dark);
	}

	@media only screen and (min-width: 640px) {
		h1 {
			font-size: 2.25rem;
		}
	}

	@media only screen and (min-width: 1024px) {
		h1 {
			font-size: 3rem;
		}

		:global(.cta-btn) {
			padding-right: 2rem;
			padding-left: 2rem;
		}
	}
</style>

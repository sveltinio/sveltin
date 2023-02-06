<script lang="ts">
	import { sveltinVersion } from '$config/defaults.js';

	export let logoSrc: string;
</script>

<footer>
	<div class="wrapper">
		<div class="content">
			<a href="/">
				<img
					src={logoSrc}
					alt="sveltin logo"
					style="width: 100px; height: 48px;"
				/>
				<span class="sr-only">Sveltin Logo</span>
			</a>
			<p>
				The CLI to Boost Your Productivity creating SvelteKit powered
				static websites.
			</p>
			<p class="semibold">You are running Sveltin v{sveltinVersion}</p>
		</div>
	</div>
</footer>

<style lang="scss">
	footer {
		width: 100%;
		background-color: $deep-dark;
		font-size: 0.875rem;
		line-height: 1.25rem;
		color: $muted;
	}

	.wrapper {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		flex-wrap: wrap;
		width: 100%;
		padding-top: 2rem;
		padding-right: 1.25rem;
		padding-bottom: 1.5rem;
		padding-left: 1.25rem;
		margin-left: auto;
		margin-right: auto;
	}

	.content {
		flex-shrink: 0;
		text-align: left;
		width: 18rem;

		p {
			margin-top: 1.5rem;
			font-size: 0.875rem;
			color: $muted;
		}
	}

	.semibold {
		font-weight: 600;
	}

	@media (min-width: 640px) {
		.wrapper {
			padding-left: 1.5rem;
			padding-right: 1.5rem;
		}
	}

	@media (min-width: 768px) {
		.wrapper {
			flex-direction: row;
			align-items: center;
		}
	}

	@media (min-width: 1024px) {
		.wrapper {
			padding-left: 3rem;
			padding-right: 3rem;
			padding-top: 3rem;
			padding-bottom: 3rem;
			flex-wrap: nowrap;
		}
	}
</style>

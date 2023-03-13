<script lang="ts">
	import { sveltinVersion, sveltekitVersion } from '$config/defaults.js';

	export let logoSrc: string;
</script>

<footer class="is-full is-size-7" style="line-height: 1.25rem;">
	<div
		class="is-full is-flex is-flex-direction-column is-align-items-start mx-auto wrapper "
	>
		<div class="is-flex-shrink-0 has-text-left" style="width: 18rem">
			<a href="/">
				<img
					src={logoSrc}
					alt="sveltin logo"
					style="width: 100px; height: 48px;"
				/>
				<span class="is-sr-only">Sveltin Logo</span>
			</a>

			<p
				class="has-text-grey-light has-text-weight-medium mt-6"
				style="letter-spacing: -0.2px;"
			>
				The CLI to Boost Your Productivity creating SvelteKit powered
				static websites.
			</p>
			<p
				class="has-text-weight-semibold has-text-grey-light has-text-weight-medium mt-6 is-size-7"
			>
				You are running Sveltin v{sveltinVersion} on SvelteKit v{sveltekitVersion}
			</p>
		</div>
	</div>
</footer>

<style lang="scss">
	@import 'bulma/sass/utilities/_all';

	footer {
		background-color: $deep-dark;
	}

	.wrapper {
		padding-top: 2rem;
		padding-right: 1.25rem;
		padding-bottom: 1.5rem;
		padding-left: 1.25rem;
		flex-wrap: wrap;

		@include from($sm) {
			padding-left: 1.5rem;
			padding-right: 1.5rem;
		}

		@include from($md) {
			flex-direction: row;
			align-items: center;
		}

		@include from($lg) {
			padding-left: 3rem;
			padding-right: 3rem;
			padding-top: 3rem;
			padding-bottom: 3rem;
			flex-wrap: nowrap;
		}
	}
</style>

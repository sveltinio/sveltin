<script lang="ts">
	import { base } from '$app/paths';
	import { sveltinVersion } from '$config/defaults.js';
	import type { IWebSite } from '@sveltinio/seo/types';

	export let websiteData: IWebSite;
</script>

<footer
	class="is-full has-background-grey-darker is-size-7"
	style="line-height: 1.25rem;"
>
	<div
		class="wrapper is-full is-flex is-flex-direction-column is-align-items-start is-flex-wrap-wrap mx-auto"
	>
		<div class="is-flex-shrink-0 has-text-left" style="width: 18rem">
			<div>
				<a href="/">
					<img
						src="{base}/{websiteData.logo}"
						alt="sveltin logo"
						style="width: 100px; height: 48px;"
					/>
					<span class="is-sr-only">Sveltin Logo</span>
				</a>
			</div>
			<p
				class="has-text-grey-light has-text-weight-medium mgt-5"
				style="letter-spacing: -0.2px;"
			>
				The CLI to Boost Your Productivity creating SvelteKit powered
				static websites.
			</p>
			<p
				class="has-text-weight-bold has-text-grey-light has-text-weight-medium mgt-6 is-size-7"
				style="letter-spacing: -0.2px;"
			>
				You are running Sveltin v{sveltinVersion}
			</p>
		</div>
	</div>
</footer>

<style lang="scss">
	@import 'bulma/sass/utilities/_all';

	.wrapper {
		padding-top: 2rem;
		padding-right: 1.25rem;
		padding-bottom: 1.5rem;
		padding-left: 1.25rem;

		@include from($sm) {
			padding-left: 1.5rem;
			padding-right: 1.5rem;
		}

		@include from($md) {
			flex-direction: row;
			align-items: center;
		}

		@include from($lg) {
			padding-left: 3rem /* 48px */;
			padding-right: 3rem /* 48px */;
			padding-top: 3rem /* 48px */;
			padding-bottom: 3rem /* 48px */;
			align-items: flex-start;
			flex-wrap: nowrap;
		}
	}
</style>

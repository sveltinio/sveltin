<script lang="ts">
	import { sveltinVersion } from '$config/defaults.js';
	import type { IWebSite } from '@sveltinio/seo/types';

	export let websiteData: IWebSite;
</script>

<footer class="lh-base bg-dark text-muted" style="line-height: 1.25rem;">
	<div class="wrapper w-100 mx-auto">
		<div class="flex-shrink-0 text-start w-xxs">
			<div>
				<a href="/">
					<img
						src="{websiteData.baseURL}/{websiteData.logo}"
						alt="sveltin logo"
						style="width: 100px; height: 48px;"
					/>
					<span class="visually-hidden">Sveltin Logo</span>
				</a>
			</div>
			<p class="mt-5">
				The CLI to Boost Your Productivity creating SvelteKit powered
				static websites.
			</p>
			<p class="mt-5">
				<strong>You are running Sveltin v{sveltinVersion}</strong>
			</p>
		</div>
	</div>
</footer>

<style lang="scss">
	p {
		font-size: $font-size-sm;
	}

	.wrapper {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		flex-wrap: wrap;
		padding-left: 1.25rem /* 20px */;
		padding-right: 1.25rem /* 20px */;
		padding-top: 2rem /* 32px */;
		padding-bottom: 1.5rem /* 24px */;
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
			padding-left: 3rem /* 48px */;
			padding-right: 3rem /* 48px */;
			padding-top: 3rem /* 48px */;
			padding-bottom: 2.25rem /* 48px */;
			align-items: flex-start;
			flex-wrap: nowrap;
		}
	}
</style>

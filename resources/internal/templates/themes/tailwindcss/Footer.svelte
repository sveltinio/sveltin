<script lang="ts">
	import { sveltinVersion } from '$config/defaults.js';
	import type { IWebSite } from '@sveltinio/seo/types';

	export let websiteData: IWebSite;
</script>

<footer class="bg-skin-deep-dark">
	<div
		class="flex flex-col items-start flex-wrap w-full px-5 pb-6 pt-8 mx-auto md:items-center md:flex-row lg:px-12 lg:py-12 lg:items-start lg:flex-nowrap"
	>
		<div class="flex-shrink-0 text-left w-72 lg:mx-0">
			<div>
				<a href="/">
					<img
						src="{websiteData.baseURL}/{websiteData.logo}"
						alt="sveltin logo"
						style="width: 100px; height: 48px;"
					/>
					<span class="sr-only">Sveltin Logo</span>
				</a>
			</div>
			<p class="text-skin-muted mt-6 text-sm">
				The CLI to Boost Your Productivity creating SvelteKit powered
				static websites.
			</p>
			<p class="text-skin-muted mt-6 text-sm font-semibold">
				You are running Sveltin v{sveltinVersion}
			</p>
		</div>
	</div>
</footer>

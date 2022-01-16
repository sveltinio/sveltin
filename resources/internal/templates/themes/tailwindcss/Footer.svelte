<script lang="ts">
	import { sveltinVersion } from '$config/defaults.js';
	import type { WebSite } from '@sveltinio/seo/types';

	export let websiteData: WebSite;
</script>

<footer>
	<div class="content">
		<div class="wrapper">
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
			<p>
				The CLI to Boost Your Productivity creating SvelteKit powered
				static websites.
			</p>
			<p><strong>You are running Sveltin v{sveltinVersion}</strong></p>
		</div>
	</div>
</footer>

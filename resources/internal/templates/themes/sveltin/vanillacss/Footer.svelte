<script lang="ts">
	import { base } from '$app/paths';
	import { sveltinVersion } from '$config/defaults.js';
	import type { IWebSite } from '@sveltinio/seo/types';

	export let websiteData: IWebSite;
</script>

<footer>
	<div class="content">
		<div class="wrapper">
			<div>
				<a href="/">
					<img
						src="{base}/{websiteData.logo}"
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

<style>
	footer {
		width: 100%;
		background-color: rgb(30 32 40/1);
		font-size: 0.875rem;
		line-height: 1.25rem;
		color: var(--color-text-muted);
	}

	footer > .content {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		flex-wrap: wrap;
		width: 100%;
		padding-top: 2rem /* 32px */;
		padding-right: 1.25rem /* 20px */;
		padding-bottom: 1.5rem /* 24px */;
		padding-left: 1.25rem /* 20px */;
		margin-left: auto;
		margin-right: auto;
	}

	footer > .content > .wrapper {
		flex-shrink: 0;
		text-align: left;
		width: 18rem /* 288px */;
	}

	footer > .content > .wrapper > p {
		margin-top: 1.5rem /* 24px */;
		font-size: 0.875rem /* 14px */;
		color: var(--color-text-muted);
	}

	@media (min-width: 640px) {
		footer > .content {
			padding-left: 1.5rem /* 24px */;
			padding-right: 1.5rem /* 24px */;
		}
	}

	@media (min-width: 768px) {
		footer > .content {
			flex-direction: row;
			align-items: center;
		}
	}

	@media (min-width: 1024px) {
		footer > .content {
			padding-left: 3rem /* 48px */;
			padding-right: 3rem /* 48px */;
			padding-top: 3rem /* 48px */;
			padding-bottom: 3rem /* 48px */;
			align-items: flex-start;
			flex-wrap: nowrap;
		}
	}
</style>

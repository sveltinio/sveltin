<script lang="ts">
	import { Link } from '@sveltinio/essentials/link';
	import { sveltinVersion, sveltekitVersion } from '$config/defaults.js';

	export let logoSrc: string;
</script>

<footer class="w-full text-sm text-muted bg-deepDark">
	<div
		class="flex flex-col items-start flex-wrap w-full px-5 pb-6 pt-8 mx-auto md:items-center md:flex-row lg:px-12 lg:py-12 lg:items-start lg:flex-nowrap"
	>
		<div class="flex-shrink-0 text-left w-72 lg:mx-0">
			<Link id="quickStartLink" href="/">
				<img
					src={logoSrc}
					alt="sveltin logo"
					style="width: 100px; height: 48px;"
				/>
				<span class="sr-only">Sveltin Logo</span>
			</Link>

			<p class="mt-6">
				The CLI to Boost Your Productivity creating SvelteKit powered
				static websites.
			</p>
			<p class="mt-6 font-semibold">
				You are running Sveltin v{sveltinVersion} on SvelteKit v{sveltekitVersion}
			</p>
		</div>
	</div>
</footer>

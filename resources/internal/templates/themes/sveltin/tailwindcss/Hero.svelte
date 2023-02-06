<script>
	import { Link } from '@sveltinio/essentials';
</script>

<section class="mx-auto px-6 py-44 max-w-7xl">
	<h1
		class="pb-4 mx-auto text-center text-skin-base text-3xl font-bold tracking-tight sm:text-4xl lg:text-5xl lg:leading-tight"
	>
		Your Sveltin project is <span
			class="block italic font-normal text-skin-accent capitalize"
			>up & running!</span
		>
	</h1>

	<div class="relative inline-flex justify-center w-full mx-auto mt-4">
		<Link
			id="quickStartLink"
			external
			icon={false}
			href="https://docs.sveltin.io/quick-start"
			class="relative inline-flex items-center p-2 text-white text-sm font-medium bg-river border-2 border-river rounded-l-lg hover:bg-haiti focus:outline-none lg:px-8"
			>Quick Start</Link
		>

		<Link
			id="documentationLink"
			external
			icon={false}
			href="https://docs.sveltin.io"
			class="relative inline-flex items-center p-2 -ml-px text-sm font-medium border-2 border-river rounded-r-lg focus:outline-none lg:px-8"
			>Read the Docs</Link
		>
	</div>
</section>

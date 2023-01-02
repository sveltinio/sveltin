<script>
	import { Link } from '@sveltinio/essentials';
</script>

<section class="hero">
	<div class="mx-auto max-w-7xl">
		<div class="py-44">
			<div class="max-w-5xl px-6 mx-auto">
				<div
					class="pb-24 text-center flex flex-col items-center w-full mx-auto"
				>
					<h1
						class="pb-4 mx-auto text-skin-base text-3xl font-bold tracking-tighter sm:text-4xl lg:text-5xl lg:leading-tight"
					>
						Your Sveltin project is <span
							class="block italic font-normal text-skin-accent capitalize"
							>up & running!</span
						>
					</h1>

					<div
						class="inline-flex justify-center max-w-md mx-auto mt-4"
					>
						<Link
							id="quickStartLink"
							external
							icon={false}
							href="https://docs.sveltin.io/quick-start"
							class="hero-btn hero-btn-accent">Quick Start</Link
						>

						<Link
							id="documentationLink"
							external
							icon={false}
							href="https://docs.sveltin.io"
							class="hero-btn hero-btn-base">Read the Docs</Link
						>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>

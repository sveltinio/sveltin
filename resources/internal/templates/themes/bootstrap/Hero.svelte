<script>
	import { ExternalLink } from '@sveltinio/essentials';
</script>

<section class="d-flex justify-content-center text-center">
	<div class="mx-auto mw-7xl">
		<div class="mx-auto w-full py-44">
			<div class="subwrapper mx-auto">
				<div class="pb-24 text-center">
					<h1 class="pb-4 mx-auto fw-bold text-center lh-sm">
						Your Sveltin project is <span
							class="text-teal block text-capitalize fst-italic fw-normal"
							>up & running!</span
						>
					</h1>

					<div
						class="mx-auto mt-4 d-inline-flex justify-content-center btn-group"
						role="group"
					>
						<ExternalLink
							id="learnMoreLink"
							name="learnMore"
							url="https://docs.sveltin.io/tutorials"
							alt="link to the tutorial pages"
							class="btn btn-accent">Learn More</ExternalLink
						>

						<ExternalLink
							id="documentationLink"
							name="docs"
							url="https://docs.sveltin.io"
							alt="link to the documentation website"
							class="btn btn-base">Read the Docs</ExternalLink
						>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>

<style lang="scss">
	h1 {
		font-size: 1.875rem;
		letter-spacing: -0.05em;
	}

	.subwrapper {
		width: 22rem;
		padding-left: 1.5rem;
		padding-right: 1.5rem;
	}

	@media (min-width: 640px) {
		h1 {
			font-size: 2.25rem;
		}
		.subwrapper {
			width: 30rem;
			padding-left: 4rem;
			padding-right: 4rem;
		}
	}

	@media (min-width: 1024px) {
		h1 {
			font-size: 3rem;
		}

		.subwrapper {
			width: 38rem;
			padding-left: 4rem;
			padding-right: 4rem;
		}
	}
</style>

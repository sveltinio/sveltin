<script>
	import { ExternalLink } from '@sveltinio/essentials';
</script>

<section class="columns is-centered has-text-center">
	<div class="column">
		<div class="mx-auto is-full pgy-44">
			<div class="container subwrapper mx-auto">
				<div
					class="mx-auto is-full has-text-left is-flex is-flex-direction-column is-justify-content-center"
				>
					<h1
						class="has-text-centered has-text-grey-dark has-text-weight-bold pgb-4"
					>
						Your Sveltin project is <span
							class="is-block has-text-weight-normal has-text-info is-italic is-capitalized"
							>up & running!</span
						>
					</h1>

					<div
						class="is-full is-inline-flex is-justify-content-center mx-auto mgt-4"
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
	@import 'bulma/sass/utilities/_all';

	.subwrapper {
		max-width: 64rem /* 1024px */;
		padding-left: 1.5rem /* 24px */;
		padding-right: 1.5rem /* 24px */;

		@include from($lg) {
			padding-left: 4rem /* 64px */;
			padding-right: 4rem /* 64px */;
		}
	}

	h1 {
		font-size: 1.875rem /* 30px */;
		letter-spacing: -0.05em;
		line-height: 1.25;

		@include from($sm) {
			font-size: 2.25rem /* 36px */;
		}

		@include from($lg) {
			font-size: 3rem /* 48px */;
		}
	}
</style>

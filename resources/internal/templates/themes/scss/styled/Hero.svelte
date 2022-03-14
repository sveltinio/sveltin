<script>
	import { ExternalLink } from '@sveltinio/essentials';
</script>

<section class="hero">
	<div class="content">
		<div class="wrapper">
			<div class="subwrapper">
				<div class="herotext">
					<h1>
						Your Sveltin project is <span class="accent"
							>up & running!</span
						>
					</h1>

					<div class="buttons-wrapper">
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
	.hero > .content {
		max-width: 80rem;
		margin-left: auto;
		margin-right: auto;

		> .wrapper {
			padding-top: 11rem;
			padding-bottom: 11rem;
		}
	}

	.subwrapper {
		max-width: 64rem;
		padding-left: 1.5rem;
		padding-right: 1.5rem;
		margin-left: auto;
		margin-right: auto;
	}

	.herotext {
		padding-bottom: 6rem;
		text-align: center;

		h1 {
			font-size: 1.875rem;
			font-weight: 700;
			text-align: center;
			letter-spacing: -0.05em;
			padding-bottom: 1rem;
			line-height: 1.25;
			$color: $dark;
		}
	}

	.accent {
		color: $color-text-accent;
		font-style: italic;
		font-weight: 400;
		display: block;
		text-transform: capitalize;
	}

	.buttons-wrapper {
		position: relative;
		display: inline-flex;
		justify-content: center;
		width: 100%;
		margin-left: auto;
		margin-right: auto;
		margin-top: 1rem;
	}

	@media (min-width: 640px) {
		.herotext h1 {
			font-size: 2.25rem;
		}
	}

	@media (min-width: 1024px) {
		.subwrapper {
			padding-left: 4rem;
			padding-right: 4rem;
		}

		.herotext h1 {
			font-size: 3rem;
		}
	}
</style>

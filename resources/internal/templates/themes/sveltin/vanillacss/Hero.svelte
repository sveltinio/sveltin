<script>
	import { Link } from '@sveltinio/essentials';
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
						<Link
							id="learnMoreLink"
							external
							icon={false}
							url="https://docs.sveltin.io/tutorials"
							alt="link to the tutorial pages"
							class="hero-btn hero-btn-accent">Learn More</Link
						>

						<Link
							id="documentationLink"
							external
							icon={false}
							url="https://docs.sveltin.io"
							alt="link to the documentation website"
							class="hero-btn hero-btn-base">Read the Docs</Link
						>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>

<style>
	.hero > .content {
		max-width: 80rem /* 1280px */;
		margin-left: auto;
		margin-right: auto;
	}

	.hero > .content > .wrapper {
		padding-top: 11rem /* 176px */;
		padding-bottom: 11rem /* 176px */;
	}

	.subwrapper {
		max-width: 64rem /* 1024px */;
		padding-left: 1.5rem /* 24px */;
		padding-right: 1.5rem /* 24px */;
		margin-left: auto;
		margin-right: auto;
	}

	.herotext {
		padding-bottom: 6rem /* 96px */;
		text-align: center;
	}

	.herotext h1 {
		font-size: 1.875rem /* 30px */;
		font-weight: 700;
		text-align: center;
		letter-spacing: -0.05em;
		padding-bottom: 1rem /* 16px */;
		line-height: 1.25;
	}

	.accent {
		color: var(--color-text-accent);
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
		margin-top: 1rem /* 16px */;
	}

	@media (min-width: 640px) {
		.herotext h1 {
			font-size: 2.25rem /* 36px */;
		}
	}

	@media (min-width: 1024px) {
		.subwrapper {
			padding-left: 4rem /* 64px */;
			padding-right: 4rem /* 64px */;
		}

		.herotext h1 {
			font-size: 3rem /* 48px */;
		}
	}
</style>

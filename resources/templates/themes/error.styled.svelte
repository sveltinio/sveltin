<script>
	import { page } from '$app/stores';
</script>

<svelte:head>
	<title>{$page.status}: {$page.error.message}</title>
</svelte:head>
<div class="error-container">
	<div class="wrapper">
		<main class="main">
			<p class="error-code">{$page.status}</p>
			<div class="message-wrapper">
				<div class="message-main">
					<h1 class="error-title">
						{$page.error.message}
					</h1>
					<p class="error-text">
						Please check the URL in the address bar and try again.
					</p>
				</div>
			</div>
		</main>
	</div>
</div>

<style>
	.error-container {
		min-height: 100%;
		padding-top: 4rem;
		padding-right: 1rem;
		padding-bottom: 4rem;
		padding-left: 1rem;
	}

	.wrapper {
		max-width: max-content;
		margin-left: auto;
		margin-right: auto;
	}

	.main {
		justify-content: center;
		align-items: center;
	}

	.error-code {
		font-size: 2.25rem;
		line-height: 2.5rem;
		font-weight: 800;
		color: #0b7599;
	}

	.error-title {
		font-size: 2.25rem;
		line-height: 2.5rem;
		font-weight: 800;
		color: #374151;
		letter-spacing: -0.025em;
	}

	.error-text {
		margin-top: 0.25rem;
		font-size: 1rem;
		line-height: 1.5rem;
		color: #6b7280;
	}

	@media (min-width: 640px) {
		.error-title {
			font-size: 3rem;
			line-height: 1;
		}
		.error-container {
			padding-left: 1.5rem;
			padding-right: 1.5rem;
			padding-top: 6rem;
			padding-bottom: 6rem;
		}

		.main {
			display: flex;
		}

		.error-code {
			font-size: 3rem;
			line-height: 1;
		}

		.message-wrapper {
			margin-left: 1.5rem;
		}

		.message-main {
			border-style: solid;
			border-top-width: 0;
			border-right-width: 0;
			border-bottom-width: 0;
			border-left-width: 1px;
			border-color: #e5e7eb;
			padding-left: 1.5rem;
		}
	}

	@media (min-width: 768) {
		.error-container {
			display: grid;
			place-items: center;
		}
	}

	@media (min-width: 1024px) {
		.error-container {
			padding-left: 2rem;
			padding-right: 2rem;
		}
	}
</style>

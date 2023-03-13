<script>
	import { page } from '$app/stores';
</script>

<svelte:head>
	<title>{$page.status}: {$page.error.message}</title>
</svelte:head>

<h1>{$page.status}</h1>
<p>{$page.error.message}</p>

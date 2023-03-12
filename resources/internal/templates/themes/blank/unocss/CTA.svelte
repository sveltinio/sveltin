<script>
	import { sveltinVersion, sveltekitVersion } from '$config/defaults.js';
	import { Link } from '@sveltinio/essentials';
</script>

<h1>Your Sveltin project is up & running!</h1>
<h2>
	You are running Sveltin v{sveltinVersion} on SvelteKit v{sveltekitVersion}
</h2>

<div>
	<Link
		id="quickStartLink"
		external
		icon={false}
		href="https://docs.sveltin.io/quick-start/">Quick Start</Link
	>

	<Link
		id="documentationLink"
		external
		icon={false}
		href="https://docs.sveltin.io">Read the Docs</Link
	>
</div>

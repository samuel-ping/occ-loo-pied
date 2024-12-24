<script lang="ts">
    import { dev } from '$app/environment';

	import { onMount } from 'svelte';

    import Header from './Header.svelte';

	const URL = dev ? 'http://localhost:3333/occupied' : 'http://192.168.0.12:3333/occupied';

	let occupied = $state(false);

    async function getOccupied() {
        try {
			const response = await fetch(URL);
			if (!response.ok) {
				throw new Error(`Response status: ${response.status}`);
			}

			const json = await response.json();
            return json.occupied
		} catch (e) {
			console.error(e);
		}
    }

    let data;
	onMount(async () => {
		data = await getOccupied();
	});
</script>

<Header />

<h1>
	{#if occupied}
		occ-loo-pied
	{:else}
		vacant
	{/if}
</h1>

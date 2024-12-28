<script lang="ts">
	import { dev } from '$app/environment';

	import { onMount } from 'svelte';

	import Status from './Status.svelte';
	import Toggle from './Toggle.svelte';

	const URL: string = dev ? 'http://localhost:3333/occupied' : 'http://192.168.0.12:3333/occupied';

	let occupied: boolean = $state(false);

	interface getOccupiedResponse {
		occupied: boolean;
		occupiedStartTime?: string;
	}

	async function getOccupied(): Promise<getOccupiedResponse> {
		const res = await fetch(URL);
		if (!res.ok) {
			throw new Error(`Response status: ${res.status}`);
		}

		const json = await res.json();
		return {
			occupied: json.occupied,
			occupiedStartTime: json.occupiedStartTime
		};
	}

	async function toggleOccupied(): Promise<void> {
		await fetch(URL, {
			method: 'PUT',
			body: JSON.stringify({ occupied: !occupied })
		});
	}

	let data;
	onMount(async () => {
		data = await getOccupied();
		occupied = data.occupied;
	});
</script>

<Status {occupied} />
<Toggle
	checked={occupied}
	onToggle={() => {
		toggleOccupied();
		occupied = !occupied;
	}}
/>

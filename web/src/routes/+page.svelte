<script lang="ts">
	import { onMount } from 'svelte';

	import Status from './Status.svelte';
	import Toggle from './Toggle.svelte';
	import * as Utils from '$lib/utils';

	let occupied: boolean = $state(false);
	let occupiedStartTime: Date = $state(new Date());

	async function toggleOccupied() {
		let data = await Utils.toggleOccupied(occupied);
		occupied = data.occupied;
		occupiedStartTime = data.occupiedStartTime ? data.occupiedStartTime : new Date();
	}

	onMount(async () => {
		let data = await Utils.getOccupied();
		occupied = data.occupied;
		occupiedStartTime = data.occupiedStartTime ? data.occupiedStartTime : new Date();
	});
</script>

<Status {occupied} {occupiedStartTime} />
<Toggle checked={occupied} onToggle={toggleOccupied} />

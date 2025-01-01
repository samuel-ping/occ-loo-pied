<script lang="ts">
	import { onMount } from 'svelte';

	import Status from './Status.svelte';
	import Toggle from './Toggle.svelte';
	import * as Utils from '$lib/utils';
	import { ONE_SECOND } from '$lib/constants';

	let isLoadingInitialState: boolean = $state(true);
	let occupied: boolean = $state(false);
	let occupiedStartTime: Date = $state(new Date());

	async function toggleOccupied() {
		let data = await Utils.toggleOccupied(occupied);
		occupied = data.occupied;
		occupiedStartTime = data.occupiedStartTime ? data.occupiedStartTime : new Date();
	}

	onMount(() => {
		// Poll server every second
		setInterval(async () => {
			let data = await Utils.getOccupied();
			if (occupied != data.occupied || occupiedStartTime.getTime() != data.occupiedStartTime?.getTime()) {
				occupied = data.occupied;
				occupiedStartTime = data.occupiedStartTime ? data.occupiedStartTime : new Date();
			}
			// Show loading state once on initial page load
			if (isLoadingInitialState) {
				isLoadingInitialState = false;
			}
		}, ONE_SECOND);
	});
</script>

<div
	class="flex h-full w-full flex-col items-center justify-center {occupied
		? 'bg-red-400'
		: 'bg-emerald-500'} transition-all"
>
	{#if isLoadingInitialState}
		<span class="text-3xl">Checking bathroom...</span>
	{:else}
		<div class="flex h-full w-full flex-col items-center justify-center gap-2">
			<Status {occupied} {occupiedStartTime} />
			<Toggle checked={occupied} onToggle={toggleOccupied} />
		</div>
	{/if}
</div>

<script lang="ts">
	import { onMount } from 'svelte';
	
	import Toggle from '$lib/Toggle.svelte';
	import * as Utils from '$lib/utils';
	import Table from './Table.svelte';
	import { ONE_SECOND } from '$lib/constants';

	let occupied: boolean = $state(false);
	let metrics: Utils.metric[] = $state([]);

	async function toggleOccupied() {
		let data = await Utils.toggleOccupied(occupied);
		occupied = data.occupied;
	}

	onMount(async () => {
		let data = await Utils.getOccupied();
		if (occupied != data.occupied) {
			occupied = data.occupied;
		}
	});

	onMount(() => {
		// Poll server every second
		setInterval(async () => {
			let data = await Utils.getOccupied();
			if (occupied != data.occupied) {
				occupied = data.occupied;
			}
		}, ONE_SECOND);
	});

	onMount(async () => {
		let data = await Utils.getMetrics();
		metrics = data.metrics;
	});

	async function onDeleteMetric(id: string) {
		await Utils.deleteMetric(id);

		// refresh state
		let data = await Utils.getMetrics();
		metrics = data.metrics;
	}
</script>

<div>
	<span>
		Toggle override: <Toggle disabled={false} checked={occupied} onToggle={toggleOccupied} />
	</span>
	<Table {metrics} onDelete={onDeleteMetric}/>
</div>

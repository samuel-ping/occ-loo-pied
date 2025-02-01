<script lang="ts">
	import { onMount } from 'svelte';

	import Toggle from '$lib/Toggle.svelte';
	import * as Utils from '$lib/utils';
	import Table from './Table.svelte';
	import { ONE_SECOND } from '$lib/constants';

	let occupied: boolean = $state(false);
	let metrics: Utils.metric[] = $state([]);
	let pagination: Utils.pagination = $state(new Utils.pagination());
	let page: number = $state(1);
	let itemsPerPage: number = $state(10);

	async function toggleOccupied() {
		let occupiedDate = await Utils.toggleOccupied(occupied);
		occupied = occupiedDate.occupied;

		let metricsData = await Utils.getMetrics(page, itemsPerPage);
		metrics = metricsData.metrics;
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
		let data = await Utils.getMetrics(page, itemsPerPage);
		metrics = data.metrics;
		pagination = data.pagination;
	});

	async function deleteMetric(id: string) {
		await Utils.deleteMetric(id);

		// refresh state
		let data = await Utils.getMetrics(page, itemsPerPage);
		metrics = data.metrics;
		pagination = data.pagination;
	}

	async function changePage(changeTo: number) {
		page = changeTo
		let data = await Utils.getMetrics(changeTo, itemsPerPage);
		metrics = data.metrics;
		pagination = data.pagination;
	}
</script>

<div>
	<span>
		Toggle override: <Toggle disabled={false} checked={occupied} onToggle={toggleOccupied} />
	</span>
	<Table
		{metrics}
		onDelete={deleteMetric}
		onPageChange={changePage}
		nextPage={pagination.nextPage}
		prevPage={pagination.prevPage}
	/>
</div>

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
	let itemsPerPage: number = $state(20);
	let itemRangeLower: number = $derived((page - 1) * itemsPerPage + 1);
	let itemRangeHigher: number = $derived(
		Math.min(itemRangeLower + itemsPerPage, pagination.totalItems)
	);

	async function toggleOccupied() {
		let occupiedData = await Utils.toggleOccupied(occupied);
		occupied = occupiedData.occupied;

		let metricsData = await Utils.getMetrics(page, itemsPerPage);
		metrics = metricsData.metrics;
		pagination = metricsData.pagination;
	}

	onMount(async () => {
		let occupiedData = await Utils.getOccupied();
		if (occupied != occupiedData.occupied) {
			occupied = occupiedData.occupied;

			let metricsData = await Utils.getMetrics(page, itemsPerPage);
			metrics = metricsData.metrics;
			pagination = metricsData.pagination;
		}
	});

	onMount(() => {
		// Poll server every second
		setInterval(async () => {
			let occupiedData = await Utils.getOccupied();
			if (occupied != occupiedData.occupied) {
				occupied = occupiedData.occupied;

				let metricsData = await Utils.getMetrics(page, itemsPerPage);
				metrics = metricsData.metrics;
				pagination = metricsData.pagination;
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
		page = changeTo;
		let data = await Utils.getMetrics(changeTo, itemsPerPage);
		metrics = data.metrics;
		pagination = data.pagination;
	}
</script>

<div class="flex flex-col h-full w-full gap-4">
	<h1 class="font-bold text-xl">Admin</h1>
	<span>
		Toggle override: <Toggle disabled={false} checked={occupied} onToggle={toggleOccupied} />
	</span>
	<Table
		{metrics}
		onDelete={deleteMetric}
		onPageChange={changePage}
		totalItems={pagination.totalItems}
		{itemRangeLower}
		{itemRangeHigher}
		nextPage={pagination.nextPage}
		prevPage={pagination.prevPage}
	/>
</div>

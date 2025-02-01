<script lang="ts">
	import { ChevronLeft, ChevronRight, Trash } from 'lucide-svelte';

	import TimeSinceDisplay from '$lib/TimeSinceDisplay.svelte';
	import * as Utils from '$lib/utils';
	import TableButton from './TableButton.svelte';

	interface Props {
		metrics: Utils.metric[];
		onDelete: (id: string) => Promise<void>;
		onPageChange: (changeTo: number) => Promise<void>;
		totalItems: number;
		itemRangeLower: number;
		itemRangeHigher: number;
		nextPage?: number;
		prevPage?: number;
	}

	let {
		metrics,
		onDelete,
		onPageChange,
		totalItems,
		itemRangeLower,
		itemRangeHigher,
		nextPage,
		prevPage
	}: Props = $props();
</script>

<div class="flex flex-col items-center justify-center">
	<table class="size-full table-auto overflow-scroll border border-gray-400">
		<caption
			>Metrics Management <i>(viewing {itemRangeLower}-{itemRangeHigher} out of {totalItems})</i
			></caption
		>
		<tbody>
			<tr>
				<th>ID</th>
				<th>Start Time</th>
				<th>End Time</th>
				<th>Duration</th>
				<th>Delete</th>
			</tr>

			{#each metrics as metric}
				<tr class="border border-gray-400 even:bg-gray-100">
					<td>{metric.id}</td>
					<td>{metric.startTime.toLocaleString()}</td>
					<td>{metric.endTime.toLocaleString()}</td>
					<td><TimeSinceDisplay timeSince={metric.duration} showMilliseconds /></td>
					<td>
						<button
							onclick={async () => {
								onDelete(metric.id);
							}}><Trash /></button
						>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
	<div class="flex flex-row gap-4">
		<TableButton
			onClick={async () => {
				if (prevPage) onPageChange(prevPage);
			}}
			disabled={prevPage == null}><ChevronLeft />previous</TableButton
		>
		<TableButton
			onClick={async () => {
				if (nextPage) onPageChange(nextPage);
			}}
			disabled={nextPage == null}>next<ChevronRight /></TableButton
		>
	</div>
</div>

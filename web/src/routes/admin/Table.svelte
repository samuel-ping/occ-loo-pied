<script lang="ts">
	import { ChevronLeft, ChevronRight, Eraser, RefreshCw, Trash } from 'lucide-svelte';

	import TimeSinceDisplay from '$lib/TimeSinceDisplay.svelte';
	import * as Utils from '$lib/utils';
	import TableButton from './TableButton.svelte';

	let confirmDeleteId: string = $state('');
	let confirmClearEndTimeId: string = $state('');

	interface Props {
		metrics: Utils.metric[];
		onRefresh: () => Promise<void>;
		onClearEndTimeAndDuration: (id: string) => Promise<void>;
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
		onRefresh,
		onClearEndTimeAndDuration,
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
			><b>Metrics Management</b>
			<i>(viewing {itemRangeLower}-{itemRangeHigher} out of {totalItems})</i><button
				class="transition-transform hover:rotate-180"
				onclick={async () => {
					onRefresh();
				}}><RefreshCw /></button
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
					<td
						><span
							>{metric.endTime.toLocaleString()}{#if confirmClearEndTimeId === metric.id}
								<button
									class="text-red-500"
									onclick={async () => {
										onClearEndTimeAndDuration(metric.id);
										confirmClearEndTimeId = '';
									}}><span><Eraser /> Confirm Clear End Time And Duration</span></button
								>
							{:else}
								<button
									onclick={() => {
										confirmClearEndTimeId = metric.id;
									}}><Eraser /></button
								>
							{/if}</span
						></td
					>
					<td
						><TimeSinceDisplay
							timeSince={Utils.nanosecondsToTimeSince(metric.duration)}
							showMilliseconds
						/></td
					>
					<td>
						{#if confirmDeleteId === metric.id}
							<button
								class="text-red-500"
								onclick={async () => {
									onDelete(metric.id);
								}}><span><Trash /> Confirm Delete</span></button
							>
						{:else}
							<button
								onclick={() => {
									confirmDeleteId = metric.id;
								}}><Trash /></button
							>
						{/if}
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

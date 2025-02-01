<script lang="ts">
	import { Trash } from 'lucide-svelte';

	import TimeSinceDisplay from '$lib/TimeSinceDisplay.svelte';
	import * as Utils from '$lib/utils';

	interface Props {
		metrics: Utils.metric[];
		onDelete: (id: string) => Promise<void>;
		onPageChange: (changeTo: number) => Promise<void>;
		nextPage?: number;
		prevPage?: number;
	}

	let { metrics, onDelete, onPageChange, nextPage, prevPage }: Props = $props();
</script>

<div class="flex flex-col items-center justify-center">
	<table class="size-full table-auto border border-gray-400">
		<tbody>
			<tr>
				<th>ID</th>
				<th>Start Time</th>
				<th>End Time</th>
				<th>Duration</th>
				<th>Delete</th>
			</tr>

			{#each metrics as metric}
				<tr>
					<td>{metric.id}</td>
					<td>{metric.startTime}</td>
					<td>{metric.endTime}</td>
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
		<button
			onclick={async () => {
				if (prevPage) onPageChange(prevPage);
			}}
			disabled={prevPage == null}>&lt; previous</button
		>
		<button
			onclick={async () => {
				if (nextPage) onPageChange(nextPage);
			}}
			disabled={nextPage == null}>next &gt;</button
		>
	</div>
</div>

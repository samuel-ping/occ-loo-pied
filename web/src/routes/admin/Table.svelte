<script lang="ts">
	import TimeSinceDisplay from '$lib/TimeSinceDisplay.svelte';
	import * as Utils from '$lib/utils';

	interface Props {
		metrics: Utils.metric[];
		onDelete: (id: string) => Promise<void>;
	}

	let { metrics, onDelete }: Props = $props();
</script>

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
						}}>delete</button
					>
				</td>
			</tr>
		{/each}
	</tbody>
</table>

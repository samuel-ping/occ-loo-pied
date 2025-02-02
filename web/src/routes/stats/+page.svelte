<script lang="ts">
	import * as Utils from '$lib/utils';
	import UsagesByDayHeatmap from './UsagesByDayHeatmap.svelte';
	import TimeSinceDisplay from '$lib/TimeSinceDisplay.svelte';
</script>

<div class="flex h-full w-full flex-col items-center gap-4">
	<h1 class="text-xl font-bold">Bathroom Stats</h1>

	{#await Utils.getStats()}
		<h2>Loading stats...</h2>
	{:then stats}
		<h2>Total Usages: {stats.totalUsages}</h2>

		<h2>
			Total time spent in bathroom is <TimeSinceDisplay
				timeSince={Utils.nanosecondsToTimeSince(stats.duration.total)}
				showMilliseconds={false}
			/>
		</h2>

		<h2>
			Longest time spent in bathroom is <TimeSinceDisplay
				timeSince={Utils.nanosecondsToTimeSince(stats.duration.longest.duration)}
				showMilliseconds={false}
			/>on {stats.duration.longest.startTime.toLocaleDateString()}
		</h2>

		<h2>
			Average time spent in bathroom is <TimeSinceDisplay
				timeSince={Utils.nanosecondsToTimeSince(stats.duration.average)}
				showMilliseconds={true}
			/>
		</h2>
	{:catch error}
		<p>Something went wrong: {error.message}</p>
	{/await}

	<UsagesByDayHeatmap />
</div>

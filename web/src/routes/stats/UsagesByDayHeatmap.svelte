<script lang="ts">
	// @ts-ignore
	import CalHeatmap, { DataRecord } from 'cal-heatmap';
	// @ts-ignore
	import Legend from 'cal-heatmap/plugins/Legend';

	import { onMount } from 'svelte';

	import * as Utils from '$lib/utils';

	import 'cal-heatmap/cal-heatmap.css';

	onMount(async () => {
		let usagesByDayData = await Utils.usagesByDay();
		let usagesByDay = usagesByDayData.usagesByDay;
		let usagesByDayRecord: DataRecord[] = usagesByDay.map((r) => ({
			date: r.date,
			value: r.timesUsed
		}));

		const cal: CalHeatmap = new CalHeatmap();

		cal.paint(
			{
				itemSelector: '#usage-heatmap',
				range: 12,
				domain: { type: 'month' },
				subDomain: {
					type: 'day'
				},
				date: { start: new Date('2025-01-01') },
				data: { source: usagesByDayRecord, x: 'date', y: 'value' },
				scale: {
					color: {
						scheme: 'YlGn',
						type: 'linear',
						domain: [0, 50]
					}
				}
			},
			[
				[
					Legend,
					{
						itemSelector: '#usage-heatmap-legend',
						label: 'Number of Usages'
					}
				]
			]
		);
	});
</script>

<div class="flex flex-col items-center">
	<div id="usage-heatmap"></div>
	<div id="usage-heatmap-legend"></div>
</div>

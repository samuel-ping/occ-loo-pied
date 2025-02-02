<script lang="ts">
	// @ts-expect-error: Types are included by default in the cal-heatmap packages. Not sure why its giving a ts-error.
	import CalHeatmap from 'cal-heatmap';
	// @ts-expect-error
	import Legend from 'cal-heatmap/plugins/Legend';
	// @ts-expect-error
	import Tooltip from 'cal-heatmap/plugins/Tooltip';

	import { onMount } from 'svelte';

	import * as Utils from '$lib/utils';

	import 'cal-heatmap/cal-heatmap.css';

	let usagesByDay: Utils.usageByDayMetric[] = $state([]);
	let mostUsagesInADay: number = $state(0);

	const calPlugins = [
		[
			Legend,
			{
				itemSelector: '#usage-heatmap-legend',
				label: 'Number of Usages'
			}
		],
		[
			Tooltip,
			{
				text: function (
					timestamp: number,
					 value: number,
					 // @ts-expect-error: dayjs type is already included in cal-heatmap package. Not sure why its giving a ts-error.
					 dayjsDate: dayjs.Dayjs
					) {
					let display = `${dayjsDate.format('MMM D')} - ${value} usages`;
					return display;
				}
			}
		]
	];

	const calOptions = (usagesByDay: Utils.usageByDayMetric[], mostUsagesInADay: number) => {
		return {
			itemSelector: '#usage-heatmap',
			range: 12,
			domain: { type: 'month' },
			subDomain: {
				type: 'day'
			},
			date: { 
				start: new Date('2025-01-01'),
				timezone: 'America/New_York',
			},
			data: { source: usagesByDay, x: 'date', y: 'timesUsed' },
			scale: {
				color: {
					scheme: 'YlGn',
					type: 'linear',
					domain: [0, mostUsagesInADay]
				}
			}
		};
	};

	onMount(async () => {
		let usagesByDayData = await Utils.usagesByDay();
		usagesByDay = usagesByDayData.usagesByDay;
		mostUsagesInADay = usagesByDayData.mostUsagesInADay;

		const cal: CalHeatmap = new CalHeatmap();

		cal.paint(calOptions(usagesByDay, mostUsagesInADay), calPlugins);
	});
</script>

<div class="flex flex-col items-center w-full overflow-visible">
	<span class="md:hidden">NOTE: Your screen is small, you probably either need to turn your device to landscape mode or turn on desktop mode to view the full heatmap 🙂</span>
	<div id="usage-heatmap"></div>
	<div id="usage-heatmap-legend"></div>
</div>

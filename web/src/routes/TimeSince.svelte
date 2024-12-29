<script lang="ts">
	import { onMount } from 'svelte';

	import * as Utils from '$lib/utils';

	const INTERVAL_MS = 1000; // 1000ms = 1s

	interface Props {
		occupiedStartTime: Date;
	}

	let { occupiedStartTime }: Props = $props();
	let timeSince: Utils.timeSince = $state({ hours: 0, minutes: 0, seconds: 0 });

	onMount(() => {
		setInterval(() => {
			timeSince = Utils.timeSince(occupiedStartTime);
		}, INTERVAL_MS);
	});
</script>

<span>
	&#40;{#if timeSince.hours != 0}{timeSince.hours}h&nbsp;
	{/if}{#if timeSince.minutes != 0}{timeSince.minutes}m&nbsp;
	{/if}{#if timeSince.seconds != 0}{timeSince.seconds}s&nbsp;{/if}ago&#41;
</span>

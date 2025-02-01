<script lang="ts">
	import * as Utils from '$lib/utils';

	interface Props {
		timeSince: Utils.timeSince;
		showMilliseconds: boolean;
	}

	let { timeSince, showMilliseconds }: Props = $props();
</script>

<span>
		{#if timeSince.hours != 0}{timeSince.hours}h&nbsp;
		{/if}{#if timeSince.minutes != 0}{timeSince.minutes}m&nbsp;
		{/if}{#if timeSince.seconds != 0}{timeSince.seconds}s&nbsp;
		{/if}{#if showMilliseconds && timeSince.milliseconds != 0}{timeSince.milliseconds}ms&nbsp{/if}
</span>

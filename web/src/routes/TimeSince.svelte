<script lang="ts">
	import { onMount } from 'svelte';

	import TimeSinceDisplay from '$lib/TimeSinceDisplay.svelte';
	import * as Utils from '$lib/utils';
	import { ONE_SECOND } from '$lib/constants';

	interface Props {
		occupiedStartTime: Date;
	}

	let { occupiedStartTime }: Props = $props();
	let timeSince: Utils.timeSince = $state({ hours: 0, minutes: 0, seconds: 0, milliseconds: 0 });

	onMount(() => {
		timeSince = Utils.timeSince(occupiedStartTime);
	});

	onMount(() => {
		setInterval(() => {
			timeSince = Utils.timeSince(occupiedStartTime);
		}, ONE_SECOND);
	});
</script>

&#40;<TimeSinceDisplay {timeSince} showMilliseconds={false} />ago&#41;

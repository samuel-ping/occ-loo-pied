<script lang="ts">
	import TimeSince from './TimeSince.svelte';

	interface Props {
		occupied: boolean;
		occupiedStartTime: Date;
	}

	let { occupied, occupiedStartTime }: Props = $props();
</script>

<div class="flex flex-col items-center justify-center">
	<h2 class="text-3xl">bathroom is</h2>

	<h1 class="text-6xl">
		{#if occupied}
			occ-loo-pied 🚽
		{:else}
			vacant 😀
		{/if}
	</h1>

	{#if occupied}
		<span class="flex flex-row text-xl">
			as of {occupiedStartTime.toLocaleTimeString()}&nbsp;
			<TimeSince {occupiedStartTime} />
		</span>
	{:else}
		<br />
	{/if}
</div>

<script lang="ts">
	interface Props {
		occupied: boolean;
	}

	let { occupied }: Props = $props();
</script>

<h2 class="text-2xl">bathroom is</h2>

<h1 class="text-5xl">
	{#if occupied}
		occ-loo-pied
	{:else}
		vacant
	{/if}
</h1>

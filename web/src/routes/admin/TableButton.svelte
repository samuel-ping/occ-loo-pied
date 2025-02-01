<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		disabled: boolean;
		onClick: () => Promise<void>;
		children: Snippet;
	}

	let { disabled, onClick, children }: Props = $props();
</script>

<button class="flex flex-row {disabled ? 'text-gray-300' : ''}" onclick={onClick} {disabled}
	>{@render children()}</button
>

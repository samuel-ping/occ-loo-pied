<script lang="ts">
	interface Props {
		disabled: boolean;
		checked: boolean;
		onToggle: () => void;
	}

	let { disabled = true, checked, onToggle }: Props = $props();
</script>

<!-- Source: https://www.material-tailwind.com/docs/html/switch -->
<div class="relative inline-block h-5 w-11">
	<input
		{checked}
		onchange={onToggle}
		id="switch-component"
		type="checkbox"
		class="peer h-5 w-11 appearance-none rounded-full bg-slate-100 transition-colors duration-300 checked:bg-slate-800 {disabled
			? 'cursor-not-allowed'
			: 'cursor-pointer'}"
		{disabled}
	/>
	<label
		for="switch-component"
		class="absolute left-0 top-0 h-5 w-5 rounded-full border border-slate-300 bg-white shadow-sm transition-transform duration-300 peer-checked:translate-x-6 peer-checked:border-slate-800 {disabled
			? 'cursor-not-allowed'
			: 'cursor-pointer'}"
	>
	</label>
</div>

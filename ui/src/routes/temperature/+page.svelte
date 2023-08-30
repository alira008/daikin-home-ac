<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	/** @type {import ('./$types').PageServerData} */
	// export let data;

	let host = $page.url.hostname;
	let temperatureInput = 76;

	onMount(async () => {
		const response = await fetch(`http://${host}:5520/temperature`);

		const temperature = await response.json();
		temperatureInput = Math.round(temperature * 1.8 + 32);
	});

	async function setTemperature() {
		// Convert Fahrenheit to celsius
		const celsius = Math.round((temperatureInput - 32.0) / 1.8);
		await fetch(`http://${host}:5520/temperature/${celsius}`);
	}
</script>

<label for="temperatureInput" class="label">
	<span>AC Temperature Fahrenheit</span>
	<input
		id="temperatureInput"
		type="number"
		class="input variant-form-material"
		bind:value={temperatureInput}
	/>
</label>
<button class="btn variant-filled-primary" on:click={() => setTemperature()}>
	Set Temperature
</button>

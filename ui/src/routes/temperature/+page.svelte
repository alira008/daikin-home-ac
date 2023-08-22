<script>
	import { page } from '$app/stores';

	// let host = $page.url.hostname;
	let host = '192.168.1.228';
	let temperatureInput = 76;
	async function setTemperature() {
        // Convert Fahrenheit to celsius
		const celsius = (temperatureInput - 32.0) / 1.8;
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

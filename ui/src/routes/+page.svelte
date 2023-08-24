<script>
	import { page } from '$app/stores';

	let host = $page.url.hostname;
	// let host = '192.168.1.228';

	/**
    @param {boolean | number} state 
    @param {string} setting 
    */
	async function setState(setting, state) {
		try {
			await fetch(`http://${host}:5520/${setting}/${state}`);
		} catch (error) {
			console.error(error);
		}
	}
</script>

<div class="flex flex-col items-center gap-4">
	<h3 class="h3">Power</h3>
	<div class="flex gap-2">
		<button class="btn variant-filled-primary" on:click={() => setState('power', true)}>
			On
		</button>
		<button class="btn variant-filled-tertiary" on:click={() => setState('power', false)}>
			Off
		</button>
	</div>
</div>
<div class="flex flex-col items-center gap-4">
	<h3 class="h3">Mode</h3>
	<div class="flex flex-col gap-2">
		<div class="flex gap-2">
			<button class="btn variant-filled-primary" on:click={() => setState('mode', 0)}>
				Auto
			</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('mode', 2)}>
				Dry
			</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('mode', 3)}>
				Cold
			</button>
			<button
				class="btn variant-filled-tertiary hidden md:inline-flex"
				on:click={() => setState('mode', 4)}
			>
				Heat
			</button>
			<button
				class="btn variant-filled-tertiary hidden md:inline-flex"
				on:click={() => setState('mode', 6)}
			>
				Fan
			</button>
		</div>
		<div class="flex gap-2 justify-center md:hidden">
			<button class="btn variant-filled-tertiary" on:click={() => setState('mode', 4)}>
				Heat
			</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('mode', 6)}>
				Fan
			</button>
		</div>
	</div>
</div>
<div class="flex flex-col items-center gap-4">
	<h3 class="h3">Fan Speed</h3>
	<div class="flex flex-col gap-2">
		<div class="flex gap-2 justify-center md:hidden">
			<button class="btn variant-filled-primary" on:click={() => setState('fanspeed', 10)}>Automatic</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('fanspeed', 11)}>Silent</button>
		</div>
		<div class="flex gap-2">
            <button class="btn variant-filled-primary hidden md:inline-flex" on:click={() => setState('fanspeed', 10)}>Automatic</button>
            <button class="btn variant-filled-tertiary hidden md:inline-flex" on:click={() => setState('fanspeed', 11)}>Silent</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('fanspeed', 3)}>Speed 1</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('fanspeed', 4)}>Speed 2</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('fanspeed', 5)}>Speed 3</button>
			<button class="btn variant-filled-tertiary hidden md:inline-flex" on:click={() => setState('fanspeed', 6)}>Speed 4</button>
			<button class="btn variant-filled-tertiary hidden md:inline-flex" on:click={() => setState('fanspeed', 7)}>Speed 5</button>
		</div>
		<div class="flex gap-2 justify-center md:hidden">
			<button class="btn variant-filled-tertiary" on:click={() => setState('fanspeed', 6)}>Speed 4</button>
			<button class="btn variant-filled-tertiary" on:click={() => setState('fanspeed', 7)}>Speed 5</button>
		</div>
	</div>
</div>
<div class="flex flex-col items-center gap-4">
	<h3 class="h3">Powerful Mode</h3>
	<div class="flex gap-2">
		<button class="btn variant-filled-primary" on:click={() => setState('powerful', true)}>
			On
		</button>
		<button class="btn variant-filled-tertiary" on:click={() => setState('powerful', false)}>
			Off
		</button>
	</div>
</div>
<!-- <div class="flex flex-col items-center gap-4"> -->
<!-- 	<h3 class="h3">Swing Mode</h3> -->
<!-- 	<div class="flex gap-2"> -->
<!-- 		<button class="btn variant-filled-primary">On</button> -->
<!-- 		<button class="btn variant-filled-tertiary">Off</button> -->
<!-- 	</div> -->
<!-- </div> -->

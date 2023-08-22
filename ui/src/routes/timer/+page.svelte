<script>
	import { DateTime } from 'luxon';
	let offDatetime = DateTime.now().toFormat('HH:mm');
	let onDatetime = DateTime.now().toFormat('HH:mm');
</script>

<div class="flex gap-4">
	<label for="datetimeOff" class="label">
		<span>Off Time</span>
		<input
			id="datetimeOff"
			type="time"
			class="w-fit input variant-form-material"
			bind:value={offDatetime}
		/>
	</label>
	<label for="datetimeOn" class="label">
		<span>On Time</span>
		<input
			id="datetimeOn"
			type="time"
			class="w-fit input variant-form-material"
			bind:value={onDatetime}
		/>
	</label>
</div>
<div class="flex gap-4">
	<button class="btn variant-filled-primary"> Set Off Timer </button>
	<button class="btn variant-filled-primary"> Set On Timer </button>
</div>

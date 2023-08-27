<script>
	import { page } from '$app/stores';
	import { DateTime } from 'luxon';
	let offDatetime = DateTime.now().toFormat('HH:mm');
	let onDatetime = DateTime.now().toFormat('HH:mm');

	let host = $page.url.hostname;

	/**
	 *
	 * @param {string} timeStr
	 * @returns {number}
	 */
	function calculateTimeDifference(timeStr) {
		const time = DateTime.fromFormat(timeStr, 'HH:mm');
		const currentTime = DateTime.now();

		return time < currentTime
			? time.plus({ hours: 24 }).diff(currentTime, 'minutes').minutes
			: time.diff(currentTime, 'minutes').minutes;
	}

	/**
	 *
	 * @param {'on' | 'off'} timer
	 */
	async function setTimer(timer) {
		/** @type {number} */
		let timeDiff = 0;
		if (timer === 'on') {
			timeDiff = Math.round(calculateTimeDifference(onDatetime));
		} else if (timer === 'off') {
			timeDiff = Math.round(calculateTimeDifference(offDatetime));
		}

		await fetch(`http://${host}:5520/timer/${timer}/true/${timeDiff}`);
	}

	/**
	 *
	 * @param {'on' | 'off'} timer
	 */
	async function clearTimer(timer) {
		await fetch(`http://${host}:5520/timer/${timer}/false/0`);
	}
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
	<button class="btn variant-filled-primary" on:click={() => setTimer('off')}>
		Set Off Timer
	</button>
	<button class="btn variant-filled-primary" on:click={() => setTimer('on')}> Set On Timer </button>
</div>
<div class="flex gap-4">
	<button class="btn variant-filled-primary" on:click={() => clearTimer('off')}>
		Clear Off Timer
	</button>
	<button class="btn variant-filled-primary" on:click={() => clearTimer('on')}> Clear On Timer </button>
</div>

<script>
	import '@skeletonlabs/skeleton/themes/theme-hamlindigo.css';
	import '@skeletonlabs/skeleton/styles/skeleton.css';
	import '../app.postcss';

	import { AppShell, AppBar, Drawer, drawerStore } from '@skeletonlabs/skeleton';
</script>

<Drawer position="top" height="h-fit">
	{#if $drawerStore.id === 'mobile-menu'}
		<div class="flex flex-col w-full gap-4 p-4">
			<a href="/" class="btn variant-filled-primary" on:click={() => drawerStore.close()}> Home </a>
			<a
				href="/temperature"
				class="btn variant-filled-primary"
				on:click={() => drawerStore.close()}
			>
				Temperature
			</a>
			<a href="/timer" class="btn variant-filled-primary" on:click={() => drawerStore.close()}>
				Timer
			</a>
		</div>
	{/if}
</Drawer>

<AppShell>
	<AppBar slotLead="gap-4">
		<svelte:fragment slot="lead">
			<button
				class="btn-icon variant-ringed md:hidden"
				on:click={() => drawerStore.open({ id: 'mobile-menu' })}
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					height="1.5em"
					class="md:hidden"
					viewBox="0 0 448 512"
				>
					<path
						d="M0 96C0 78.3 14.3 64 32 64H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32C14.3 128 0 113.7 0 96zM0 256c0-17.7 14.3-32 32-32H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32zM448 416c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32H416c17.7 0 32 14.3 32 32z"
					/>
				</svg>
			</button>
			<a href="/"class="btn variant-filled-primary hidden md:inline-flex">Home</a>
			<a href="/temperature"class="btn variant-filled-primary hidden md:inline-flex">Temperature</a>
			<a href="/timer"class="btn variant-filled-primary hidden md:inline-flex">Timer</a>
		</svelte:fragment>
		<svelte:fragment slot="trail">Daikin AC Remote</svelte:fragment>
	</AppBar>
	<div class="flex flex-col gap-8 items-center my-8">
		<slot />
	</div>
</AppShell>

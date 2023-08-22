/** @type {import('tailwindcss').Config} */
export default {
	content: [
		'./src/**/*.{html,js,svelte}',
		require('path').join(require.resolve('@skeletonlabs/skeleton'), '../**/*.{html,js,svelte}')
	],
	theme: {
		extend: {}
	},
	plugins: [
		require('@tailwindcss/forms'),
		...require('@skeletonlabs/skeleton/tailwind/skeleton.cjs')()
	]
};

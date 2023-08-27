/** @type {import('./$types').PageServerLoad} */
export async function load({url}) {
    const host = url.hostname
	const response = await fetch(`http://${host}:5520/temperature`);

	const responseJson = await response.json();

	return { temperature: responseJson };
}

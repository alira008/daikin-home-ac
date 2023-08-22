/** @type {import('./$types').PageServerLoad} */
export async function load({ url }) {
	// let host = url.host;
	const host = '192.168.1.228';
	const response = await fetch(`http://${host}:5520/temperature`);

	const responseJson = response.json();

	return responseJson;
}

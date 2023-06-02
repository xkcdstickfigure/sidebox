export const censorAddress = (address) => {
	const username = address.split("@")[0]
	const domain = address.split("@")[1]
	return `${username.substring(0, 4)}****@${domain}`
}

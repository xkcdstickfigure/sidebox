import { createAPIClient } from "../../api.js"

chrome.storage.local.get(null, ({ token }) => {
	const api = createAPIClient(token)
	api.account().then(console.log).catch(console.error)
})

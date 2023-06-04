import { API } from "./api.js"

chrome.runtime.onMessage.addListener((request, sender) => {
	if (request.type === "boxes-auth" && sender.tab) {
		const params = new URL(sender.tab.url).searchParams
		const code = params.get("code")
		const state = params.get("state")

		if (code)
			new API()
				.login(code, state)
				.then(({ token }) => {
					chrome.storage.local.set({ token })
				})
				.catch(console.error)
	}
})

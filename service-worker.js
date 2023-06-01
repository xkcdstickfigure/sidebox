import { API } from "./api.js"

chrome.runtime.onMessage.addListener((request, sender) => {
	if (request.type === "boxes-auth") {
		const code = new URL(sender.tab?.url).searchParams.get("code")

		if (code)
			new API()
				.login(code)
				.then(({ token }) => {
					chrome.storage.local.set({ token })
				})
				.catch(console.error)
	}
})

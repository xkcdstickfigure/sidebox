import { api } from "../api.js"

chrome.runtime.onMessage.addListener((request, sender) => {
	if (request.type === "boxes-auth") {
		const code = new URL(sender.tab?.url).searchParams.get("code")

		if (code)
			api
				.login(code)
				.then(({ token }) => {
					console.log(token)
				})
				.catch(console.error)
	}
})

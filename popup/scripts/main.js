import { createAPIClient } from "../../api.js"

// screens
const setScreen = (name) => {
	document.querySelectorAll(".screen").forEach((elem) => {
		elem.style.display = elem.dataset.screen === name ? "block" : "none"
	})
}
setScreen("loading")

// retreive data from storage
chrome.storage.local.get(null, ({ token, accountCache }) => {
	// launch auth flow if no token
	if (!token) return launchAuthFlow()

	// create api client
	const api = createAPIClient(token)

	// render page from account cache
	if (accountCache) setScreen("home")

	// fetch account data
	api
		.account()
		.then((data) => {
			// store account data
			chrome.storage.local.set({ accountCache: data })

			setScreen("home")
		})
		.catch((err) => {
			// launch auth flow if bad token
			if (err === "bad authorization") return launchAuthFlow()
		})
})

// reauthenticate
const launchAuthFlow = () => {
	window.open("http://localhost:3000/auth")
}

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
	if (accountCache) showMessage(accountCache.email)

	// fetch account data
	api
		.account()
		.then((data) => {
			// store account data
			chrome.storage.local.set({ accountCache: data })

			showMessage(data.email)
		})
		.catch((err) => {
			// launch auth flow if bad token
			if (err === "bad authorization") return launchAuthFlow()

			showMessage("unable to connect to server")
		})
})

// reauthenticate
const launchAuthFlow = () => {
	window.open("http://localhost:3000/auth")
}

// show message
const showMessage = (text) => {
	//document.querySelector(".message").innerText = text
}

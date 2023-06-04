import { API } from "../../api.js"
import { setScreen } from "./util/screen.js"
import { renderHomeScreen } from "./screens/home.js"

setScreen("loading")

// retreive data from storage
chrome.storage.local.get(null, ({ token, accountCache }) => {
	// launch auth flow if no token
	if (!token) return launchAuthFlow()

	// create api client
	const api = new API(token)

	// render home screen from account cache
	if (accountCache) {
		renderHomeScreen(api, accountCache)
		setScreen("home")
	}

	// fetch account data
	api
		.account()
		.then((data) => {
			// open update url
			if (data.updateUrl) window.open(data.updateUrl)

			// store account data
			chrome.storage.local.set({ accountCache: data })

			// render home screen
			renderHomeScreen(api, data)
			if (!accountCache) setScreen("home")
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

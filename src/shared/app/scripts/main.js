import { API, server } from "../../api.js"
import { setScreen } from "./util/screen.js"
import { renderHomeScreen, renderInboxList } from "./screens/home.js"

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
		.then((account) => {
			// open update url
			if (account.updateUrl) {
				chrome.tabs.create({ url: account.updateUrl })
				return window.close()
			}

			// cache account data
			chrome.storage.local.set({ accountCache: account })

			// render home screen
			renderHomeScreen(api, account)
			if (!accountCache) setScreen("home")

			// account refetch interval
			startAccountRefetchInterval(api)
		})
		.catch((err) => {
			// launch auth flow if bad token
			if (err === "bad authorization") return launchAuthFlow()
		})
})

// reauthenticate
const launchAuthFlow = () => {
	chrome.tabs.create({ url: `${server}/auth` })
	window.close()
}

// account refetch interval
const startAccountRefetchInterval = (api) =>
	setInterval(() => {
		api.account().then((account) => {
			// cache account data
			chrome.storage.local.set({ accountCache: account })

			// render inbox list
			renderInboxList(api, account.inboxes)
		})
	}, 5000)

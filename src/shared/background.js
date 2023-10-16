import { API } from "../api.js"

chrome.runtime.onMessage.addListener((request, sender) => {
	if (request.type === "sidebox-auth" && sender.tab) {
		const params = new URL(sender.tab.url).searchParams
		const code = params.get("code")
		const state = params.get("state")
		
		if (code)
		new API().login(code, state).then(({ token }) => {
			chrome.storage.local.set({ token })
		})
	}
})

chrome.runtime.onInstalled.addListener(() => {
	chrome.contextMenus.removeAll();
	if (chrome.sidePanel) {
		chrome.contextMenus.create({
			id: "open-sidebar",
			title: "Open in sidebar",
			contexts: ["action"]
		})
	}
})

if (chrome.sidePanel) {
	chrome.contextMenus.onClicked.addListener((info, tab) => {
		if (info.menuItemId === "open-sidebar") {
			chrome.sidePanel.open({ windowId: tab.windowId })
		}
	})
}

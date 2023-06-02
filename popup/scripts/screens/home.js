import { getColor } from "../util/color.js"
import { setScreen } from "../util/screen.js"
import { renderInboxScreen } from "./inbox.js"
import { censorAddress } from "../util/address.js"

const inboxes = document.querySelector(".homeScreen .inboxes")

export const renderHomeScreen = (api, account) => {
	inboxes.innerHTML = ""
	inboxes.append(...account.inboxes.map((inbox) => createInbox(api, inbox)))
}

const createInbox = (api, { id, name, address, unread }) => {
	const inbox = document.createElement("button")
	inbox.className = "inbox"
	inbox.dataset.inboxId = id

	inbox.onclick = () => {
		renderInboxScreen(api, { id, name, address })
		setScreen("inbox")
	}

	const icon = document.createElement("div")
	icon.className = "icon"
	icon.style.backgroundColor = getColor(id)
	icon.innerText = name[0].toUpperCase()

	const info = document.createElement("div")
	info.className = "info"

	const nameText = document.createElement("p")
	nameText.className = "name"
	nameText.innerText = name

	const addressText = document.createElement("p")
	addressText.className = "address"
	addressText.innerText = censorAddress(address)

	if (unread) {
		const unreadDot = document.createElement("div")
		unreadDot.className = "dot"
		inbox.append(unreadDot)
	}

	info.append(nameText, addressText)
	inbox.append(icon, info)
	return inbox
}

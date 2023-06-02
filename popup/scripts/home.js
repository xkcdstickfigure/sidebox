import { getColor } from "./color.js"
import { setScreen } from "./screen.js"
import { renderInboxScreen } from "./inbox.js"

const inboxes = document.querySelector(".homeScreen .inboxes")

export const renderHomeScreen = (account) => {
	inboxes.innerHTML = ""
	inboxes.append(
		...account.inboxes.map((inbox) =>
			createInbox(inbox.id, inbox.name, inbox.address, inbox.unread)
		)
	)
}

const createInbox = (id, name, address, unread) => {
	const inbox = document.createElement("button")
	inbox.className = "inbox"

	inbox.onclick = () => {
		renderInboxScreen({ id, name, address })
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

const censorAddress = (address) => {
	const username = address.split("@")[0]
	const domain = address.split("@")[1]
	return `${username.substring(0, 4)}****@${domain}`
}

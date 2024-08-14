import { setScreen } from "../util/screen.js"
import { renderInboxScreen } from "./inbox.js"
import { censorAddress } from "../util/address.js"
import { getIconText, getIconColor } from "../util/icon.js"

const screen = document.querySelector(".homeScreen")
const $ = (str) => screen.querySelector(str)

// create button
$(".create").onclick = () => {
	$(".createForm").style.display = "block"
	$(".createForm input").focus()
}

// cancel create button
$(".createForm .cancel").onclick = () => {
	$(".createForm").style.display = "none"
}

// render screen
export const renderHomeScreen = (api, account) => {
	// inbox list
	renderInboxList(api, account.inboxes)

	// search
	$(".search").oninput = () => {
		renderInboxList(api, account.inboxes)
	}

	// create inbox
	$(".createForm").onsubmit = (e) => {
		e.preventDefault()

		const name = $(".createForm input").value.trim()
		if (!name) return

		$(".createForm .confirm").disabled = true

		api
			.inboxCreate(name)
			.then((data) => {
				// move to inbox screen
				renderInboxScreen(api, data, true)
				setScreen("inbox")

				// add inbox to list
				$(".inboxes").prepend(createInbox(api, data))

				// close form
				$(".createForm").reset()
				$(".createForm").style.display = "none"
			})
			.finally(() => {
				$(".createForm .confirm").disabled = false
			})
	}
}

// render inbox list
export const renderInboxList = (api, inboxes) => {
	const value = $(".search").value.trim()
	const filteredInboxes = value
		? inboxes.filter((inbox) => inbox.name.includes(value))
		: inboxes

	$(".inboxes").innerHTML = ""
	$(".inboxes").append(
		...filteredInboxes.map((inbox) => createInbox(api, inbox)),
	)
}

// create inbox list row
const createInbox = (api, data) => {
	const { id, name, address, unread } = data
	const inbox = document.createElement("button")
	inbox.className = "inbox"
	inbox.dataset.inboxId = id
	inbox.dataset.inbox = JSON.stringify(data)

	inbox.onclick = () => {
		// display inbox as read
		inbox.classList.remove("unread")
		inbox.querySelector(".dot")?.remove()

		// render inbox screen
		renderInboxScreen(api, JSON.parse(inbox.dataset.inbox), false)
		setScreen("inbox")
	}

	const icon = document.createElement("div")
	icon.className = "icon"
	icon.style.backgroundColor = getIconColor(id)
	icon.innerText = getIconText(name)

	const info = document.createElement("div")

	const nameText = document.createElement("p")
	nameText.className = "name"
	nameText.innerText = name

	const addressText = document.createElement("p")
	addressText.className = "address"
	addressText.innerText = censorAddress(address)

	if (unread) {
		inbox.classList.add("unread")
		const unreadDot = document.createElement("div")
		unreadDot.className = "dot"
		inbox.append(unreadDot)
	}

	info.append(nameText, addressText)
	inbox.append(icon, info)
	return inbox
}

// type in search input
window.addEventListener("keydown", (event) => {
	const isTyping = event.key.length === 1 || event.key === "Backspace"
	const isScreenActive = screen.style.display === "block"
	const isPopupActive = $(".createForm").style.display === "block"

	if (isTyping && isScreenActive && !isPopupActive) {
		$(".search").focus()
	}
})

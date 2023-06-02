import { censorAddress } from "../util/address.js"

const screen = document.querySelector(".inboxScreen")

export const renderInboxScreen = ({ id, name, address }) => {
	screen.querySelector(".name").innerText = name
	screen.querySelector(".address").innerText = censorAddress(address)
}

import { censorAddress } from "../util/address.js"

const screen = document.querySelector(".inboxScreen")

export const renderInboxScreen = ({ id, name, address }) => {
	// name
	screen.querySelector(".name").innerText = name

	// address
	const addressButton = document.createElement("button")
	addressButton.className = "address"
	addressButton.innerText = censorAddress(address)
	screen.querySelector(".address").replaceWith(addressButton)

	// click to reveal full address
	addressButton.onclick = () => {
		const revealedAddressText = document.createElement("p")
		revealedAddressText.className = "address"
		revealedAddressText.innerText = address
		addressButton.replaceWith(revealedAddressText)
	}
}

const screen = document.querySelector(".inboxScreen")

export const renderInboxScreen = ({ id, name, address }) => {
	screen.querySelector(".title").innerText = name
}

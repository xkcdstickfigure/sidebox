export const setScreen = (name) => {
	document.querySelectorAll(".screen").forEach((elem) => {
		elem.style.display = elem.dataset.screen === name ? "block" : "none"
	})
}

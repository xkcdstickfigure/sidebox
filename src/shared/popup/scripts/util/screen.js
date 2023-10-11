export const setScreen = (name) => {
	document.body.dataset.screen = name;
	document.body.querySelectorAll("[data-screen]").forEach((elem) => {
		elem.classList[(elem.dataset.screen === name ? "add" : "remove")]("activeScreen");
	})
}

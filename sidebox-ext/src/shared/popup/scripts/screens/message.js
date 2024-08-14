import { setScreen } from "../util/screen.js"

const screen = document.querySelector(".messageScreen")
const $ = (str) => screen.querySelector(str)

// back
$(".back").onclick = () => setScreen("inbox")

// render screen
export const renderMessageScreen = (
	api,
	{ id, fromName, fromAddress, subject, date }
) => {
	$(".subject").innerText = subject
	$(".from").innerText = `${fromName} (${fromAddress})`
	$(".date").innerText = Intl.DateTimeFormat(undefined, {
		dateStyle: "medium",
		timeStyle: "short",
	}).format(new Date(date))
	$(".content").style.display = "none"
	$(".content").srcdoc = ""

	api.messageGet(id).then((data) => {
		$(".content").srcdoc = htmlContentPrefix + data.body
		$(".content").style.display = "block"
	})
}

// html content prefix
const htmlContentPrefix = `<style>
	* {
		font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont,
			Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans, sans-serif,
			Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol, Noto Color Emoji;
	}
</style>
<base target="_blank" />`

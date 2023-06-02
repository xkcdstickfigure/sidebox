import { setScreen } from "../util/screen.js"

const screen = document.querySelector(".messageScreen")
const $ = (str) => screen.querySelector(str)

// back
$(".back").onclick = () => setScreen("inbox")

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
	$(".content").srcdoc = ""

	api.messageGet(id).then((data) => {
		$(".content").srcdoc = data.body
	})
}

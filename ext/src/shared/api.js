export const server = "{{SERVER}}"

export class API {
	#token

	constructor(token) {
		this.#token = token
	}

	#fetch = (method, url, body) =>
		new Promise((resolve, reject) =>
			fetch(`${server}/api/${url}`, {
				method,
				headers: {
					authorization: this.#token,
					"content-type": "application/json",
				},
				body: body ? JSON.stringify(body) : undefined,
			})
				.then((res) => {
					if (res.status === 200) res.json().then((data) => resolve(data))
					else
						res
							.json()
							.then((err) => reject(err.errorName))
							.catch(reject)
				})
				.catch(reject)
		)

	login = (code, state) => this.#fetch("POST", "login", { code, state })
	account = () => this.#fetch("GET", "account")
	inboxCreate = (name) => this.#fetch("POST", "inbox", { name })
	inboxGet = (id) => this.#fetch("GET", `inbox/${encodeURIComponent(id)}`)
	inboxSetName = (id, name) =>
		this.#fetch("POST", `inbox/${encodeURIComponent(id)}/name`, { name })
	inboxSetMuted = (id, muted) =>
		this.#fetch("POST", `inbox/${encodeURIComponent(id)}/muted`, { muted })
	inboxDelete = (id) => this.#fetch("DELETE", `inbox/${encodeURIComponent(id)}`)
	messageGet = (id) => this.#fetch("GET", `message/${encodeURIComponent(id)}`)
}

const server = "http://localhost:3000"

export const createAPIClient = (token) => {
	const apiFetch = (method, url, body) =>
		new Promise((resolve, reject) =>
			fetch(`${server}/api/${url}`, {
				method,
				headers: {
					authorization: token,
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
							.catch(resolve)
				})
				.catch(reject)
		)

	return {
		login: (code) => apiFetch("POST", "login", { code }),
		account: () => apiFetch("GET", "account"),
		inboxCreate: (name) => apiFetch("POST", "inbox", { name }),
		inboxGet: (id) => apiFetch("GET", `inbox/${encodeURIComponent(id)}`),
		inboxDelete: (id) => apiFetch("DELETE", `inbox/${encodeURIComponent(id)}`),
		messageGet: (id) => apiFetch("GET", `message/${encodeURIComponent(id)}`),
	}
}

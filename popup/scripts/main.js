chrome.storage.local.get(null, (storage) => {
	document.write(JSON.stringify(storage))
})

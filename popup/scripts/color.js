export const getColor = (id) =>
	colors[parseInt(id.split("-")[0], 16) % colors.length]

const colors = [
	"#1abc9c",
	"#2ecc71",
	"#3498db",
	"#9b59b6",
	"#e91e63",
	"#f1c40f",
	"#e67e22",
	"#e74c3c",
	"#607d8b",
	"#11806a",
	"#1f8b4c",
	"#206694",
	"#71368a",
	"#ad1457",
	"#c27c0e",
	"#a84300",
	"#992d22",
]

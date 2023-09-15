package response

type Error struct {
	Number  int    `json:"code"`
	Message string `json:"message"`
}

package response

type Token struct {
	Token   string `json:"token"`
	Created string `json:"created"`
	Expired string `json:"expired"`
}

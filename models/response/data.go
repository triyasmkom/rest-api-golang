package response

type Data struct {
	Iat      string `json:"iat,omitempty"`
	Exp      string `json:"exp,omitempty"`
	Email    string `json:"email,omitempty"`
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
	Created  string `json:"created,omitempty"`
	Expired  string `json:"expired,omitempty"`
	Password string `json:"password,omitempty"`
}

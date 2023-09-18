package response

type Data struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	Iat     string `json:"iat,omitempty"`
	Exp     string `json:"exp,omitempty"`
	Token   string `json:"token,omitempty"`
	Created string `json:"created,omitempty"`
	Expired string `json:"expired,omitempty"`

	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`

	Alamat    string `json:"alamat,omitempty"`
	Kelurahan string `json:"kelurahan,omitempty"`
	Kecamatan string `json:"kecamatan,omitempty"`
	Kabupaten string `json:"kabupaten,omitempty"`
	Provinsi  string `json:"provinsi,omitempty"`
}

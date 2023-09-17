package response

import "server-golang/models/database"

type Claims struct {
	Iat      float64         `json:"iat,omitempty"`
	Exp      float64         `json:"exp,omitempty"`
	Email    string          `json:"email,omitempty"`
	Id       uint            `json:"id,omitempty"`
	Username string          `json:"username,omitempty"`
	Role     []database.Role `json:"role"`
}

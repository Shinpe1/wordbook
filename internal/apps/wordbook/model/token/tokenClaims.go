package token

type TokenClaims struct {
	UserId int    `json:"userId"`
	Iss    string `json:"issuer"`
	Iat    string `json:"iat"`
}

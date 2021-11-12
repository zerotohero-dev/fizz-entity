package endpoint

var Crypto = struct {
	SecureHash       string
	SecureHashVerify string
	Jwt              string
	JwtVerify        string
	SecureToken      string
}{
	SecureHash:       "/v1/hash",
	SecureHashVerify: "/v1/hash/verify",
	Jwt:              "/v1/jwt",
	JwtVerify:        "/v1/jwt/verify",
	SecureToken:      "/v1/token",
}

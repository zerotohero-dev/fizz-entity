/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

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

var Mailer = struct {
	Verification         string
	Welcome              string
	PasswordReset        string
	PasswordResetConfirm string
	SubscriptionConfirm  string
}{
	Verification:         "/v1/relay/verification",
	Welcome:              "/v1/relay/welcome",
	PasswordReset:        "/v1/relay/reset",
	PasswordResetConfirm: "/v1/relay/confirm",
	SubscriptionConfirm:  "/v1/relay/subscribed",
}

var Idm = struct {
	Info           string
	Login          string
	SignUp         string
	CreateAccount  string
	RemindPassword string
	ResetPassword  string
}{
	Info:           "/v1/info",
	Login:          "/v1/login",
	SignUp:         "/v1/signup",
	CreateAccount:  "/v1/create",
	RemindPassword: "/v1/remind",
	ResetPassword:  "/v1/reset",
}

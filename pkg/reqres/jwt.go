/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package reqres

type JwtCreateRequest struct {
	Email   string   `json:"email"`
	Err     string   `json:"err,omitempty"`
}

type JwtCreateResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type JwtVerifyRequest struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// JwtVerifyResponse is a verification response to a passed-in JWT.
//     Valid: Whether the token is still valid.
//     Expires: Unix timestamp at when the token will expire.
//     Email: The email of the token’s owner.
type JwtVerifyResponse struct {
	Valid   bool     `json:"valid"`
	Email   string   `json:"email"`
	Expires int64    `json:"expires"`
	Err     string   `json:"err,omitempty"`
}
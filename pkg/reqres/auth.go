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


type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Err      string `json:"err,omitempty"`
}

// TODO: maybe some additional payload about subscription status.
type LogInResponse struct {
	AuthToken string   `json:"token"`
	Err       string   `json:"err,omitempty"`
}

type LogOutRequest struct {
	Err string `json:"err,omitempty"`
}

type LogOutResponse struct {
	Err string `json:"err,omitempty"`
}

type SignUpRequest struct {
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	OptIn    bool   `json:"optIn"`
	Err      string `json:"err,omitempty"`
}

type SignUpResponse struct {
	Err string `json:"err,omitempty"`
}

type EmailVerificationRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type EmailVerificationResponse struct {
	Verified bool   `json:"verified"`
	Err      string `json:"err,omitempty"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
	Err   string `json:"err,omitempty"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Err      string `json:"err,omitempty"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}
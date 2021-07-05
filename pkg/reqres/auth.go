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

package reqres

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Err      string `json:"err,omitempty"`
}

type LogInResponse struct {
	AuthToken string `json:"token"`
	Err       string `json:"err,omitempty"`
}

type LogOutRequest struct {
	Err string `json:"err,omitempty"`
}

type LogOutResponse struct {
	Err string `json:"err,omitempty"`
}

type SignUpRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Err   string `json:"err,omitempty"`
}

type SignUpResponse struct {
	Err string `json:"err,omitempty"`
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

type CreateAccountRequest struct {
	Email                  string `json:"email"`
	Name                   string `json:"name"`
	Token                  string `json:"token"`
	Password               string `json:"password"`
	SubscribeToMailingList bool   `json:"optIn"`
	Err                    string `json:"err,omitempty"`
}

type CreateAccountResponse struct {
	Verified bool   `json:"verified"`
	Err      string `json:"err,omitempty"`
}

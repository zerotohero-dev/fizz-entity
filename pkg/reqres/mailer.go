package reqres

type RelayWelcomeEmailRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayWelcomeEmailResponse struct {
	Relayed bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

type RelayEmailVerifiedEmailRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayEmailVerifiedEmailResponse struct {
	Relayed bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

type RelaySendEmailVerificationEmailRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type RelaySendEmailVerificationEmailResponse struct {
	Relayed bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

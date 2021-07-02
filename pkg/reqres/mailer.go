package reqres

type RelayWelcomeEmailRequest struct {
	Email string `json:"message"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayWelcomeEmailResponse struct {
	Relayed bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

type RelayEmailVerifiedEmailRequest struct {
	Email string `json:"message"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayEmailVerifiedEmailResponse struct {
	Relayed bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

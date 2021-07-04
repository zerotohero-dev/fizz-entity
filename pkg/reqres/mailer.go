package reqres

type RelayWelcomeMessageRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayWelcomeMessageResponse struct {
	Success bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

type RelayEmailVerifiedMessageRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayEmailVerifiedMessageResponse struct {
	Success bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

type RelayEmailVerificationMessageRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type RelayEmailVerificationMessageResponse struct {
	Success bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

type RelayPasswordResetConfirmationMessageRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayPasswordResetConfirmationMessageResponse struct {
	Success bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

type RelayPasswordResetMessageRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Name  string `json:"name"`
	Err   string `json:"err,omitempty"`
}

type RelayPasswordResetMessageResponse struct {
	Success bool   `json:"relayed"`
	Err     string `json:"err,omitempty"`
}

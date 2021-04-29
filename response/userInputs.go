package response

type UserInputID string

const (
	UserInputOTP            = "otp"
	UserInputSecretQuestion = "secret_question"
	UserInputCaptcha        = "captcha"
	UserInputPin            = "pin"
	UserInputConfirmation   = "confirmation"
	UserInputToken          = "token"
)

type UserInput struct {
	ID     UserInputID `json:"id,omitempty"`
	Query  string      `json:"query,omitempty"`
	Answer string      `json:"answer,omitempty"`
	Index  int         `json:"index,omitempty"`
}

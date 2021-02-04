package constants

// ApiStatus is the status of the responses returned from the API
type ApiStatus string

const (
	StatusInitialized       ApiStatus = "initialized"
	StatusFailed            ApiStatus = "failed"
	StatusUserInputRequired ApiStatus = "user_input_required"
	StatusDone              ApiStatus = "done"
)

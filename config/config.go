package config

// Config holds the fields which is specific to the app and which all products need.
type Config struct {
	BundleID  string
	AppKey    string
	AppSecret string
}

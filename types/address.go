package types

type IdentityAddress struct {
	Flat     string `json:"flat,omitempty"`
	Building string `json:"building,omitempty"`
	Full     string `json:"full,omitempty"`
	Area     string `json:"area,omitempty"`
	PoBox    string `json:"poBox,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
}

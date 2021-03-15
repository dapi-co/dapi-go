package types

type Identity struct {
	Nationality    string           `json:"nationality,omitempty"`
	DateOfBirth    string           `json:"dateOfBirth,omitempty"`
	Numbers        []PhoneNumber    `json:"numbers,omitempty"`
	EmailAddress   string           `json:"emailAddress,omitempty"`
	Name           string           `json:"name,omitempty"`
	Address        IdentityAddress  `json:"address,omitempty"`
	Identification []Identification `json:"identification,omitempty"`
}

package response

type Identity struct {
	Nationality    string           `json:"nationality,omitempty"`
	DateOfBirth    string           `json:"dateOfBirth,omitempty"`
	Numbers        []PhoneNumber    `json:"numbers,omitempty"`
	EmailAddress   string           `json:"emailAddress,omitempty"`
	Name           string           `json:"name,omitempty"`
	Address        Address          `json:"address,omitempty"`
	Identification []Identification `json:"identification,omitempty"`
}

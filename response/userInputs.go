package response

type UserInput struct {
	ID     string `json:"id,omitempty"`
	Query  string `json:"query,omitempty"`
	Answer string `json:"answer,omitempty"`
	Index  int    `json:"index,omitempty"`
}

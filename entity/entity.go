package entity

type Request struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

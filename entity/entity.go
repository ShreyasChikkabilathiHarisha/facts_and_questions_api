package entity

type Request struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

type PostItem struct {
	QuestionID string `json:"questionid,omitempty"`
	UserID     string `json:"userid,omitempty"`
	Question   string `json:"question,omitempty"`
	Answer     Answer
	VoteID     Vote
}

type Vote struct {
	vode_id         string
	usersUpvoting   map[string]struct{}
	usersDownvoting map[string]struct{}
}

type Answer struct {
	AnswerID string
	Answer   string
}

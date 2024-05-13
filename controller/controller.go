package controller

import (
	"facts_and_questions_api/entity"
)

type Option func(*Controller)

type Logger interface {
	Printf(format string, v ...interface{})
}

func LogWith(logger Logger) Option {
	return func(c *Controller) {
		c.logger = logger
	}
}

type Controller struct {
	logger       Logger
	questionRepo map[string]interface{}
}

func NewController(options ...Option) *Controller {
	c := &Controller{
		questionRepo: map[string]interface{}{
			"q1": "a1",
			"q2": "a2",
			"q3": "a3",
		},
	}

	for _, o := range options {
		o(c)
	}

	return c
}

// func (c *Controller) log(format string, v ...interface{}) {
// 	if c.logger != nil {
// 		c.logger.Printf(format+"\n", v...)
// 	}
// }

func (c *Controller) Create(request entity.Request) map[string]interface{} {
	c.questionRepo[request.Question] = request.Answer

	return map[string]interface{}{
		request.Question: request.Answer,
	}
}

func (c *Controller) Fetch(question string) interface{} {
	return c.questionRepo[question]
}

func (c *Controller) FetchAllQuestions() map[string]interface{} {
	return c.questionRepo
}

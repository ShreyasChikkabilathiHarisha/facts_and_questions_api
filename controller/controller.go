package controller

import (
	"facts_and_questions_api/entity"
	"log"
)

type Controller struct {
	postItems map[string]entity.PostItem
}

func NewController() *Controller {
	c := &Controller{
		postItems: make(map[string]entity.PostItem),
	}

	return c
}

func (c *Controller) Create(request entity.PostItem) map[string]interface{} {
	c.postItems[request.QuestionID] = request

	return map[string]interface{}{
		request.QuestionID: request,
	}
}

// func (c *Controller) Fetch(question string) interface{} {
// 	return c.questionRepo[question]
// }

func (c *Controller) FetchAllQuestions() map[string]entity.PostItem {
	log.Println(c.postItems)

	return c.postItems
}

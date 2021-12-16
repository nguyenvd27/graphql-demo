package static_database

import "github.com/nguyenvd27/graphql-test/model"

type StaticDatabase struct {
	todos  []*model.Todo
	events []*model.Event
}

func NewStaticDatabase(todos []*model.Todo, events []*model.Event) *StaticDatabase {
	return &StaticDatabase{
		todos:  todos,
		events: events,
	}
}

var Todos = []*model.Todo{
	{
		ID:     "1",
		Text:   "text 1",
		Done:   false,
		UserID: "1",
	},
	{
		ID:     "2",
		Text:   "text 2",
		Done:   false,
		UserID: "2",
	},
}

var Events = []*model.Event{
	{
		ID:      "1",
		Title:   "title 1",
		Content: "content 1",
		UserID:  "1",
	},
	{
		ID:      "2",
		Title:   "title 2",
		Content: "content 2",
		UserID:  "2",
	},
}

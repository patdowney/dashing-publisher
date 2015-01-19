package dashing

type Event struct {
	WidgetID string
	Body     map[string]interface{}
}

type Job interface {
	Work(chan Event)
}

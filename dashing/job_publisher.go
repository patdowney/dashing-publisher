package dashing

import (
	"bytes"
	"encoding/json"
	"net/url"
)

type JobPublisher struct {
	client    *http.Client
	TargetURL *url.URL
	Job       Job
}

func NewJobPublisher(targetURL *url.URL, job Job) *JobPublisher {
	client := &http.Client{}
	jp := &JobPublisher{
		client:    client,
		TargetURL: targetURL,
		Job:       job,
	}

	return jp
}

func (j *JobPublisher) SendEvent(event Event) error {
	data := &bytes.Buffer{}

	enc := json.NewEncoder(data)
	enc.Encode(event.Body)

	resp, err := j.client.Post(j.TargetURL.String(), "application/json", data)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

func (j *JobPublisher) Start() {
	events := make(chan Event)

	go j.Job.Work(events)

	for {
		select {
		case e := <-events:
			j.SendEvent(e)
		}
	}
}

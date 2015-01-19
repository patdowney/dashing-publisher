package dashing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendEvent(client *http.Client, event Event) error {
	data := &bytes.Buffer{}

	enc := json.NewEncoder(data)
	enc.Encode(event.Body)

	widgetUrl := fmt.Sprintf("http://localhost:3000/widgets/%s", event.WidgetID)
	resp, err := client.Post(widgetUrl, "application/json", data)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

func StartPublishLoop(j Job) {
	client := &http.Client{}
	events := make(chan Event)

	go j.Work(events)

	for {
		select {
		case e := <-events:
			SendEvent(client, e)
		}
	}
}

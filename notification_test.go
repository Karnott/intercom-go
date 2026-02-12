package intercom

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var errTest = errors.New("Test Error")

type alwaysErrorReader struct{}

func (r *alwaysErrorReader) Read(p []byte) (n int, err error) {
	return 0, errTest
}

func TestParsingFromReader(t *testing.T) {
	r, _ := os.Open("fixtures/notification.json")
	n, _ := NewNotification(r)
	if n.ID != "notif_ccd8a4d0-f965-11e3-a367-c779cae3e1b3" {
		t.Errorf("Notification did not have ID")
	}
	if n.CreatedAt != 1392731331 {
		t.Errorf("Notification did not have CreatedAt")
	}
	if n.DeliveryAttempts != 1 {
		t.Errorf("Notification did not have DeliveryAttempts")
	}
	if n.Topic != "company.created" {
		t.Errorf("Notification did not have Topic")
	}
	if n.FirstSentAt != 1392731392 {
		t.Errorf("Notification did not have FirstSentAt")
	}
	if n.Company == nil {
		t.Errorf("Notification did not have Company")
	}
	if n.Conversation != nil {
		t.Errorf("Notification should not have Conversation")
	}
	if n.Event != nil {
		t.Errorf("Notification should not have Event")
	}
	if n.Tag != nil {
		t.Errorf("Notification should not have Tag")
	}
	if n.Contact != nil {
		t.Errorf("Notification should not have Contact")
	}
}

func TestParsingError(t *testing.T) {
	r := &alwaysErrorReader{}
	_, err := NewNotification(r)
	if err == nil {
		t.Errorf("Error not returned")
	}
}

func TestParsingConverationFromReader(t *testing.T) {
	topics := []string{
		"conversation.user.created",
		"conversation.user.replied",
		"conversation.admin.replied",
		"conversation.admin.single.created",
		"conversation.admin.assigned",
		"conversation.admin.noted",
		"conversation.admin.closed",
		"conversation.admin.opened",
	}

	for _, topic := range topics {
		payload, _ := ioutil.ReadFile("fixtures/conversation.json")
		r := strings.NewReader(fmt.Sprintf(`{
			"topic": "%s",
			"data": {
				"item": %s
			}
		}`, topic, string(payload)))
		n, _ := NewNotification(r)
		if n.Conversation == nil {
			t.Errorf("Notification did not have Conversation")
		}
	}
}

func TestParsingContactFromReader(t *testing.T) {
	topics := []string{
		"contact.created",
		"contact.deleted",
		"contact.unsubscribed",
		"contact.email.updated",
		"user.created",
		"user.deleted",
		"user.unsubscribed",
		"user.email.updated",
	}

	for _, topic := range topics {
		payload, _ := ioutil.ReadFile("fixtures/contact.json")
		r := strings.NewReader(fmt.Sprintf(`{
			"topic": "%s",
			"data": {
				"item": %s
			}
		}`, topic, string(payload)))
		n, _ := NewNotification(r)
		if n.Contact == nil {
			t.Errorf("Notification did not have Contact for topic %s", topic)
		}
	}
}

func TestParsingTagFromReader(t *testing.T) {
	topics := []string{
		"contact.tag.created",
		"contact.tag.deleted",
		"user.tag.created",
		"user.tag.deleted",
	}

	for _, topic := range topics {
		payload, _ := ioutil.ReadFile("fixtures/tag.json")
		r := strings.NewReader(fmt.Sprintf(`{
			"topic": "%s",
			"data": {
				"item": %s
			}
		}`, topic, string(payload)))
		n, _ := NewNotification(r)
		if n.Tag == nil {
			t.Errorf("Notification did not have Tag")
		}
	}
}

func TestParsingEventFromReader(t *testing.T) {
	topics := []string{
		"event.created",
	}

	for _, topic := range topics {
		payload, _ := ioutil.ReadFile("fixtures/event.json")
		r := strings.NewReader(fmt.Sprintf(`{
			"topic": "%s",
			"data": {
				"item": %s
			}
		}`, topic, string(payload)))
		n, _ := NewNotification(r)
		if n.Event == nil {
			t.Errorf("Notification did not have Event")
		}
	}
}

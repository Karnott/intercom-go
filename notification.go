package intercom

import (
	"encoding/json"
	"io"
)

// Notification is the object delivered to a webhook.
type Notification struct {
	ID               string        `json:"id,omitempty"`
	CreatedAt        int64         `json:"created_at,omitempty"`
	Topic            string        `json:"topic,omitempty"`
	DeliveryAttempts int64         `json:"delivery_attempts,omitempty"`
	FirstSentAt      int64         `json:"first_sent_at,omitempty"`
	RawData          *Data         `json:"data,omitempty"`
	Conversation     *Conversation `json:"-"`
	Contact          *Contact      `json:"-"`
	Tag              *Tag          `json:"-"`
	Company          *Company      `json:"-"`
	Event            *Event        `json:"-"`
}

// Data is the data node of the notification.
type Data struct {
	Item json.RawMessage `json:"item,omitempty"`
}

// NewNotification parses a Notification from json read from an io.Reader.
// It may only contain partial objects (such as a single conversation part)
// depending on what is provided by the webhook.
func NewNotification(r io.Reader) (*Notification, error) {
	notification := &Notification{
		RawData: &Data{},
	}
	err := json.NewDecoder(r).Decode(notification)
	if err != nil {
		return nil, err
	}
	switch notification.Topic {
	case "conversation.user.created",
		"conversation.user.replied",
		"conversation.admin.replied",
		"conversation.admin.single.created",
		"conversation.admin.assigned",
		"conversation.admin.noted",
		"conversation.admin.closed",
		"conversation.admin.opened":
		c := &Conversation{}
		json.Unmarshal(notification.RawData.Item, c)
		notification.Conversation = c
	case "contact.created",
		"contact.deleted",
		"contact.unsubscribed",
		"contact.email.updated",
		"contact.user.created",
		"contact.lead.created",
		"contact.lead.updated",
		"user.created",
		"user.deleted",
		"user.unsubscribed",
		"user.email.updated":
		ct := &Contact{}
		json.Unmarshal(notification.RawData.Item, ct)
		notification.Contact = ct
	case "contact.tag.created",
		"contact.tag.deleted",
		"contact.user.tag.created",
		"contact.user.tag.deleted",
		"contact.lead.tag.created",
		"contact.lead.tag.deleted",
		"user.tag.created",
		"user.tag.deleted":
		t := &Tag{}
		json.Unmarshal(notification.RawData.Item, t)
		notification.Tag = t
	case "company.created",
		"company.updated",
		"company.deleted":
		c := &Company{}
		json.Unmarshal(notification.RawData.Item, c)
		notification.Company = c
	case "event.created":
		e := &Event{}
		json.Unmarshal(notification.RawData.Item, e)
		notification.Event = e
	}
	return notification, nil
}

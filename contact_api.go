package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/karnott/intercom-go/interfaces"
)

// ContactRepository defines the interface for working with Contacts through the API.
type ContactRepository interface {
	find(ContactIdentifiers) (Contact, error)
	list(contactListParams) (ContactList, error)
	create(*Contact) (Contact, error)
	update(*Contact) (Contact, error)
	merge(sourceID, targetID string) (Contact, error)
	archive(id string) (Contact, error)
	unarchive(id string) (Contact, error)
	delete(id string) (Contact, error)
}

// ContactAPI implements ContactRepository
type ContactAPI struct {
	httpClient interfaces.HTTPClient
}

func (api ContactAPI) find(params ContactIdentifiers) (Contact, error) {
	switch {
	case params.ID != "":
		return unmarshalToContact(api.httpClient.Get(fmt.Sprintf("/contacts/%s", params.ID), nil))
	case params.ExternalID != "":
		return api.findByExternalID(params.ExternalID)
	}
	return Contact{}, errors.New("Missing Contact Identifier")
}

func (api ContactAPI) findByExternalID(externalID string) (Contact, error) {
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"field":    "external_id",
			"operator": "=",
			"value":    externalID,
		},
	}
	data, err := api.httpClient.Post("/contacts/search", &searchQuery)
	if err != nil {
		return Contact{}, err
	}
	contactList := ContactList{}
	err = json.Unmarshal(data, &contactList)
	if err != nil {
		return Contact{}, err
	}
	if len(contactList.Contacts) == 0 {
		return Contact{}, errors.New("Contact not found")
	}
	return contactList.Contacts[0], nil
}

func (api ContactAPI) list(params contactListParams) (ContactList, error) {
	contactList := ContactList{}
	data, err := api.httpClient.Get("/contacts", params)
	if err != nil {
		return contactList, err
	}
	err = json.Unmarshal(data, &contactList)
	return contactList, err
}

func (api ContactAPI) create(contact *Contact) (Contact, error) {
	return unmarshalToContact(api.httpClient.Post("/contacts", contact))
}

func (api ContactAPI) update(contact *Contact) (Contact, error) {
	if contact.ID == "" {
		return Contact{}, errors.New("Missing Contact ID for update")
	}
	return unmarshalToContact(api.httpClient.Patch(fmt.Sprintf("/contacts/%s", contact.ID), contact))
}

type mergeRequest struct {
	From string `json:"from"`
	Into string `json:"into"`
}

func (api ContactAPI) merge(sourceID, targetID string) (Contact, error) {
	return unmarshalToContact(api.httpClient.Post("/contacts/merge", &mergeRequest{From: sourceID, Into: targetID}))
}

func (api ContactAPI) archive(id string) (Contact, error) {
	return unmarshalToContact(api.httpClient.Post(fmt.Sprintf("/contacts/%s/archive", id), nil))
}

func (api ContactAPI) unarchive(id string) (Contact, error) {
	return unmarshalToContact(api.httpClient.Post(fmt.Sprintf("/contacts/%s/unarchive", id), nil))
}

func (api ContactAPI) delete(id string) (Contact, error) {
	contact := Contact{}
	data, err := api.httpClient.Delete(fmt.Sprintf("/contacts/%s", id), nil)
	if err != nil {
		return contact, err
	}
	err = json.Unmarshal(data, &contact)
	return contact, err
}

func unmarshalToContact(data []byte, err error) (Contact, error) {
	savedContact := Contact{}
	if err != nil {
		return savedContact, err
	}
	err = json.Unmarshal(data, &savedContact)
	return savedContact, err
}

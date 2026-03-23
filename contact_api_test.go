package intercom

import (
	"io/ioutil"
	"testing"
)

func TestContactAPIFind(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/54c42e7ea7a765fa7", t: t}
	api := ContactAPI{httpClient: &http}
	contact, err := api.find(ContactIdentifiers{ID: "54c42e7ea7a765fa7"})
	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}
	if contact.ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contact.ID)
	}
	if contact.Phone != "+1234567890" {
		t.Errorf("Phone was %s, expected +1234567890", contact.Phone)
	}
	if contact.ExternalID != "123" {
		t.Errorf("ExternalID was %s, expected 123", contact.ExternalID)
	}
	if contact.Role != "user" {
		t.Errorf("Role was %s, expected user", contact.Role)
	}
}

func TestContactAPIFindByExternalID(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/find_by_external_id/123", t: t}
	api := ContactAPI{httpClient: &http}
	contact, err := api.find(ContactIdentifiers{ExternalID: "123"})
	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}
	if contact.ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contact.ID)
	}
	if contact.ExternalID != "123" {
		t.Errorf("ExternalID was %s, expected 123", contact.ExternalID)
	}
}

func TestContactAPIListDefault(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contacts.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contactList, _ := api.list(contactListParams{})
	contacts := contactList.Contacts
	if contacts[0].ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contacts[0].ID)
	}
	pages := contactList.Pages
	if pages.Page != 1 {
		t.Errorf("Page was %d, expected 1", pages.Page)
	}
	if pages.Next == nil {
		t.Errorf("Expected cursor-based next page")
	}
	if pages.Next.StartingAfter != "abc123cursor" {
		t.Errorf("StartingAfter was %s, expected abc123cursor", pages.Next.StartingAfter)
	}
	if contactList.TotalCount != 180 {
		t.Errorf("TotalCount was %d, expected 180", contactList.TotalCount)
	}
}

func TestContactAPIListByEmail(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contacts.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contactList, _ := api.list(contactListParams{Email: "mycontact@example.io"})
	contacts := contactList.Contacts
	if contacts[0].ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contacts[0].ID)
	}
	if clParams, ok := http.lastQueryParams.(contactListParams); !ok || clParams.Email != "mycontact@example.io" {
		t.Errorf("Email expected to be mycontact@example.io, but was %s", clParams.Email)
	}
}

func TestContactAPICreate(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{Email: "mycontact@example.io"}
	_, err := api.create(contact)
	if err != nil {
		t.Errorf("Error creating contact %s", err)
	}
}

func TestContactAPIUpdate(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/54c42e7ea7a765fa7", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{ID: "54c42e7ea7a765fa7", Email: "mycontact@example.io"}
	_, err := api.update(contact)
	if err != nil {
		t.Errorf("Error updating contact %s", err)
	}
}

func TestContactAPIMerge(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/merge", t: t}
	api := ContactAPI{httpClient: &http}
	_, err := api.merge("source123", "target456")
	if err != nil {
		t.Errorf("Error merging contacts %s", err)
	}
}

func TestContactAPIArchive(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/54c42e7ea7a765fa7/archive", t: t}
	api := ContactAPI{httpClient: &http}
	_, err := api.archive("54c42e7ea7a765fa7")
	if err != nil {
		t.Errorf("Error archiving contact %s", err)
	}
}

func TestContactAPIUnarchive(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/54c42e7ea7a765fa7/unarchive", t: t}
	api := ContactAPI{httpClient: &http}
	_, err := api.unarchive("54c42e7ea7a765fa7")
	if err != nil {
		t.Errorf("Error unarchiving contact %s", err)
	}
}

func TestContactAPIDelete(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/b123d", t: t}
	api := ContactAPI{httpClient: &http}
	returned, _ := api.delete("b123d")
	if returned.ExternalID != "123" {
		t.Errorf("Expected ExternalID %s, got %s", "123", returned.ExternalID)
	}
}

type TestContactHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
	lastQueryParams interface{}
}

func (t *TestContactHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	t.lastQueryParams = queryParams
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t *TestContactHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called: %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t *TestContactHTTPClient) Put(uri string, body interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called: %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t *TestContactHTTPClient) Patch(uri string, body interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called: %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t *TestContactHTTPClient) Delete(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

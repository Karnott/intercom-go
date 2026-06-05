package intercom

import (
	"testing"
)

func TestContactFindByID(t *testing.T) {
	contact, _ := (&ContactService{Repository: TestContactAPI{t: t}}).FindByID("46adad3f09126dca")
	if contact.ID != "46adad3f09126dca" {
		t.Errorf("Contact not found")
	}
}

func TestContactFindByExternalID(t *testing.T) {
	contact, _ := (&ContactService{Repository: TestContactAPI{t: t}}).FindByExternalID("134d")
	if contact.ExternalID != "134d" {
		t.Errorf("Contact not found")
	}
}

func TestContactList(t *testing.T) {
	contactList, _ := (&ContactService{Repository: TestContactAPI{t: t}}).ListByEmail("jamie@example.io", PageParams{})
	contacts := contactList.Contacts
	if contacts[0].ID != "46adad3f09126dca" {
		t.Errorf("Contact not listed")
	}
}

func TestContactListEmail(t *testing.T) {
	contactList, _ := (&ContactService{Repository: TestContactAPI{t: t}}).List(PageParams{})
	contacts := contactList.Contacts
	if contacts[0].ID != "46adad3f09126dca" {
		t.Errorf("Contact not listed")
	}
}

func TestContactCreate(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	contact := Contact{Email: "some@email.com"}
	c, _ := contactService.Create(&contact)
	if c.Email != contact.Email {
		t.Errorf("expected returned contact to have email %s, got %s", contact.Email, c.Email)
	}
}

func TestContactUpdate(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	contact := Contact{ID: "abc123", Email: "some@email.com"}
	c, _ := contactService.Update(&contact)
	if c.Email != contact.Email {
		t.Errorf("expected returned contact to have email %s, got %s", contact.Email, c.Email)
	}
}

func TestContactMerge(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	c, _ := contactService.Merge("source123", "target456")
	if c.ID != "target456" {
		t.Errorf("expected merged contact to have ID target456, got %s", c.ID)
	}
}

func TestContactArchive(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	c, _ := contactService.Archive("abc123")
	if c.ID != "abc123" {
		t.Errorf("expected archived contact to have ID abc123, got %s", c.ID)
	}
}

func TestContactUnarchive(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	c, _ := contactService.Unarchive("abc123")
	if c.ID != "abc123" {
		t.Errorf("expected unarchived contact to have ID abc123, got %s", c.ID)
	}
}

func TestContactDelete(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	c, _ := contactService.Delete("abc123")
	if c.ID != "abc123" {
		t.Errorf("expected deleted contact to have ID abc123, got %s", c.ID)
	}
}

func TestContactAttachCompany(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	company, _ := contactService.AttachCompany("contact123", "company456")
	if company.ID != "company456" {
		t.Errorf("expected attached company to have ID company456, got %s", company.ID)
	}
}

func TestContactDetachCompany(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	company, _ := contactService.DetachCompany("contact123", "company456")
	if company.ID != "company456" {
		t.Errorf("expected detached company to have ID company456, got %s", company.ID)
	}
}

func TestContactMessageAddress(t *testing.T) {
	contact := Contact{ID: "abc123", Email: "some@email.com"}
	address := contact.MessageAddress()
	if address.ID != "abc123" {
		t.Errorf("Contact address had wrong ID")
	}
	if address.Type != "contact" {
		t.Errorf("Contact address was not of type contact, was %s", address.Type)
	}
	if address.Email != "some@email.com" {
		t.Errorf("Contact address had wrong Email")
	}
}

type TestContactAPI struct {
	t *testing.T
}

func (t TestContactAPI) find(params ContactIdentifiers) (Contact, error) {
	return Contact{ID: params.ID, Email: params.ExternalID, ExternalID: params.ExternalID}, nil
}

func (t TestContactAPI) list(params contactListParams) (ContactList, error) {
	return ContactList{Contacts: []Contact{{ID: "46adad3f09126dca", Email: "jamie@example.io", ExternalID: "aa123"}}}, nil
}

func (t TestContactAPI) create(c *Contact) (Contact, error) {
	return Contact{ID: "new123", Email: c.Email}, nil
}

func (t TestContactAPI) update(c *Contact) (Contact, error) {
	return Contact{ID: c.ID, Email: c.Email, ExternalID: c.ExternalID}, nil
}

func (t TestContactAPI) merge(sourceID, targetID string) (Contact, error) {
	return Contact{ID: targetID}, nil
}

func (t TestContactAPI) archive(id string) (Contact, error) {
	return Contact{ID: id}, nil
}

func (t TestContactAPI) unarchive(id string) (Contact, error) {
	return Contact{ID: id}, nil
}

func (t TestContactAPI) delete(id string) (Contact, error) {
	return Contact{ID: id}, nil
}

func (t TestContactAPI) attachCompany(contactID, companyID string) (Company, error) {
	return Company{ID: companyID, CompanyID: "337631"}, nil
}

func (t TestContactAPI) detachCompany(contactID, companyID string) (Company, error) {
	return Company{ID: companyID}, nil
}

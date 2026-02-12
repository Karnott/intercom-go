package intercom

import "fmt"

// ContactService handles interactions with the API through a ContactRepository.
type ContactService struct {
	Repository ContactRepository
}

// ContactList holds a list of Contacts and cursor-based pagination info.
type ContactList struct {
	Pages      CursorPages `json:"pages"`
	TotalCount int64       `json:"total_count,omitempty"`
	Contacts   []Contact   `json:"data,omitempty"`
}

// Contact represents a Contact (user or lead) within Intercom.
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Contact struct {
	ID                     string                 `json:"id,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	ExternalID             string                 `json:"external_id,omitempty"`
	Role                   string                 `json:"role,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	EmailDomain            string                 `json:"email_domain,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`
	FormattedPhone         string                 `json:"formatted_phone,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	WorkspaceID            string                 `json:"workspace_id,omitempty"`
	Avatar                 *ContactAvatar         `json:"avatar,omitempty"`
	OwnerID                int64                  `json:"owner_id,omitempty"`
	SocialProfiles         *SocialProfileList     `json:"social_profiles,omitempty"`
	HasHardBounced         bool                   `json:"has_hard_bounced,omitempty"`
	MarkedEmailAsSpam      bool                   `json:"marked_email_as_spam,omitempty"`
	UnsubscribedFromEmails *bool                  `json:"unsubscribed_from_emails,omitempty"`
	CreatedAt              int64                  `json:"created_at,omitempty"`
	UpdatedAt              int64                  `json:"updated_at,omitempty"`
	SignedUpAt             int64                  `json:"signed_up_at,omitempty"`
	LastSeenAt             int64                  `json:"last_seen_at,omitempty"`
	LastRepliedAt          int64                  `json:"last_replied_at,omitempty"`
	LastContactedAt        int64                  `json:"last_contacted_at,omitempty"`
	LastEmailOpenedAt      int64                  `json:"last_email_opened_at,omitempty"`
	LastEmailClickedAt     int64                  `json:"last_email_clicked_at,omitempty"`
	LanguageOverride       string                 `json:"language_override,omitempty"`
	Browser                string                 `json:"browser,omitempty"`
	BrowserVersion         string                 `json:"browser_version,omitempty"`
	BrowserLanguage        string                 `json:"browser_language,omitempty"`
	OS                     string                 `json:"os,omitempty"`
	Tags                   *TagList               `json:"tags,omitempty"`
	Segments               *SegmentList           `json:"segments,omitempty"`
	Companies              *CompanyList           `json:"companies,omitempty"`
	Location               *ContactLocation       `json:"location,omitempty"`
	Referrer               string                 `json:"referrer,omitempty"`
	UTMCampaign            string                 `json:"utm_campaign,omitempty"`
	UTMContent             string                 `json:"utm_content,omitempty"`
	UTMMedium              string                 `json:"utm_medium,omitempty"`
	UTMSource              string                 `json:"utm_source,omitempty"`
	UTMTerm                string                 `json:"utm_term,omitempty"`
	EnabledPushMessaging   *bool                  `json:"enabled_push_messaging,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
}

// ContactAvatar represents a contact's avatar.
type ContactAvatar struct {
	Type     string `json:"type,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

// ContactLocation holds location data for a contact.
type ContactLocation struct {
	Type          string `json:"type,omitempty"`
	Country       string `json:"country,omitempty"`
	Region        string `json:"region,omitempty"`
	City          string `json:"city,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	ContinentCode string `json:"continent_code,omitempty"`
}

// SocialProfileList is a list of SocialProfiles for a Contact.
type SocialProfileList struct {
	SocialProfiles []SocialProfile `json:"social_profiles,omitempty"`
}

// SocialProfile represents a social account for a Contact.
type SocialProfile struct {
	Name     string `json:"name,omitempty"`
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
}

// ContactIdentifiers hold identifiers for finding a Contact.
type ContactIdentifiers struct {
	ID         string `url:"-"`
	ExternalID string `url:"external_id,omitempty"`
}

type contactListParams struct {
	PageParams
	StartingAfter string `url:"starting_after,omitempty"`
	Email         string `url:"email,omitempty"`
}

// FindByID looks up a Contact by their Intercom ID.
func (c *ContactService) FindByID(id string) (Contact, error) {
	return c.Repository.find(ContactIdentifiers{ID: id})
}

// FindByExternalID looks up a Contact by their external ID.
func (c *ContactService) FindByExternalID(externalID string) (Contact, error) {
	return c.Repository.find(ContactIdentifiers{ExternalID: externalID})
}

// List all Contacts for App.
func (c *ContactService) List(params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params})
}

// ListByEmail looks up a list of Contacts by their Email.
func (c *ContactService) ListByEmail(email string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, Email: email})
}

// Create Contact.
func (c *ContactService) Create(contact *Contact) (Contact, error) {
	return c.Repository.create(contact)
}

// Update Contact.
func (c *ContactService) Update(contact *Contact) (Contact, error) {
	return c.Repository.update(contact)
}

// Merge two Contacts. The source contact is merged into the target.
func (c *ContactService) Merge(sourceID, targetID string) (Contact, error) {
	return c.Repository.merge(sourceID, targetID)
}

// Archive a Contact.
func (c *ContactService) Archive(id string) (Contact, error) {
	return c.Repository.archive(id)
}

// Unarchive a Contact.
func (c *ContactService) Unarchive(id string) (Contact, error) {
	return c.Repository.unarchive(id)
}

// Delete Contact permanently.
func (c *ContactService) Delete(id string) (Contact, error) {
	return c.Repository.delete(id)
}

// MessageAddress gets the address for a Contact in order to message them.
func (c Contact) MessageAddress() MessageAddress {
	return MessageAddress{
		Type:  "contact",
		ID:    c.ID,
		Email: c.Email,
	}
}

func (c Contact) String() string {
	return fmt.Sprintf("[intercom] contact { id: %s, role: %s, name: %s, email: %s }", c.ID, c.Role, c.Name, c.Email)
}

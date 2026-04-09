# Intercom-Go

Thin client for the [Intercom](https://www.intercom.io) API (v2.15).

## Install

```
go get github.com/karnott/intercom-go/v2
```

## Usage

### Getting a Client

```go
import (
	intercom "github.com/karnott/intercom-go/v2"
)

ic := intercom.NewClient("your_access_token")
```

This client can then be used to make requests.

You can find your access token [here](https://app.intercom.com/developers/_). Learn more about access tokens [here](https://developers.intercom.io/docs/personal-access-tokens).

If you are building a third party application you can get your OAuth token by [setting up OAuth](https://developers.intercom.io/page/setting-up-oauth) for Intercom.

#### Client Options

The client can be configured with different options by calls to `ic.Option`:

```go
ic.Option(intercom.TraceHTTP(true)) // turn http tracing on
ic.Option(intercom.BaseURI("http://intercom.dev")) // change the base uri used, useful for testing
ic.Option(intercom.SetHTTPClient(myHTTPClient)) // set a new HTTP client, see below for more info
```

or combined:

```go
ic.Option(intercom.TraceHTTP(true), intercom.BaseURI("http://intercom.dev"))
```

### Contacts

Since API v2.0, Users and Leads have been unified into a single **Contact** model. A Contact has a `role` field that is either `"user"` or `"lead"`.

#### Create

```go
contact := intercom.Contact{
	Role:  "user",
	ExternalID: "27",
	Email: "test@example.com",
	Name:  "InterGopher",
	SignedUpAt: int64(time.Now().Unix()),
	CustomAttributes: map[string]interface{}{"is_cool": true},
}
savedContact, err := ic.Contacts.Create(&contact)
```

#### Update

```go
contact := intercom.Contact{
	ID:   "6329e838deab13e266c3602d",
	Name: "Updated Name",
	CustomAttributes: map[string]interface{}{"is_cool": true},
}
savedContact, err := ic.Contacts.Update(&contact)
```

* `ID` is required for updates.

#### Find

```go
contact, err := ic.Contacts.FindByID("6329e838deab13e266c3602d")
```

```go
contact, err := ic.Contacts.FindByExternalID("27")
```

#### List

```go
contactList, err := ic.Contacts.List(intercom.PageParams{Page: 1, PerPage: 50})
contactList.Pages    // CursorPages with cursor-based pagination
contactList.Contacts // []Contact
```

```go
contactList, err := ic.Contacts.ListByEmail("test@example.com", intercom.PageParams{})
```

#### Merge

Merge a source contact into a target contact:

```go
mergedContact, err := ic.Contacts.Merge("source_contact_id", "target_contact_id")
```

#### Archive / Unarchive / Delete

```go
contact, err := ic.Contacts.Archive("6329e838deab13e266c3602d")
contact, err := ic.Contacts.Unarchive("6329e838deab13e266c3602d")
contact, err := ic.Contacts.Delete("6329e838deab13e266c3602d")
```

##### Adding/Removing Companies

Adding a Company:

```go
companyList := intercom.CompanyList{
	Companies: []intercom.Company{
		{CompanyID: "5"},
	},
}
contact := intercom.Contact{
	ID: "6329e838deab13e266c3602d",
	Companies: &companyList,
}
```

Removing is similar, but adding a `Remove: intercom.Bool(true)` attribute to a company.

### Companies

#### Save

```go
company := intercom.Company{
	CompanyID: "27",
	Name: "My Co",
	CustomAttributes: map[string]interface{}{"is_cool": true},
	Plan: &intercom.Plan{Name: "MyPlan"},
}
savedCompany, err := ic.Companies.Save(&company)
```

* `CompanyID` is required.

#### Find

```go
company, err := ic.Companies.FindByID("46adad3f09126dca")
```

```go
company, err := ic.Companies.FindByCompanyID("27")
```

```go
company, err := ic.Companies.FindByName("My Co")
```

#### List

```go
companyList, err := ic.Companies.List(intercom.PageParams{Page: 2})
companyList.Pages     // page information
companyList.Companies // []Company
```

```go
companyList, err := ic.Companies.ListBySegment("segmentID123", intercom.PageParams{})
```

```go
companyList, err := ic.Companies.ListByTag("42", intercom.PageParams{})
```

#### List Contacts

```go
contactList, err := ic.Companies.ListContactsByID("46adad3f09126dca", intercom.PageParams{})
contactList.Contacts // []Contact
```

```go
contactList, err := ic.Companies.ListContactsByCompanyID("27", intercom.PageParams{})
```

### Events

#### Save

```go
event := intercom.Event{
	UserID: "27",
	EventName: "bought_item",
	CreatedAt: int64(time.Now().Unix()),
	Metadata: map[string]interface{}{"item_name": "PocketWatch"},
}
err := ic.Events.Save(&event)
```

* One of `UserID`, `ID`, or `Email` is required.
* `EventName` is required.
* `CreatedAt` is optional, must be an integer representing seconds since Unix Epoch. Will be set to _now_ unless given.
* `Metadata` is optional, and can be constructed as a `map[string]interface{}`.

### Admins

#### List

```go
adminList, err := ic.Admins.List()
admins := adminList.Admins
```

### Tags

#### List

```go
tagList, err := ic.Tags.List()
tags := tagList.Tags
```

#### Save

```go
tag := intercom.Tag{Name: "GoTag"}
savedTag, err := ic.Tags.Save(&tag)
```

`Name` is required. Passing an `ID` will attempt to update the tag with that ID.

#### Delete

```go
err := ic.Tags.Delete("6")
```

#### Tagging Contacts/Companies

```go
taggingList := intercom.TaggingList{Name: "GoTag", Users: []intercom.Tagging{{UserID: "27"}}}
savedTag, err := ic.Tags.Tag(&taggingList)
```

A `Tagging` can identify a Contact or Company, and can be set to `Untag`:

```go
taggingList := intercom.TaggingList{Name: "GoTag", Users: []intercom.Tagging{{UserID: "27", Untag: intercom.Bool(true)}}}
savedTag, err := ic.Tags.Tag(&taggingList)
```

### Segments

#### List

```go
segmentList, err := ic.Segments.List()
segments := segmentList.Segments
```

#### Find

```go
segment, err := ic.Segments.Find("abc312daf2397")
```

### Messages

#### New Admin to Contact Email

```go
msg := intercom.NewEmailMessage(intercom.PERSONAL_TEMPLATE, &intercom.Admin{ID: "1234"}, &intercom.Contact{Email: "test@example.com"}, "subject", "body")
savedMessage, err := ic.Messages.Save(&msg)
```

Can use `intercom.PLAIN_TEMPLATE` too.

#### New Admin to Contact InApp

```go
msg := intercom.NewInAppMessage(&intercom.Admin{ID: "1234"}, &intercom.Contact{Email: "test@example.com"}, "body")
savedMessage, err := ic.Messages.Save(&msg)
```

#### New Contact Message

```go
msg := intercom.NewContactMessage(&intercom.Contact{Email: "test@example.com"}, "body")
savedMessage, err := ic.Messages.Save(&msg)
```

### Conversations

#### Find Conversation

```go
convo, err := ic.Conversations.Find("1234")
```

#### List Conversations

All:

```go
convoList, err := ic.Conversations.ListAll(intercom.PageParams{})
```

By Contact:

```go
convoList, err := ic.Conversations.ListByContact(&contact, intercom.SHOW_ALL, intercom.PageParams{})
convoList, err := ic.Conversations.ListByContact(&contact, intercom.SHOW_UNREAD, intercom.PageParams{})
```

By Admin:

```go
convoList, err := ic.Conversations.ListByAdmin(&admin, intercom.SHOW_ALL, intercom.PageParams{})
convoList, err := ic.Conversations.ListByAdmin(&admin, intercom.SHOW_OPEN, intercom.PageParams{})
convoList, err := ic.Conversations.ListByAdmin(&admin, intercom.SHOW_CLOSED, intercom.PageParams{})
```

#### Reply

Contact reply:

```go
convo, err := ic.Conversations.Reply("1234", &contact, intercom.CONVERSATION_COMMENT, "my message")
```

Contact reply with attachment:

```go
convo, err := ic.Conversations.ReplyWithAttachmentURLs("1234", &contact, intercom.CONVERSATION_COMMENT, "my message", []string{"http://www.example.com/attachment.jpg"})
```

Admin reply:

```go
convo, err := ic.Conversations.Reply("1234", &admin, intercom.CONVERSATION_COMMENT, "my message")
```

Admin note:

```go
convo, err := ic.Conversations.Reply("1234", &admin, intercom.CONVERSATION_NOTE, "my message to just admins")
```

#### Open and Close

```go
convo, err := ic.Conversations.Open("1234", &openerAdmin)
convo, err := ic.Conversations.Close("1234", &closerAdmin)
```

#### Assign

```go
convo, err := ic.Conversations.Assign("1234", &assignerAdmin, &assigneeAdmin)
```

### Articles

#### List

```go
articleList, err := ic.Articles.List(intercom.PageParams{Page: 1, PerPage: 50})
articleList.Articles // []Article
```

#### Find

```go
article, err := ic.Articles.Find("123")
```

### Collections

Collections support hierarchical nesting via `ParentID`. Top-level collections have a nil `ParentID`; sub-collections (formerly "sections" in API < v2.10) have a `ParentID` pointing to their parent collection.

#### List

```go
collectionList, err := ic.Collections.List(intercom.PageParams{Page: 1, PerPage: 50})
collectionList.Collections // []Collection
```

#### Find

```go
collection, err := ic.Collections.Find("123")
```

### Webhooks / Notifications

If you have received a JSON webhook notification, you can convert it into Intercom objects. A Notification can be created from any `io.Reader`, typically an http request body:

```go
notif, err := intercom.NewNotification(r)
```

The returned Notification will contain exactly 1 of the `Company`, `Conversation`, `Event`, `Tag` or `Contact` fields populated. It may only contain partial objects depending on what is provided by the webhook.

### Errors

Errors returned from the API implement `intercom.IntercomError` and can be checked:

```go
_, err := ic.Contacts.FindByID("doesnotexist")
if herr, ok := err.(intercom.IntercomError); ok && herr.GetCode() == "not_found" {
	fmt.Print(herr)
}
```

### HTTP Client

The HTTP Client used by this package can be swapped out for one of your choosing, with your own configuration, it just needs to implement the HTTPClient interface:

```go
type HTTPClient interface {
	Get(string, interface{}) ([]byte, error)
	Post(string, interface{}) ([]byte, error)
	Patch(string, interface{}) ([]byte, error)
	Delete(string, interface{}) ([]byte, error)
}
```

It'll need to work with `accessToken` and `baseURI` values. See the provided client for an example. Then create an Intercom Client and inject the HTTPClient:

```go
ic := intercom.Client{}
ic.Option(intercom.SetHTTPClient(myHTTPClient))
// ready to go!
```

### On Bools

Due to the way Go represents the zero value for a bool, it's necessary to pass pointers to bool instead in some places.

The helper `intercom.Bool(true)` creates these for you.

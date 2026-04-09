package main

import (
	"log"
	"os"

	intercom "github.com/karnott/intercom-go/v2"
)

func main() {
	token := os.Getenv("INTERCOM_ACCESS_TOKEN")
	if token == "" {
		log.Fatal("INTERCOM_ACCESS_TOKEN environment variable is required")
	}

	ic := intercom.NewClient(token)
	ic.Option(intercom.TraceHTTP(true))

	u := new(intercom.Contact)
	// user: "[intercom] contact { id: , role: user, name: PAUL BOUSQUET, email: paulbousquet46@gmail.com }"
	// user_id: "34585"
	u.Role = "user"
	u.Name = "PAUL BOUSQUET"
	u.Email = "paulbousquet46@gmail.com"
	u.ExternalID = "34585"

	saved, createErr := ic.Contacts.Create(u)
	if createErr != nil {
		// Contact may already exist, try to find and update
		if u.ExternalID != "" {
			existing, findErr := ic.Contacts.FindByExternalID(u.ExternalID)
			log.Printf("FindByExternalID result: %v, err: %v", existing, findErr)
			if findErr == nil {
				u.ID = existing.ID
				log.Printf("Updating with ID: %s", u.ID)
				update, err := ic.Contacts.Update(u)
				if err != nil {
					log.Printf("Error update contact: %v", err)
					return
				}
				log.Printf("Updated contact: %v", update)
			}
		}
		//return saved, createErr
		log.Printf("Error creating contact: %v", createErr)
	}
	log.Printf("Created contact: %v", saved)
	//return saved, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	intercom "github.com/karnott/intercom-go/v2"
)

type ExportArticle struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	State       string `json:"state"`
	URL         string `json:"url"`
}

type ExportCollection struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Description    string             `json:"description,omitempty"`
	URL            string             `json:"url,omitempty"`
	SubCollections []ExportCollection `json:"sub_collections,omitempty"`
	Articles       []ExportArticle    `json:"articles,omitempty"`
}

func main() {
	token := os.Getenv("INTERCOM_ACCESS_TOKEN")
	if token == "" {
		log.Fatal("INTERCOM_ACCESS_TOKEN environment variable is required")
	}

	ic := intercom.NewClient(token)

	collections := fetchAllCollections(ic)
	fmt.Fprintf(os.Stderr, "Fetched %d collections\n", len(collections))

	articles := fetchAllArticles(ic)
	fmt.Fprintf(os.Stderr, "Fetched %d articles\n", len(articles))

	// Separate top-level collections from sub-collections (formerly sections).
	// In API v2.10+, sections are collections with a non-null parent_id.
	topLevel := []intercom.Collection{}
	subByParent := map[string][]intercom.Collection{}
	for _, c := range collections {
		if c.ParentID == nil {
			topLevel = append(topLevel, c)
		} else {
			subByParent[*c.ParentID] = append(subByParent[*c.ParentID], c)
		}
	}

	// Group articles by parent_id
	articlesByParent := map[string][]ExportArticle{}
	for _, a := range articles {
		if a.State != "published" {
			continue
		}
		ea := ExportArticle{
			ID:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			Body:        a.Body,
			State:       a.State,
			URL:         a.URL,
		}
		parentID := strconv.FormatInt(a.ParentID, 10)
		articlesByParent[parentID] = append(articlesByParent[parentID], ea)
	}

	// Build the tree
	tree := make([]ExportCollection, 0, len(topLevel))
	for _, c := range topLevel {
		ec := ExportCollection{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
			URL:         c.URL,
			Articles:    articlesByParent[c.ID],
		}
		// Add sub-collections (formerly sections)
		for _, sub := range subByParent[c.ID] {
			eSub := ExportCollection{
				ID:       sub.ID,
				Name:     sub.Name,
				URL:      sub.URL,
				Articles: articlesByParent[sub.ID],
			}
			ec.SubCollections = append(ec.SubCollections, eSub)
		}
		tree = append(tree, ec)
	}

	output, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	if err := os.MkdirAll("export", 0755); err != nil {
		log.Fatalf("Error creating export directory: %v", err)
	}
	if err := os.WriteFile("export/faq.json", output, 0644); err != nil {
		log.Fatalf("Error writing export/faq.json: %v", err)
	}

	fmt.Fprintf(os.Stderr, "Written to export/faq.json (%d collections)\n", len(tree))
}

func fetchAllCollections(ic *intercom.Client) []intercom.Collection {
	var all []intercom.Collection
	page := int64(1)
	for {
		list, err := ic.Collections.List(intercom.PageParams{Page: page, PerPage: 50})
		if err != nil {
			log.Fatalf("Error fetching collections (page %d): %v", page, err)
		}
		all = append(all, list.Collections...)
		if page >= list.Pages.TotalPages {
			break
		}
		page++
	}
	return all
}

func fetchAllArticles(ic *intercom.Client) []intercom.Article {
	var all []intercom.Article
	page := int64(1)
	for {
		list, err := ic.Articles.List(intercom.PageParams{Page: page, PerPage: 50})
		if err != nil {
			log.Fatalf("Error fetching articles (page %d): %v", page, err)
		}
		all = append(all, list.Articles...)
		if page >= list.Pages.TotalPages {
			break
		}
		page++
	}
	return all
}

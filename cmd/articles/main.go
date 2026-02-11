package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	intercom "github.com/karnott/intercom-go"
)

type ExportArticle struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	State       string `json:"state"`
	URL         string `json:"url"`
}

type ExportSection struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	URL      string          `json:"url,omitempty"`
	Articles []ExportArticle `json:"articles"`
}

type ExportCollection struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Sections    []ExportSection `json:"sections,omitempty"`
	Articles    []ExportArticle `json:"articles,omitempty"`
}

func main() {
	token := os.Getenv("INTERCOM_ACCESS_TOKEN")
	if token == "" {
		log.Fatal("INTERCOM_ACCESS_TOKEN environment variable is required")
	}

	ic := intercom.NewClient(token, "")

	collections := fetchAllCollections(ic)
	fmt.Fprintf(os.Stderr, "Fetched %d collections\n", len(collections))

	sections := fetchAllSections(ic)
	fmt.Fprintf(os.Stderr, "Fetched %d sections\n", len(sections))

	articles := fetchAllArticles(ic)
	fmt.Fprintf(os.Stderr, "Fetched %d articles\n", len(articles))

	// Group articles by (parent_type, parent_id)
	collectionArticles := map[string][]ExportArticle{}
	sectionArticles := map[string][]ExportArticle{}

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
		switch a.ParentType {
		case "section":
			sectionArticles[parentID] = append(sectionArticles[parentID], ea)
		default:
			collectionArticles[parentID] = append(collectionArticles[parentID], ea)
		}
	}

	// Group sections by parent collection
	collectionSections := map[string][]ExportSection{}
	for _, s := range sections {
		es := ExportSection{
			ID:       s.ID,
			Name:     s.Name,
			URL:      s.URL,
			Articles: sectionArticles[s.ID],
		}
		if es.Articles == nil {
			es.Articles = []ExportArticle{}
		}
		collectionSections[strconv.FormatInt(s.ParentID, 10)] = append(collectionSections[strconv.FormatInt(s.ParentID, 10)], es)
	}

	// Build the tree
	tree := make([]ExportCollection, 0, len(collections))
	for _, c := range collections {
		ec := ExportCollection{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
			URL:         c.URL,
			Sections:    collectionSections[c.ID],
			Articles:    collectionArticles[c.ID],
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

func fetchAllSections(ic *intercom.Client) []intercom.Section {
	var all []intercom.Section
	page := int64(1)
	for {
		list, err := ic.Sections.List(intercom.PageParams{Page: page, PerPage: 50})
		if err != nil {
			log.Fatalf("Error fetching sections (page %d): %v", page, err)
		}
		all = append(all, list.Sections...)
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

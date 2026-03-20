package intercom

import (
	"encoding/json"
	"testing"
)

func TestStartingAfterParamUnmarshalObject(t *testing.T) {
	data := []byte(`{"starting_after": "WzE2OTYsIjY3OGIwZTQ1MDAwMDAwMDAwMDAwMDAwMCIsMl0=", "per_page": 25}`)
	var param StartingAfterParam
	if err := json.Unmarshal(data, &param); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if param.StartingAfter != "WzE2OTYsIjY3OGIwZTQ1MDAwMDAwMDAwMDAwMDAwMCIsMl0=" {
		t.Errorf("expected cursor string, got %q", param.StartingAfter)
	}
	if param.PerPage != 25 {
		t.Errorf("expected per_page 25, got %d", param.PerPage)
	}
}

func TestStartingAfterParamUnmarshalString(t *testing.T) {
	data := []byte(`"https://api.intercom.io/companies/123/contacts?page=2"`)
	var param StartingAfterParam
	if err := json.Unmarshal(data, &param); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if param.StartingAfter != "https://api.intercom.io/companies/123/contacts?page=2" {
		t.Errorf("expected URL string, got %q", param.StartingAfter)
	}
}

func TestCursorPagesUnmarshalWithStringNext(t *testing.T) {
	data := []byte(`{"type":"pages","page":1,"per_page":50,"total_pages":3,"next":"https://api.intercom.io/companies/123/contacts?page=2"}`)
	var pages CursorPages
	if err := json.Unmarshal(data, &pages); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pages.Next == nil {
		t.Fatal("expected Next to be non-nil")
	}
	if pages.Next.StartingAfter != "https://api.intercom.io/companies/123/contacts?page=2" {
		t.Errorf("expected URL in StartingAfter, got %q", pages.Next.StartingAfter)
	}
}

func TestCursorPagesUnmarshalWithObjectNext(t *testing.T) {
	data := []byte(`{"type":"pages","next":{"starting_after":"WzE2OTY="}}`)
	var pages CursorPages
	if err := json.Unmarshal(data, &pages); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pages.Next == nil {
		t.Fatal("expected Next to be non-nil")
	}
	if pages.Next.StartingAfter != "WzE2OTY=" {
		t.Errorf("expected cursor, got %q", pages.Next.StartingAfter)
	}
}

func TestCursorPagesUnmarshalWithNullNext(t *testing.T) {
	data := []byte(`{"type":"pages","page":1,"next":null}`)
	var pages CursorPages
	if err := json.Unmarshal(data, &pages); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pages.Next != nil {
		t.Errorf("expected Next to be nil, got %+v", pages.Next)
	}
}
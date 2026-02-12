package intercom

import (
	"io/ioutil"
	"testing"
)

func TestJobAPISaveContact(t *testing.T) {
	http := TestJobHTTPClient{t: t, expectedURI: "/bulk/users", fixtureFilename: "fixtures/job.json"}
	api := JobAPI{httpClient: &http}
	contact := Contact{ExternalID: "1234"}
	job := JobRequest{Items: []*JobItem{NewContactJobItem(&contact, JOB_POST)}, bulkType: "users"}
	savedJob, _ := api.save(&job)
	if savedJob.ID != "job_5ca1ab1eca11ab1e" {
		t.Errorf("Did not respond with correct job")
	}
}

func TestJobAPISaveEvent(t *testing.T) {
	http := TestJobHTTPClient{t: t, expectedURI: "/bulk/events", fixtureFilename: "fixtures/job.json"}
	api := JobAPI{httpClient: &http}
	event := Event{UserID: "1234"}
	job := JobRequest{Items: []*JobItem{NewEventJobItem(&event)}, bulkType: "events"}
	savedJob, _ := api.save(&job)
	if savedJob.ID != "job_5ca1ab1eca11ab1e" {
		t.Errorf("Did not respond with correct job")
	}
}

type TestJobHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	f               func(job *JobRequest)
	fixtureFilename string
	expectedURI     string
}

func (t *TestJobHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called")
	}
	if t.f != nil {
		t.f(body.(*JobRequest))
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

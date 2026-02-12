package intercom

import "testing"

func TestNewJob(t *testing.T) {
	repo := &TestJobRepository{t: t}
	repo.f = func(job *JobRequest) {
		if job.Items[0].Method != JOB_POST.String() {
			repo.t.Errorf("Wrong job method")
		}
		c := job.Items[0].Data.(*Contact)
		if c.Email != "foo@bar.com" {
			repo.t.Errorf("Wrong contact email")
		}
	}
	contact := Contact{Email: "foo@bar.com"}
	js := JobService{Repository: repo}
	js.NewUserJob(NewContactJobItem(&contact, JOB_POST))
}

func TestAppendJob(t *testing.T) {
	repo := &TestJobRepository{t: t}
	js := JobService{Repository: repo}
	newJob, _ := js.NewUserJob()

	repo.f = func(job *JobRequest) {
		if job.Items[0].Method != JOB_POST.String() {
			repo.t.Errorf("Wrong job method")
		}
		c := job.Items[0].Data.(*Contact)
		if c.Email != "foo@bar.com" {
			repo.t.Errorf("Wrong contact email")
		}
	}
	contact := Contact{Email: "foo@bar.com"}

	js.AppendUsers(newJob.ID, NewContactJobItem(&contact, JOB_POST))
}

type TestJobRepository struct {
	t *testing.T
	f func(job *JobRequest)
}

func (api *TestJobRepository) save(job *JobRequest) (JobResponse, error) {
	if api.f != nil {
		api.f(job)
	}
	return JobResponse{}, nil
}

func (api *TestJobRepository) find(id string) (JobResponse, error) {
	return JobResponse{}, nil
}

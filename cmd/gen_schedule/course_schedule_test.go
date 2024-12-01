package main

import (
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"testing"
)

func TestCourseSchedule(t *testing.T) {
	curriculum, err := data.ImportCurriculumFromCSV("Python II Programming Honors")
	if err != nil {
		t.Error()
	}
	terms, err := TermsLoader()
	if err != nil {
		t.Error()
	}
	if terms[0].ID == 0 {
		t.Error()
	}
	schedule, err := domain.NewCourseSchedule(terms[0], *curriculum)
	if err != nil {
		t.Error()
	}
	if schedule.Term.ID == 0 {
		t.Error()
	}

}

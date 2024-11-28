package main

import (
	"gh_static_portfolio/cmd/domain"
	"log"
	"testing"
)

func TestTermsLoader(t *testing.T) {
	terms, err := TermsLoader()
	if err != nil {
		t.Errorf("error loading terms: %s", err)
	}
	for _, term := range terms {
		log.Println("Term: ", term.Start.String(), " to ", term.End.String(), "ID: ", term.ID, " Name: ", term.Name)
		for _, date := range term.NonInstructionalDays {
			log.Println("NonIDay: ", date.String())
		}
		for _, date := range domain.InstructionDays(term) {
			log.Println("IDay: ", date.String()[:10])
		}
	}

}

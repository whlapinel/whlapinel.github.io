package data

import (
	"log"
	"testing"
)

func TestCurricSoaToOop(t *testing.T) {
	curric, err := ImportCurriculumFromCSV("Python II Programming Honors")
	if err != nil {
		t.Errorf("error importing curriculum: %s", err)
	}
	course := curricSoaToOop(*curric)
	for _, unit := range course.Units {
		log.Println("Unit:", unit.Number)
		log.Println("Lessons:", unit.Lessons)
		log.Println()
	}

}

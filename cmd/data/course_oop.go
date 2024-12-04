package data

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type courseRepo struct {
}

func NewCoursesRepo() domain.Repository[domain.Course] {
	return &courseRepo{}
}

func (c *courseRepo) All() ([]*domain.Course, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *courseRepo) WriteToCSV(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

func (c *courseRepo) ReadFromCSV() ([]*domain.Course, error) {

	instances, err := importInstancesFromCSV()
	if err != nil {
		return nil, err
	}
	var courses []*domain.Course
	for _, instance := range instances {
		courses = append(courses, courseInstanceSoaToOop(*instance))
	}
	return courses, nil

}

func (c *courseRepo) Save(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

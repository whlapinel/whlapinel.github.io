package data

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type courseRepo struct {
}

func NewCoursesRepo() domain.Repository[domain.CourseOOP] {
	return &courseRepo{}
}

func (c *courseRepo) All() ([]*domain.CourseOOP, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *courseRepo) WriteToCSV(course *domain.CourseOOP) error {
	return fmt.Errorf("not implemented")
}

func (c *courseRepo) ReadFromCSV() ([]*domain.CourseOOP, error) {

	instances, err := importInstancesFromCSV()
	if err != nil {
		return nil, err
	}
	var courses []*domain.CourseOOP
	for _, instance := range instances {
		courses = append(courses, courseInstanceSoaToOop(*instance))
	}
	return courses, nil

}

func (c *courseRepo) Save(course *domain.CourseOOP) error {
	return fmt.Errorf("not implemented")
}

package services

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type CourseService interface {
	CreateTemplate(template *domain.Course) error
	CreateInstance(course *domain.Course) error
	GetTemplates() ([]*domain.Course, error)
	GetInstances() ([]*domain.Course, error)
	ReadFromCSV() ([]*domain.Course, error)
}

type courseService struct {
	repo domain.CourseRepo
}

func NewCourseService(courseRepo domain.CourseRepo) CourseService {
	return courseService{repo: courseRepo}

}

func (svc courseService) CreateTemplate(course *domain.Course) error {
	return svc.repo.Save(course)
}
func (svc courseService) CreateInstance(course *domain.Course) error {
	return svc.repo.Save(course)
}

func (svc courseService) All() ([]*domain.Course, error) {
	return svc.repo.All()
}
func (svc courseService) GetTemplates() ([]*domain.Course, error) {
	return svc.repo.GetTemplates()
}
func (svc courseService) GetInstances() ([]*domain.Course, error) {
	return svc.repo.GetInstances()
}

func (svc courseService) ReadFromCSV() ([]*domain.Course, error) {
	courses, err := svc.repo.ReadFromCSV()
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		svc.repo.Save(course)
	}
	return courses, nil
}

func (svc courseService) Update(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

func (svc courseService) Delete(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

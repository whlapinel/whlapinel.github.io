package services

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type CourseService interface {
	CreatorService[domain.Course]
	UpdaterService[domain.Course]
	DeleterService[domain.Course]
	CSVService[domain.Course]
	GetAllService[domain.Course]
	GetTemplates() ([]*domain.Course, error)
}

type courseService struct {
	repo domain.CourseRepository
}

func NewCourseService(courseRepo domain.CourseRepository) CourseService {
	return courseService{repo: courseRepo}

}

func (svc courseService) Create(course *domain.Course) error {
	return svc.repo.Save(course)
}

func (svc courseService) All() ([]*domain.Course, error) {
	return svc.repo.All()
}
func (svc courseService) GetTemplates() ([]*domain.Course, error) {
	return svc.repo.GetTemplates()
}

func (svc courseService) ReadFromCSV() ([]*domain.Course, error) {
	return svc.repo.ReadFromCSV()
}

func (svc courseService) Update(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

func (svc courseService) Delete(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

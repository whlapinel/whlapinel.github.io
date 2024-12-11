package services

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type CourseOOPService interface {
	CreatorService[domain.CourseOOP]
	UpdaterService[domain.CourseOOP]
	DeleterService[domain.CourseOOP]
	CSVService[domain.CourseOOP]
	GetAllService[domain.CourseOOP]
}

type courseOOPService struct {
	repo domain.Repository[domain.CourseOOP]
}

func NewCourseOOPService(courseRepo domain.Repository[domain.CourseOOP]) CourseOOPService {
	return courseOOPService{repo: courseRepo}

}

func (svc courseOOPService) Create(course *domain.CourseOOP) error {
	return svc.repo.Save(course)
}

func (svc courseOOPService) All() ([]*domain.CourseOOP, error) {
	return svc.repo.All()
}

func (svc courseOOPService) ReadFromCSV() ([]*domain.CourseOOP, error) {
	return svc.repo.ReadFromCSV()
}

func (svc courseOOPService) Update(course *domain.CourseOOP) error {
	return fmt.Errorf("not implemented")
}

func (svc courseOOPService) Delete(course *domain.CourseOOP) error {
	return fmt.Errorf("not implemented")
}

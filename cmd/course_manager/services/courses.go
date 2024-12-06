package services

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type CourseService interface {
	CreatorService[domain.CourseSOA]
	UpdaterService[domain.CourseSOA]
	DeleterService[domain.CourseSOA]
	CSVService[domain.CourseSOA]
	GetAllService[domain.CourseSOA]
}

type courseService struct {
	repo domain.Repository[domain.CourseSOA]
}

func NewCourseService(courseRepo domain.Repository[domain.CourseSOA]) CourseService {
	return courseService{repo: courseRepo}

}

func (svc courseService) Create(course *domain.CourseSOA) error {
	return svc.repo.Save(course)
}

func (svc courseService) All() ([]*domain.CourseSOA, error) {
	return svc.repo.All()
}

func (svc courseService) ReadFromCSV() ([]*domain.CourseSOA, error) {
	return svc.repo.ReadFromCSV()
}

func (svc courseService) Update(course *domain.CourseSOA) error {
	return fmt.Errorf("not implemented")
}

func (svc courseService) Delete(course *domain.CourseSOA) error {
	return fmt.Errorf("not implemented")
}

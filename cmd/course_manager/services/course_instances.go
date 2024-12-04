package services

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type CourseInstanceService interface {
	Service[domain.CourseInstance]
}

type courseInstanceService struct {
	instanceRepo domain.Repository[domain.CourseInstance]
	courseRepo   domain.Repository[domain.CourseSOA]
	termRepo     domain.Repository[domain.Term]
}

func NewcourseInstanceService(
	instanceRepo domain.Repository[domain.CourseInstance],
	courseRepo domain.Repository[domain.CourseSOA],
	termRepo domain.Repository[domain.Term]) CourseInstanceService {
	return courseInstanceService{instanceRepo: instanceRepo, courseRepo: courseRepo, termRepo: termRepo}

}

func (svc courseInstanceService) Create(course *domain.CourseInstance) error {
	return svc.instanceRepo.Save(course)
}

func (svc courseInstanceService) All() ([]*domain.CourseInstance, error) {
	return svc.instanceRepo.All()
}

func (svc courseInstanceService) ReadFromCSV() ([]*domain.CourseInstance, error) {
	return svc.instanceRepo.ReadFromCSV()
}

func (svc courseInstanceService) Update(course *domain.CourseInstance) error {
	return fmt.Errorf("not implemented")
}

func (svc courseInstanceService) Delete(course *domain.CourseInstance) error {
	return fmt.Errorf("not implemented")
}

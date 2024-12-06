package services

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type CourseInstanceService interface {
	FetchOne(courseID int, termID int) ([]*domain.CourseInstance, error)
	CreatorService[domain.CourseInstance]
	UpdaterService[domain.CourseInstance]
	DeleterService[domain.CourseInstance]
	CSVService[domain.CourseInstance]
}

type courseInstanceService struct {
	instanceRepo domain.CourseInstanceRepository
	courseRepo   domain.Repository[domain.CourseSOA]
	termRepo     domain.Repository[domain.Term]
}

// FetchOne implements CourseInstanceService.
func (svc courseInstanceService) FetchOne(courseID int, termID int) ([]*domain.CourseInstance, error) {
	panic("unimplemented")
}

func NewcourseInstanceService(
	instanceRepo domain.CourseInstanceRepository,
	courseRepo domain.Repository[domain.CourseSOA],
	termRepo domain.Repository[domain.Term]) CourseInstanceService {
	return courseInstanceService{instanceRepo: instanceRepo, courseRepo: courseRepo, termRepo: termRepo}

}

func (svc courseInstanceService) Create(course *domain.CourseInstance) error {
	return svc.instanceRepo.Save(course)
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

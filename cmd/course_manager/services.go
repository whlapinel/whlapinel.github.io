package main

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
)

type Service[T any] interface {
	Create(*T) error
	All() ([]*T, error)
	Update(*T) error
	Delete(*T) error
}

type CourseService interface {
	Service[domain.CourseSOA]
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

func (svc courseService) Update(course *domain.CourseSOA) error {
	return fmt.Errorf("not implemented")
}

func (svc courseService) Delete(course *domain.CourseSOA) error {
	return fmt.Errorf("not implemented")
}

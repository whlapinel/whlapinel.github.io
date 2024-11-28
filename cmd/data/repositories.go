package data

import "gh_static_portfolio/cmd/domain"

type blogRepo struct {
}

func NewBlogRepo() domain.BlogRepository {
	return &blogRepo{}
}

func (r *blogRepo) GetAll() ([]*domain.Blog, error) {
	return blogs, nil
}

type educationRepo struct {
}

func NewEducationRepo() domain.EducationRepository {
	return &educationRepo{}
}

func (r *educationRepo) GetAll() ([]*domain.EducationItem, error) {
	return educationItems, nil
}

type projectRepo struct {
}

func NewProjectRepo() domain.ProjectRepository {
	return &projectRepo{}
}

func (r *projectRepo) GetAll() ([]*domain.ProjectItem, error) {
	return projectItems, nil
}

type workHistoryRepo struct {
}

func NewWorkHistoryRepo() domain.WorkHistoryRepository {
	return &workHistoryRepo{}
}

func (r *workHistoryRepo) GetAll() ([]*domain.WorkHistoryItem, error) {
	return workHistoryItems, nil
}

type skillRepo struct {
}

func NewSkillRepo() domain.SkillRepository {
	return &skillRepo{}
}

func (r *skillRepo) GetAll() ([]*domain.SkillItem, error) {
	return skillItems, nil
}

type courseRepo struct {
}

func NewCoursesRepo() domain.CourseRepository {
	return &courseRepo{}
}

func (c *courseRepo) GetAll() ([]domain.Course, error) {
	courses, err := Courses()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

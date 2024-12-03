package data

import "gh_static_portfolio/cmd/domain"

type projectRepo struct {
}

func NewProjectRepo() domain.ProjectRepository {
	return &projectRepo{}
}

func (r *projectRepo) GetAll() ([]*domain.ProjectItem, error) {
	return projectItems, nil
}
var projectItems = []*domain.ProjectItem{
	{
		Image:       "https://via.placeholder.com/150",
		Title:       "Project 1",
		Subtitle:    "Subtitle 1",
		Description: "This is a description of project 1.",
		Link:        "",
	},
}

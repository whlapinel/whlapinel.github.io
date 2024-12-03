package data

import "gh_static_portfolio/cmd/domain"

type skillRepo struct {
}

func NewSkillRepo() domain.SkillRepository {
	return &skillRepo{}
}

func (r *skillRepo) GetAll() ([]*domain.SkillItem, error) {
	return skillItems, nil
}

var skillItems = []*domain.SkillItem{
	{
		Title: "Backend Development using Go, Python, and Java",
	},
	{
		Title: "Frontend Development including React, Typescript/Javascript, HTML, and CSS",
	},
	{
		Title: "Containers including Docker",
	},
	{
		Title: "Operating systems including Windows and Linux",
	},
	{
		Title: "Version control including Git and GitHub",
	},
}

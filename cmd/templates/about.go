package templates

import (
	"gh_static_portfolio/cmd/data"

	"github.com/a-h/templ"
)

const AboutDir = "./docs/about/"

func AboutPages() []Templifier {
	return []Templifier{
		NewSkillsPage(data.SkillItems),
		NewEducationPage(data.EducationItems),
		NewWorkHistoryPage(data.WorkHistoryItems),
		NewProjectListPage(data.ProjectItems),
	}
}

type EducationPage struct {
	Items []*data.EducationItem
}

func NewEducationPage(items []*data.EducationItem) Templifier {
	return &EducationPage{
		Items: items,
	}
}

func (p *EducationPage) Templify() templ.Component {
	return EducationListComponent(p.Items)
}

func (p *EducationPage) Title() string {
	return "Education"
}

func (p *EducationPage) Directory() string {
	return AboutDir
}

type WorkHistoryPage struct {
	Items []*data.WorkHistoryItem
}

func NewWorkHistoryPage(items []*data.WorkHistoryItem) Templifier {
	return &WorkHistoryPage{
		Items: data.WorkHistoryItems,
	}
}

func (p *WorkHistoryPage) Templify() templ.Component {
	return WorkHistoryComponent(p.Items)
}

func (p *WorkHistoryPage) Title() string {
	return "Work History"
}

func (p *WorkHistoryPage) Directory() string {
	return "./docs/about/"
}

type ProjectListPage struct {
	Items []*data.ProjectItem
}

func NewProjectListPage(items []*data.ProjectItem) Templifier {
	return &ProjectListPage{
		Items: items,
	}
}

func (p *ProjectListPage) Templify() templ.Component {
	return ProjectListComponent(p.Items)
}

func (p *ProjectListPage) Title() string {
	return "Projects"
}

func (p *ProjectListPage) Directory() string {
	return AboutDir
}

type SkillsPage struct {
	Items []*data.SkillItem
}

func NewSkillsPage(items []*data.SkillItem) Templifier {
	return &SkillsPage{
		Items: items,
	}
}

func (p *SkillsPage) Templify() templ.Component {
	return SkillsListComponent(p.Items)
}

func (p *SkillsPage) Title() string {
	return "Skills"
}

func (p *SkillsPage) Directory() string {
	return AboutDir
}

package templates

import (
	"gh_static_portfolio/cmd/domain"

	"github.com/a-h/templ"
)

type aboutPage struct {
}

func NewAboutPage() Templifier {
	return &aboutPage{}
}

func (p *aboutPage) Templify() templ.Component {
	return AboutComponent()
}

func (p *aboutPage) Title() string {
	return "About"
}

func (p *aboutPage) Directory() string {
	return rootDir
}

type educationPage struct {
	Items []*domain.EducationItem
}

func NewEducationPage(items []*domain.EducationItem) Templifier {
	return &educationPage{
		Items: items,
	}
}

func (p *educationPage) Templify() templ.Component {
	return EducationListComponent(p.Items)
}

func (p *educationPage) Title() string {
	return "Education"
}

func (p *educationPage) Directory() string {
	return string(aboutDir)
}

type classesPage struct {
	Item *domain.EducationItem
}

func NewClassesPage(item *domain.EducationItem) Templifier {
	return &classesPage{
		Item: item,
	}
}

func (p *classesPage) Templify() templ.Component {
	return ClassListComponent(p.Item)
}

func (p *classesPage) Title() string {
	return p.Item.School
}

func (p *classesPage) Directory() string {
	return educationDir
}

type workHistoryPage struct {
	Items []*domain.WorkHistoryItem
}

func NewWorkHistoryPage(items []*domain.WorkHistoryItem) Templifier {
	return &workHistoryPage{
		Items: items,
	}
}

func (p *workHistoryPage) Templify() templ.Component {
	return WorkHistoryComponent(p.Items)
}

func (p *workHistoryPage) Title() string {
	return "Work History"
}

func (p *workHistoryPage) Directory() string {
	return "./docs/about/"
}

type projectsPage struct {
	Items []*domain.ProjectItem
}

func NewProjectsPage(items []*domain.ProjectItem) Templifier {
	return &projectsPage{
		Items: items,
	}
}

func (p *projectsPage) Templify() templ.Component {
	return ProjectListComponent(p.Items)
}

func (p *projectsPage) Title() string {
	return "Projects"
}

func (p *projectsPage) Directory() string {
	return string(aboutDir)
}

type skillsPage struct {
	Items []*domain.SkillItem
}

func NewSkillsPage(items []*domain.SkillItem) Templifier {
	return &skillsPage{
		Items: items,
	}
}

func (p *skillsPage) Templify() templ.Component {
	return SkillsListComponent(p.Items)
}

func (p *skillsPage) Title() string {
	return "Skills"
}

func (p *skillsPage) Directory() string {
	return string(aboutDir)
}

type personalPage struct {
}

func NewPersonalPage() Templifier {
	return &personalPage{}
}

func (p *personalPage) Templify() templ.Component {
	return PersonalComponent()
}

func (p *personalPage) Title() string {
	return "Personal"
}

func (p *personalPage) Directory() string {
	return string(aboutDir)
}

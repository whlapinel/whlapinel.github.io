package data

import "gh_static_portfolio/cmd/domain"

type workHistoryRepo struct {
}

func NewWorkHistoryRepo() domain.WorkHistoryRepository {
	return &workHistoryRepo{}
}

func (r *workHistoryRepo) GetAll() ([]*domain.WorkHistoryItem, error) {
	return workHistoryItems, nil
}

var workHistoryItems = []*domain.WorkHistoryItem{
	{
		Year:        "Aug 2024 - Present",
		Company:     "Phillip O. Berry Academy of Technology",
		Position:    "Teacher, Python Programming",
		Description: "Developed the school's first Python 2 course.",
	},
	{
		Year:        "Aug 2021 - Jun 2024",
		Company:     "South Mecklenburg High School",
		Position:    "Teacher, Earth & Environmental Science",
		Description: "Led the EES Professional Learning Community during the 2022-2023 school year.",
	},
	{
		Year:        "Jan 2020 - Jun 2021",
		Company:     "Mallard Creek High School",
		Position:    "Teacher, Earth & Environmental Science",
		Description: "Taught Earth & Environmental Science during the virtual learning phase of the pandemic.",
	},
	{
		Year:        "Aug 2001 - Oct 2015",
		Company:     "US Navy",
		Position:    "Various",
		Description: "10 years of active duty with 4 deployments aboard various ships.",
	},
}

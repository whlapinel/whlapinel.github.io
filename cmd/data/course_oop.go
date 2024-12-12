package data

import (
	"context"
	"database/sql"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"strconv"
)

type courseRepo struct {
	queries *database.Queries
}

// GetInstances implements domain.CourseRepository.
func (c *courseRepo) GetInstances() ([]*domain.Course, error) {
	panic("unimplemented")
}

// GetTemplates implements domain.CourseRepository.
func (c *courseRepo) GetTemplates() ([]*domain.Course, error) {
	dbCourses, err := c.queries.GetTemplates(context.Background())
	if err != nil {
		return nil, err
	}
	var courses []*domain.Course
	for _, dbCourse := range dbCourses {
		course := &domain.Course{
			ID:          int(dbCourse.CourseID),
			Name:        dbCourse.CourseName,
			Description: dbCourse.CourseDescr.String,
		}
		dbUnits, err := c.queries.GetUnits(context.Background(), dbCourse.CourseID)
		if err != nil {
			return nil, err
		}
		var units []domain.Unit
		for _, dbUnit := range dbUnits {
			unit := domain.Unit{
				ID:          int(dbUnit.ID),
				CourseID:    int(dbCourse.CourseID),
				Number:      int(dbUnit.Number),
				Name:        dbUnit.Name,
				Description: dbUnit.Description.String,
			}
			dbLessons, err := c.queries.GetLessons(context.Background(), dbUnit.ID)
			if err != nil {
				return nil, err
			}
			var lessons []domain.Lesson
			for _, dbLesson := range dbLessons {
				lesson := domain.Lesson{
					ID:          int(dbLesson.ID),
					Number:      int(dbLesson.Number),
					Name:        dbLesson.Name.String,
					Description: dbLesson.Description.String,
				}
				lessons = append(lessons, lesson)
			}
			unit.Lessons = lessons
			units = append(units, unit)
		}
		course.Units = units
		courses = append(courses, course)
	}
	return courses, nil
}

func NewCoursesRepo(db *database.Queries) domain.CourseRepository {
	return &courseRepo{queries: db}
}

func (c *courseRepo) All() ([]*domain.Course, error) {
	panic("not implemented")
}

func (c *courseRepo) WriteToCSV(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

func (c *courseRepo) ReadFromCSV() ([]*domain.Course, error) {

	instances, err := importInstancesFromCSV()
	if err != nil {
		return nil, err
	}
	var courses []*domain.Course
	for _, instance := range instances {
		courses = append(courses, courseInstanceSoaToOop(*instance))
	}
	return courses, nil

}

// Save updates the ID for Course, Units, and Lessons
func (c *courseRepo) Save(course *domain.Course) error {
	ctx := context.Background()
	dbCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{Name: course.Name})
	if err != nil {
		return err
	}
	course.ID = int(dbCourse.ID)
	for i, unit := range course.Units {
		var hasDescr = unit.Description == ""
		currUnit := database.Unit{
			CourseID: int64(course.ID),
			Number:   int64(unit.Number),
			Name:     unit.Name,
			Description: sql.NullString{
				String: unit.Description,
				Valid:  hasDescr,
			},
		}
		currUnit, err = c.queries.SaveUnit(ctx, database.SaveUnitParams{
			Number:      currUnit.Number,
			Name:        currUnit.Name,
			Description: currUnit.Description,
			CourseID:    currUnit.CourseID,
		})
		if err != nil {
			return err
		}
		unit.ID = int(currUnit.ID)
		for j, lesson := range unit.Lessons {
			log.Println("lesson name:", lesson.Name, "lesson number: ", lesson.Number)
			currLesson := database.Lesson{
				UnitID: int64(unit.ID),
				Number: int64(lesson.Number),
				Name: sql.NullString{
					String: lesson.Name,
					Valid:  lesson.Name != "",
				},
				Description: sql.NullString{
					String: lesson.Description,
					Valid:  lesson.Description != "",
				},
			}
			currLesson, err = c.queries.SaveLesson(ctx, database.SaveLessonParams{
				Number:      currLesson.Number,
				Name:        currLesson.Name,
				Description: currLesson.Description,
				UnitID:      currLesson.UnitID,
			})
			if err != nil {
				return fmt.Errorf("courseRepo.Save(), SaveLesson(): %v, unit id: %s, lesson#: %s", err, strconv.Itoa(int(currUnit.ID)), strconv.Itoa(int(currLesson.Number)))
			}
			course.Units[i].Lessons[j].ID = int(currLesson.ID)
		}

	}
	return nil

}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const deleteCourse = `-- name: DeleteCourse :one
DELETE FROM courses WHERE id = ?
RETURNING id, template_id, term_id, name, description
`

func (q *Queries) DeleteCourse(ctx context.Context, id int64) (Course, error) {
	row := q.db.QueryRowContext(ctx, deleteCourse, id)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.TemplateID,
		&i.TermID,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const deleteLesson = `-- name: DeleteLesson :one
DELETE FROM lessons WHERE id = ?
RETURNING id, unit_id, template_id, number, name, description
`

func (q *Queries) DeleteLesson(ctx context.Context, id int64) (Lesson, error) {
	row := q.db.QueryRowContext(ctx, deleteLesson, id)
	var i Lesson
	err := row.Scan(
		&i.ID,
		&i.UnitID,
		&i.TemplateID,
		&i.Number,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const deleteNonInstructDays = `-- name: DeleteNonInstructDays :one
DELETE FROM non_instruct_days WHERE id = ?
RETURNING id, term_id, date
`

func (q *Queries) DeleteNonInstructDays(ctx context.Context, id int64) (NonInstructDay, error) {
	row := q.db.QueryRowContext(ctx, deleteNonInstructDays, id)
	var i NonInstructDay
	err := row.Scan(&i.ID, &i.TermID, &i.Date)
	return i, err
}

const deleteTerm = `-- name: DeleteTerm :one
DELETE FROM terms WHERE id = ?
RETURNING id, name, start, "end"
`

func (q *Queries) DeleteTerm(ctx context.Context, id int64) (Term, error) {
	row := q.db.QueryRowContext(ctx, deleteTerm, id)
	var i Term
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Start,
		&i.End,
	)
	return i, err
}

const deleteUnit = `-- name: DeleteUnit :one
DELETE FROM units WHERE id = ?
RETURNING id, course_id, template_id, number, sequence, name, description
`

func (q *Queries) DeleteUnit(ctx context.Context, id int64) (Unit, error) {
	row := q.db.QueryRowContext(ctx, deleteUnit, id)
	var i Unit
	err := row.Scan(
		&i.ID,
		&i.CourseID,
		&i.TemplateID,
		&i.Number,
		&i.Sequence,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getCourses = `-- name: GetCourses :many
SELECT id, template_id, term_id, name, description FROM courses
`

func (q *Queries) GetCourses(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.ID,
			&i.TemplateID,
			&i.TermID,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDate = `-- name: GetDate :one
SELECT id, term_id, day_number, date FROM dates WHERE date = ?
`

func (q *Queries) GetDate(ctx context.Context, date string) (Date, error) {
	row := q.db.QueryRowContext(ctx, getDate, date)
	var i Date
	err := row.Scan(
		&i.ID,
		&i.TermID,
		&i.DayNumber,
		&i.Date,
	)
	return i, err
}

const getInstances = `-- name: GetInstances :many
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr,
  c.term_id
FROM 
  courses c
WHERE 
  c.template_id IS NOT NULL
`

type GetInstancesRow struct {
	CourseID    int64
	CourseName  string
	CourseDescr sql.NullString
	TermID      sql.NullInt64
}

func (q *Queries) GetInstances(ctx context.Context) ([]GetInstancesRow, error) {
	rows, err := q.db.QueryContext(ctx, getInstances)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInstancesRow
	for rows.Next() {
		var i GetInstancesRow
		if err := rows.Scan(
			&i.CourseID,
			&i.CourseName,
			&i.CourseDescr,
			&i.TermID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLessons = `-- name: GetLessons :many
SELECT
  l.id,
  l.number,
  l.name,
  l.description
FROM
  lessons l
WHERE
  l.unit_id = ?
`

type GetLessonsRow struct {
	ID          int64
	Number      int64
	Name        sql.NullString
	Description sql.NullString
}

func (q *Queries) GetLessons(ctx context.Context, unitID int64) ([]GetLessonsRow, error) {
	rows, err := q.db.QueryContext(ctx, getLessons, unitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLessonsRow
	for rows.Next() {
		var i GetLessonsRow
		if err := rows.Scan(
			&i.ID,
			&i.Number,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemplates = `-- name: GetTemplates :many
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_descr
FROM 
  courses c
WHERE 
  c.template_id IS NULL
`

type GetTemplatesRow struct {
	CourseID    int64
	CourseName  string
	CourseDescr sql.NullString
}

func (q *Queries) GetTemplates(ctx context.Context) ([]GetTemplatesRow, error) {
	rows, err := q.db.QueryContext(ctx, getTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTemplatesRow
	for rows.Next() {
		var i GetTemplatesRow
		if err := rows.Scan(&i.CourseID, &i.CourseName, &i.CourseDescr); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnits = `-- name: GetUnits :many
SELECT
  u.id,
  u.number,
  u.name,
  u.description
FROM
  units u
WHERE
  u.course_id = ?
`

type GetUnitsRow struct {
	ID          int64
	Number      int64
	Name        string
	Description sql.NullString
}

func (q *Queries) GetUnits(ctx context.Context, courseID int64) ([]GetUnitsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUnits, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUnitsRow
	for rows.Next() {
		var i GetUnitsRow
		if err := rows.Scan(
			&i.ID,
			&i.Number,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveCourse = `-- name: SaveCourse :one
INSERT INTO courses (
  template_id, term_id, name, description
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, template_id, term_id, name, description
`

type SaveCourseParams struct {
	TemplateID  sql.NullInt64
	TermID      sql.NullInt64
	Name        string
	Description sql.NullString
}

func (q *Queries) SaveCourse(ctx context.Context, arg SaveCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, saveCourse,
		arg.TemplateID,
		arg.TermID,
		arg.Name,
		arg.Description,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.TemplateID,
		&i.TermID,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const saveDate = `-- name: SaveDate :one
INSERT INTO dates (
  term_id, day_number, date
) VALUES (
  ?, ?, ?
)
RETURNING id, term_id, day_number, date
`

type SaveDateParams struct {
	TermID    int64
	DayNumber int64
	Date      string
}

func (q *Queries) SaveDate(ctx context.Context, arg SaveDateParams) (Date, error) {
	row := q.db.QueryRowContext(ctx, saveDate, arg.TermID, arg.DayNumber, arg.Date)
	var i Date
	err := row.Scan(
		&i.ID,
		&i.TermID,
		&i.DayNumber,
		&i.Date,
	)
	return i, err
}

const saveLesson = `-- name: SaveLesson :one
INSERT INTO lessons (
  number, name, description, unit_id, template_id
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING id, unit_id, template_id, number, name, description
`

type SaveLessonParams struct {
	Number      int64
	Name        sql.NullString
	Description sql.NullString
	UnitID      int64
	TemplateID  sql.NullInt64
}

func (q *Queries) SaveLesson(ctx context.Context, arg SaveLessonParams) (Lesson, error) {
	row := q.db.QueryRowContext(ctx, saveLesson,
		arg.Number,
		arg.Name,
		arg.Description,
		arg.UnitID,
		arg.TemplateID,
	)
	var i Lesson
	err := row.Scan(
		&i.ID,
		&i.UnitID,
		&i.TemplateID,
		&i.Number,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const saveLessonDate = `-- name: SaveLessonDate :one
INSERT INTO lesson_dates (
  lesson_id, date_id
) VALUES (
  ?, ?
)
RETURNING lesson_id, date_id
`

type SaveLessonDateParams struct {
	LessonID int64
	DateID   int64
}

func (q *Queries) SaveLessonDate(ctx context.Context, arg SaveLessonDateParams) (LessonDate, error) {
	row := q.db.QueryRowContext(ctx, saveLessonDate, arg.LessonID, arg.DateID)
	var i LessonDate
	err := row.Scan(&i.LessonID, &i.DateID)
	return i, err
}

const saveTerm = `-- name: SaveTerm :one
INSERT INTO terms (
  name, start, end
) VALUES (
  ?, ?, ?
)
RETURNING id, name, start, "end"
`

type SaveTermParams struct {
	Name  string
	Start string
	End   string
}

func (q *Queries) SaveTerm(ctx context.Context, arg SaveTermParams) (Term, error) {
	row := q.db.QueryRowContext(ctx, saveTerm, arg.Name, arg.Start, arg.End)
	var i Term
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Start,
		&i.End,
	)
	return i, err
}

const saveUnit = `-- name: SaveUnit :one
INSERT INTO units (
  number, sequence, name, description, course_id, template_id
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING id, course_id, template_id, number, sequence, name, description
`

type SaveUnitParams struct {
	Number      int64
	Sequence    int64
	Name        string
	Description sql.NullString
	CourseID    int64
	TemplateID  sql.NullInt64
}

func (q *Queries) SaveUnit(ctx context.Context, arg SaveUnitParams) (Unit, error) {
	row := q.db.QueryRowContext(ctx, saveUnit,
		arg.Number,
		arg.Sequence,
		arg.Name,
		arg.Description,
		arg.CourseID,
		arg.TemplateID,
	)
	var i Unit
	err := row.Scan(
		&i.ID,
		&i.CourseID,
		&i.TemplateID,
		&i.Number,
		&i.Sequence,
		&i.Name,
		&i.Description,
	)
	return i, err
}

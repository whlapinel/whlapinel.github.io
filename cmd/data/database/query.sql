-- name: GetCourses :many
SELECT
  c.id as course_id,
  c.name as course_name,
  c.description as course_description,
  u.number as unit_number,
  u.name as unit_name,
  l.number as lesson_number,
  l.name as lesson_name
FROM
  courses c
JOIN
  units u ON u.course_id = c.id AND u.instance_id = NULL
JOIN
  lessons l ON l.unit_id = u.id;

-- name: GetCourseInstances :many
SELECT
  ci.id as instance_id,
  t.id as term_id,
  t.name as term_name,
  c.name as course_name, 
  u.number as unit_number, 
  u.name as unit_name, 
  l.number as lesson_number, 
  l.name as lesson_name, 
  d.date, 
  d.number as day_number
FROM
  course_instances as ci
JOIN
  courses as c ON c.id = ci.course_id
JOIN
  units as u on u.instance_id = ci.id
JOIN
  terms as t on t.id = ci.term_id
JOIN
  lessons as l on l.unit_id = u.id
JOIN
  dates as d ON d.lesson_id = l.id
WHERE
  c.id = ? AND term_id = ?
ORDER BY
  d.number;

-- name: SaveCourse :one
INSERT INTO courses (
  name, description
) VALUES (
  ?, ?
)
RETURNING *;

-- name: SaveCourseInstance :one
INSERT INTO course_instances (
  course_id, term_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: SaveUnit :one
INSERT INTO units (
  number, name, description, course_id, instance_id
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: SaveLesson :one
INSERT INTO lessons (
  number, name, description, unit_id
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: SaveTerm :one
INSERT INTO terms (
  name, start, end
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: DeleteLesson :one
DELETE FROM lessons WHERE id = ?
RETURNING *;

-- name: DeleteUnit :one
DELETE FROM units WHERE id = ?
RETURNING *;

-- name: DeleteCourseInstance :one
DELETE FROM course_instances WHERE id = ?
RETURNING *;

-- name: DeleteTerm :one
DELETE FROM terms WHERE id = ?
RETURNING *;

-- name: DeleteNonInstructDays :one
DELETE FROM non_instruct_days WHERE id = ?
RETURNING *;

-- name: DeleteCourse :one
DELETE FROM courses WHERE id = ?
RETURNING *;







-- name: GetCourses :many
SELECT
    c.id as course_id,
    c.name as course_name,
    d.day AS day_number,
    u.number AS unit_number,
    u.name AS unit_name,
    l.number AS lesson_number,
    l.name AS lesson_name
FROM 
    day_numbers d
JOIN 
    lessons l ON d.lesson_id = l.id
JOIN 
    units u ON l.unit_id = u.id
JOIN 
    courses c ON u.course_id = c.id
ORDER BY 
    d.day;


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
  number, name, description, course_id 
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: SaveLesson :one
INSERT INTO lessons (
  number, name, description, unit_id
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: SaveDayNumber :one
INSERT INTO day_numbers (
  lesson_id, day
) VALUES (
  ?, ?
)
RETURNING *;

-- name: SaveTerm :one
INSERT INTO terms (
  start, end
) VALUES (
  ?, ?
)
RETURNING *;

-- name: SaveDate :one
INSERT INTO dates (
  course_instance_id, lesson_id, date
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: DeleteDayNumber :one
DELETE FROM day_numbers WHERE id = ?
RETURNING *;

-- name: DeleteLesson :one
DELETE FROM lessons WHERE id = ?
RETURNING *;

-- name: DeleteUnit :one
DELETE FROM units WHERE id = ?
RETURNING *;

-- name: DeleteDate :one
DELETE FROM dates WHERE course_instance_id = ? and lesson_id = ?
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







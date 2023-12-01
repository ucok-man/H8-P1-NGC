-- Retrieve the list of all students enrolled in a specific course.
SELECT
	students.name,
    students.email,
    courses.title
FROM students
INNER JOIN student_course USING(student_id)
INNER JOIN courses USING(course_id)
WHERE courses.title = 'Fisika Course';

-- Find all the courses a particular student has registered for.
SELECT
    courses.title
FROM students
INNER JOIN student_course USING(student_id)
INNER JOIN courses USING(course_id)

-- Get the preferred schedule of a student for a specific course.
SELECT
	students.name,
    courses.title,
    student_course.prefered_schedule
FROM students
INNER JOIN student_course USING(student_id)
INNER JOIN courses USING(course_id);
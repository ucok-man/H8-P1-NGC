INSERT INTO students(name, email, major, year_of_study)
VALUES 
    ('Ali', 'ali@example.com', 'Biologi', 2021),
    ('Budi', 'budi@example.com', 'Fisika', 2021),
    ('Zilong', 'zilong@example.com', 'Fisika', 2021),
    ('Yu Zong', 'yuzong@example.com', 'Manajemen', 2021),
    ('Eren', 'eren@example.com', 'Biologi', 2021);

INSERT INTO courses(title, instructor, schedule, credit_hours)
VALUES 
    ('Biologi Course', 'Instructor 1', '2023-01-1', 70),
    ('Fisika Course', 'Instructor 2', '2023-02-5', 70),
    ('Manajemen Course', 'Instructor 3', '2023-02-10', 70),
    ('IT Course', 'Instructor 4', '2023-03-15', 70),
    ('Seni Rupa Course', 'Instructor 5', '2023-03-20', 70);

INSERT INTO student_course(student_id, course_id, prefered_schedule)
VALUES 
    (1, 1, '2023-01-5'),
    (5, 1, '2023-01-10'),
    (2, 2, '2023-02-10'),
    (3, 2, '2023-02-10'),
    (4, 5, '2023-03-25');

    

CREATE DATABASE IF NOT EXISTS student_course;
USE student_course;

CREATE TABLE students(
    student_id INT AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    major VARCHAR(255) NOT NULL,
    year_of_study YEAR NOT NULL,

    PRIMARY KEY (student_id),
    UNIQUE(email)
);

CREATE TABLE courses(
    course_id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    instructor VARCHAR(255) NOT NULL,
    schedule DATE NOT NULL,
    credit_hours INT NOT NULL,

   PRIMARY KEY(course_id),
   CHECK (credit_hours > 0)  
);

CREATE TABLE student_course(
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    prefered_schedule DATE NOT NULL,

    FOREIGN KEY (student_id) REFERENCES students(student_id)
        ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
        ON DELETE CASCADE
);
BEGIN;

CREATE TABLE students (
    id int NOT NULL,
    lastName varchar(100) NOT NULL,
    firstName varchar(100) NOT NULL,
    age int NOT NULL,
    CONSTRAINT student_key PRIMARY KEY (id)
);

CREATE TABLE classes (
    id int NOT NULL,
    title varchar(100) NOT NULL,
    description varchar(100) NOT NULL,
    CONSTRAINT person_key PRIMARY KEY (id)
);

CREATE TABLE student_classes(
    studentId int NOT NULL,
    classId int NOT NULL,
    CONSTRAINT student_class_key PRIMARY KEY (studentId,classId),
    CONSTRAINT student_fk FOREIGN KEY (studentId) REFERENCES students(id),
    CONSTRAINT class_fk FOREIGN KEY (classId) REFERENCES classes(id)
);

COMMIT;
BEGIN;

CREATE TABLE IF NOT EXISTS students (
    id int NOT NULL,
    lastName varchar(100) NOT NULL,
    firstName varchar(100) NOT NULL,
    age int NOT NULL,
    CONSTRAINT student_key PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS classes (
    id int NOT NULL,
    title varchar(100) NOT NULL,
    description varchar(100) NOT NULL,
    CONSTRAINT person_key PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS student_classes(
    student_id int NOT NULL,
    class_id int NOT NULL,
    CONSTRAINT student_class_key PRIMARY KEY (student_id,class_id),
    CONSTRAINT student_fk FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
    CONSTRAINT class_fk FOREIGN KEY (class_id) REFERENCES classes(id) ON DELETE CASCADE
);

COMMIT;
CREATE TABLE IF NOT EXISTS students (
        rollno integer primary key ,
        name text,
        dept_no integer references departments(id)
);
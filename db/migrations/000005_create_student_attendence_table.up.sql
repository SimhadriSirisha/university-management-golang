CREATE TABLE IF NOT EXISTS student_attendance (
    id SERIAL,
    rollno integer,
    date date,
    login_time time ,
    logout_time time,
    primary key (rollno, date)
);
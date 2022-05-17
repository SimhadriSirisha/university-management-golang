CREATE TABLE IF NOT EXISTS student_attendance (
    id SERIAL PRIMARY KEY,
    rollno integer not null,
    date date,
    login_time time ,
    logout_time time
);
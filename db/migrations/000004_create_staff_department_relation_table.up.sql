CREATE TABLE IF NOT EXISTS staff_department (
        id SERIAL PRIMARY KEY,
        staffId integer not null,
        deptId integer not null,
        constraint fk_staff_department_staff foreign key (staffId) references staffs(id),
        constraint fk_staff_department_dept foreign key (deptId) references departments(id)
);
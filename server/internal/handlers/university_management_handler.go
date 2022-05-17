package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"university-management-golang/db/connection"
	um "university-management-golang/protoclient/university_management"
)

type universityManagementServer struct {
	um.UniversityManagementServiceServer
	connectionManager connection.DatabaseConnectionManager
}

func (u *universityManagementServer) GetDepartment(ctx context.Context, request *um.GetDepartmentRequest) (*um.GetDepartmentResponse, error) {
	connection, err := u.connectionManager.GetConnection()
	//defer u.connectionManager.CloseConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	var department um.Department
	connection.GetSession().Select("id", "name").From("departments").Where("id = ?", request.GetId()).LoadOne(&department)

	_, err = json.Marshal(&department)
	 if err != nil {
		 log.Fatalf("Error while marshaling %+v", err)
	 }

	return &um.GetDepartmentResponse{Department: &um.Department{
		Id:   department.Id,
		Name: department.Name,
	}}, nil
}

func (u *universityManagementServer) GetStudentsForDepartment(ctx context.Context, request *um.GetStudentsForDepartmentRequest) (*um.GetStudentsForDepartmentResponse, error) {
	connection, err := u.connectionManager.GetConnection()
	//defer u.connectionManager.CloseConnection()

	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	var students []*um.Student
	_,err = connection.GetSession().Select("rollno", "name", "dept_no").From("students").Where("dept_no = ?", request.GetId()).Load(&students)
	if err != nil {
		log.Fatalf("Error occured while fetching: %v\n",err);
	}

	_, err = json.Marshal(&students)

	return &um.GetStudentsForDepartmentResponse{
		Students: students,
	},err
}

func (u *universityManagementServer)  GetStaffsTeachingToStudent(ctx context.Context,request *um.GetStaffsTeachingToStudentRequest) (*um.GetStaffsTeachingToStudentResponse, error) {
	connection, err := u.connectionManager.GetConnection()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	var dept_id int32
	_,err = connection.GetSession().Select("dept_no").From("students").Where("rollno = ?", request.GetStudentId()).Load(&dept_id)
	if err != nil {
		log.Fatalf("Error occured while fetching dept id: %v\n",err)
	}

	var staffIds []int32
	_,err = connection.GetSession().Select("staffid").From("staff_department").Where("deptid = ?", dept_id).Distinct().Load(&staffIds)
	if err != nil {
		log.Fatalf("Error occured while fetching staff id's: %v\n",err)
	}

	var staffs []*um.Staff
	_,err = connection.GetSession().Select("id", "name").From("staffs").Where("id in ?", staffIds).Load(&staffs)

	return &um.GetStaffsTeachingToStudentResponse{Staffs: staffs},err
}

func (u *universityManagementServer) InsertLoginTime(ctx context.Context, request *um.InsertLoginTimeRequest) (*um.InsertLoginTimeResponse, error) {
	connection, err := u.connectionManager.GetConnection()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	now := time.Now()
	rq := request.GetRollno()
	var rn int32

	_,err = connection.GetSession().Select("rollno").From("student_attendance").Where("rollno = ?", rq).Load(&rn)
	if rn == rq {
		_,err = connection.GetSession().Update("student_attendance").Set("login_time",  now.Format("15:04:05")).Where("rollno = ?", rq).Exec()
	}else {
		_,err = connection.GetSession().InsertInto("student_attendance").Columns("rollno", "date", "login_time").Values(rq,now.Format("Jan 2, 2006"), now.Format("15:04:05")).Exec()
	}

	if err != nil {
		return &um.InsertLoginTimeResponse{Msg: "Unsuccessful in inserting login time"},err
	}
	return &um.InsertLoginTimeResponse{Msg: "Successfully login time loaded" }, err
}

func (u *universityManagementServer) InsertLogoutTime(ctx context.Context, request *um.InsertLogoutTimeRequest) (*um.InsertLogoutTimeResponse, error){
	connection, err := u.connectionManager.GetConnection()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	now := time.Now()
	rq := request.GetRollno()
	var rn int32

	_,err = connection.GetSession().Select("rollno").From("student_attendance").Where("rollno = ?", rq).Load(&rn)

	if rn == rq {
		_,err = connection.GetSession().Update("student_attendance").Set("logout_time",  now.Format("15:04:05")).Where("rollno = ?", rq).Exec()
	}else {
		return &um.InsertLogoutTimeResponse{Msg: "Insert Login time first"},nil
	}

	if err != nil {
		return &um.InsertLogoutTimeResponse{Msg: "Unsuccessful in inserting logout time"},err
	}
	return &um.InsertLogoutTimeResponse{Msg: "Successfully logout time loaded" }, err
}

func NewUniversityManagementHandler(connectionmanager connection.DatabaseConnectionManager) um.UniversityManagementServiceServer {
	return &universityManagementServer {
		connectionManager: connectionmanager,
	}
}
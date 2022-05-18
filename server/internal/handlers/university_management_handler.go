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

	_,err = connection.GetSession().
		InsertBySql("Insert into student_attendance(rollno,date,login_time) values(?,?,?) " +
			"ON CONFLICT (rollno,date) DO NOTHING",
			rq,now.Format("Jan 2, 2006"), now.Format("15:04:05")).Returning("login_time").Exec()

	if err != nil {
		return &um.InsertLoginTimeResponse{Msg: "Unsuccessful login"},err
	}
	return &um.InsertLoginTimeResponse{Msg: "Login Successful" }, err
}

func (u *universityManagementServer) InsertLogoutTime(ctx context.Context, request *um.InsertLogoutTimeRequest) (*um.InsertLogoutTimeResponse, error){
	connection, err := u.connectionManager.GetConnection()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	now := time.Now()
	rq := request.GetRollno()

	_,err = connection.GetSession().Update("student_attendance").Set("logout_time",  now.Format("15:04:05")).Where("rollno = ? and date = ?", rq, now.Format("Jan 2, 2006")).Exec()

	if err != nil {
		return &um.InsertLogoutTimeResponse{Msg: "Logout Unsuccessful"},err
	}
	return &um.InsertLogoutTimeResponse{Msg: "Logout Successful" }, err
}

func NewUniversityManagementHandler(connectionmanager connection.DatabaseConnectionManager) um.UniversityManagementServiceServer {
	return &universityManagementServer {
		connectionManager: connectionmanager,
	}
}
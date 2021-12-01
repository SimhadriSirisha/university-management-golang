package handlers

import (
	"context"
	"encoding/json"
	"log"
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
	if err != nil {
		log.Fatalf("Error while marshaling %+v", err)
	}

	return &um.GetStudentsForDepartmentResponse{
		Students: students,
	},nil
}

func NewUniversityManagementHandler(connectionmanager connection.DatabaseConnectionManager) um.UniversityManagementServiceServer {
	return &universityManagementServer {
		connectionManager: connectionmanager,
	}
}
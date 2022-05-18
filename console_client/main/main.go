package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"university-management-golang/protoclient/university_management"
)

const (
	host = "localhost"
	port = "2345"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error occured %+v", err)
	}
	client := university_management.NewUniversityManagementServiceClient(conn)
	var departmentID int32 = 101
	departmentResponse, err := client.GetDepartment(context.TODO(), &university_management.GetDepartmentRequest{Id: departmentID})
	if err != nil {
		log.Fatalf("Error occured while fetching department for id %d,err: %+v", departmentID, err)
	}
	log.Println(departmentResponse)

	studentRequest := &university_management.GetStudentsForDepartmentRequest{Id: departmentID}
	studentsResponse, err := client.GetStudentsForDepartment(context.Background(), studentRequest)
	if err != nil {
		log.Fatalf("Error occured while fetching students of department id %d,err: %+v", departmentID, err)
	}
	log.Println(studentsResponse)

	var studentId int32 = 1
	staffRequest := &university_management.GetStaffsTeachingToStudentRequest{StudentId: studentId}
	staffResponse,err := client.GetStaffsTeachingToStudent(context.Background(), staffRequest)
	if err != nil {
		log.Fatalf("Error occured while fetching the staffs teaching the student having rollno %d, err: %v", studentId, err)
	}
	log.Println(staffResponse)

	studentLoginAttendanceRequest := &university_management.InsertLoginTimeRequest{Rollno: 1}
	loginRes,err := client.InsertLoginTime(context.Background(), studentLoginAttendanceRequest)
	if err != nil {
		log.Fatalf("Error occured while inserting login time of student having rollno %d, err: %v", studentId, err)
	}
	log.Println(loginRes)

	loginRes,err = client.InsertLoginTime(context.Background(), studentLoginAttendanceRequest)
	if err != nil {
		log.Fatalf("Error occured while inserting login time of student having rollno %d, err: %v", studentId, err)
	}
	log.Println(loginRes)


	studentLogoutAttendanceRequest := &university_management.InsertLogoutTimeRequest{Rollno: 1}
	logoutRes,err := client.InsertLogoutTime(context.Background(), studentLogoutAttendanceRequest)
	if err != nil {
		log.Fatalf("Error occured while inserting logout time of student having rollno %d, err: %v", studentId, err)
	}
	log.Println(logoutRes)

	logoutRes,err = client.InsertLogoutTime(context.Background(), studentLogoutAttendanceRequest)
	if err != nil {
		log.Fatalf("Error occured while inserting logout time of student having rollno %d, err: %v", studentId, err)
	}
	log.Println(logoutRes)
}

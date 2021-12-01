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
}

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	migrations "university-management-golang/db"
	"university-management-golang/db/connection"
	um "university-management-golang/protoclient/university_management"
	"university-management-golang/server/internal/handlers"
)

const port = "2345"


//db
const (
	username = "postgres"
	password = "admin"
	host = "localhost"
	dbPort   = "5436"
	dbName = "postgres"
	schema = "public"
)

func main() {
	err := migrations.MigrationsUp(username, password, host, dbPort, dbName, schema)
	if err != nil {
		log.Fatalf("Failed to migrate, err: %+v\n", err)
	}

	connectionmanager := &connection.DatabaseConnectionManagerImpl{
			&connection.DBConfig{
				host,dbPort,username,password,dbName,schema,
			},
			nil,
	}

	insertSeedData(connectionmanager)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen to port: %s, err: %+v\n", port, err)
	}
	log.Printf("Starting to listen on port: %s\n", port)

	um.RegisterUniversityManagementServiceServer(grpcServer, handlers.NewUniversityManagementHandler(connectionmanager))
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to start GRPC Server: %+v\n", err)
	}
}

func insertSeedData(connectionManager connection.DatabaseConnectionManager) {
	connection, err := connectionManager.GetConnection()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	log.Println("Cleaning up students table")
	_, err = connection.GetSession().DeleteFrom("students").Exec()
	log.Println("Cleaning up departments table")
	_, err = connection.GetSession().DeleteFrom("departments").Exec()
	if err != nil {
		log.Fatalf("Could not delete from department table. Err: %+v", err)
	}

	log.Println("Inserting into departments table")
	_, err = connection.GetSession().InsertInto("departments").Columns("id","name").
		Values(101, "Computer Science").Exec()
	_, err = connection.GetSession().InsertInto("departments").Columns("id","name").
		Values(102, "Information Technology").Exec()

	log.Println("Inserting into students table")
	_,err = connection.GetSession().InsertInto("students").Columns("rollno", "name", "dept_no").
		Values(1, "Justin", 101).Exec()
	_,err = connection.GetSession().InsertInto("students").Columns("rollno", "name", "dept_no").
		Values(2, "Crystal", 101).Exec()
	_,err = connection.GetSession().InsertInto("students").Columns("rollno", "name", "dept_no").
		Values(3, "James", 102).Exec()
	_,err = connection.GetSession().InsertInto("students").Columns("rollno", "name", "dept_no").
		Values(4, "Sandy", 102).Exec()

	if err != nil {
		log.Fatalf("Could not insert into department table. Err: %+v", err)
	}

	//defer connectionManager.CloseConnection()
}

package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func CreateMYSQL(name string, pw string, host string, port string) {
	log.Println(name, pw, host, port)
	connect(name, pw, host, port)
	createDB()
	chooseDB()
	createTables()
	CloseConnection()
	DB = nil
}

func connectMYSQL(name string, pw string, host string, port string) {
	connect(name, pw, host, port)
	chooseDB()
}

func GetDB(name string, pw string, host string, port string) *sql.DB {
	if DB == nil {
		connectMYSQL(name, pw, host, port)
	}
	return DB
}

func connect(name string, pw string, host string, port string) {
	dataSource := name + ":" + pw + "@tcp(" + host + ":" + port + ")/"
	dbConnect, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Println(err.Error())
	}
	DB = dbConnect
}

func createDB() {
	_, err := DB.Exec("CREATE DATABASE IF NOT EXISTS coyoMock")
	if err != nil {
		log.Println(err.Error())
	}
}

func chooseDB() {
	_, err := DB.Exec("USE coyoMock")
	if err != nil {
		log.Println(err.Error())
	}
}

func createTables() {
	createGroupTable()
	createRoleTable()
	createUserTable()
	createUserGroupTable()
	createUserRoleTable()
}

func createRoleTable() {
	stmt, err := DB.Prepare("CREATE TABLE IF NOT EXISTS coyoRole(id int NOT NULL AUTO_INCREMENT, roleId varchar(100), name varchar(50), PRIMARY KEY (id));")
	if err != nil {
		log.Println(err.Error())
	}
	stmtExec(stmt)
}

func createGroupTable() {
	stmt, err := DB.Prepare("CREATE TABLE IF NOT EXISTS coyoGroup(id int NOT NULL AUTO_INCREMENT, groupId varchar(100), name varchar(50), PRIMARY KEY (id));")
	if err != nil {
		log.Println(err.Error())
	}
	stmtExec(stmt)
}

//These fields (according to createUser Request) are ignored:
/*
	remoteLogonName,
	persistedDisplayName,
	welcomeMail,
	generatePassword,
	language,
	temporaryPassword,
	initialUser
*/
func createUserTable() {
	stmt, err := DB.Prepare("CREATE TABLE IF NOT EXISTS coyoUser(id int NOT NULL AUTO_INCREMENT, email varchar(30), loginName varchar(30), firstname varchar(30), lastname varchar(30), active boolean, superadmin boolean, password varchar(50), PRIMARY KEY (id));")
	if err != nil {
		log.Println(err.Error())
	}
	stmtExec(stmt)
}

func createUserRoleTable() {
	stmt, err := DB.Prepare("CREATE TABLE IF NOT EXISTS coyoUserrole(userid int NOT NULL, roleid int NOT NULL, PRIMARY KEY (userid, roleid), FOREIGN KEY (userid) REFERENCES coyoUser (id), FOREIGN KEY (roleid) REFERENCES coyoRole (id));")
	if err != nil {
		log.Println(err.Error())
	}
	stmtExec(stmt)
}

func createUserGroupTable() {
	stmt, err := DB.Prepare("CREATE TABLE IF NOT EXISTS coyoUsergroup(userid int NOT NULL, groupid int NOT NULL, PRIMARY KEY (userid, groupid), FOREIGN KEY (userid) REFERENCES coyoUser (id), FOREIGN KEY (groupid) REFERENCES coyoGroup (id));")
	if err != nil {
		log.Println(err.Error())
	}
	stmtExec(stmt)
}

func stmtExec(stmt *sql.Stmt) {
	_, err := stmt.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}

func CloseConnection() {
	err := DB.Close()
	if err != nil {
		log.Println(err.Error())
	}
}

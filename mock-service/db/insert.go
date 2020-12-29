package db

import (
	"log"
	"strconv"
)

//Due to a mock Service there will be NO SQL-Escaping

func InsertUser(email string, lName string, first string, last string, active bool, superadmin bool, password string) {
	/*
		email varchar(30),
		loginName varchar(30),
		firstname varchar(30),
		lastname varchar(30),
		active boolean,
		superadmin boolean,
		password varchar(30)
	*/
	query := "INSERT INTO coyoUser VALUES (NULL, '" + email + "' , '" + lName + "' , '" + first + "' , '" + last + "' ," + strconv.FormatBool(active) + " ," + strconv.FormatBool(superadmin) + " , '" + password + "')"
	insert, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer insert.Close()
}

func InsertRole(roleId string, name string) {
	/*
		roleId varchar(100),
		name varchar(30)
	*/
	query := "INSERT INTO coyoRole VALUES (NULL, '" + roleId + "' , '" + name + "')"
	insert, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer insert.Close()
}

func InsertUserRole(userId int, roleId int) {
	/*
		userId int id,
		roleId int id
	*/
	insertIntoUserGroupOrUserRole("coyoUserrole", userId, roleId)
}

func InsertGroup(groupId string, name string) {
	/*
		groupId varchar(100),
		name varchar(30)
	*/
	query := "INSERT INTO coyoGroup VALUES (NULL, '" + groupId + "' , '" + name + "')"
	insert, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer insert.Close()
}

func InsertUserGroup(userId int, groupId int) {
	/*
		userId int id,
		roleId int id
	*/
	insertIntoUserGroupOrUserRole("coyoUsergroup", userId, groupId)
}

func insertIntoUserGroupOrUserRole(tableName string, userId int, groupRoleId int) {
	query := "INSERT INTO " + tableName + " VALUES ('" + strconv.Itoa(userId) + "' , '" + strconv.Itoa(groupRoleId) + "')"
	insert, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer insert.Close()
}

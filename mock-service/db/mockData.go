package db

import (
	"database/sql"
	"github.com/Pallinder/go-randomdata"
)

var users []string
var roleNames []string
var groupNames []string

func FillDBWithMockData(db *sql.DB) {
	DB = db
	fillUser()
	fillRole()
	fillUserRole()
	fillGroup()
	fillUserGroup()
}

func fillUser() {
	for counter := 0; counter != 50; counter++ {
		profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)
		user := profile.Login.Username
		InsertUser(profile.Email, user, profile.Name.First, profile.Name.Last, randomdata.Boolean(), randomdata.Boolean(), profile.Login.Md5)
		users = append(users, user)
	}
}

func fillRole() {
	for counter := 0; counter != 5; counter++ {
		name := randomdata.SillyName()
		roleNames = append(roleNames, name)
		InsertRole(randomdata.StringSample("mycb32-2ca13e7-aea1e2"+name, "681-13e7aea1-4351afs2"+name, "mst-9afa-b7b1ring3"+name), name)
	}
}

func fillUserRole() {
	for counter := 0; counter != 5; counter++ {
		userId := SelectUserByUsername(users[counter])
		roleId := SelectRoleByName(roleNames[counter])
		InsertUserRole(userId, roleId)
	}
}

func fillGroup() {
	for counter := 0; counter != 5; counter++ {
		name := randomdata.SillyName()
		groupNames = append(groupNames, name)
		InsertGroup(randomdata.StringSample("mycb32-2ca13e7-aea1e2"+name, "681-13e7aea1-4351afs2"+name, "mst-9afa-b7b1ring3"+name), name)
	}
}

func fillUserGroup() {
	for counter := 0; counter != 5; counter++ {
		userId := SelectUserByUsername(users[counter])
		groupId := SelectGroupByName(groupNames[counter])
		InsertUserRole(userId, groupId)
	}
}

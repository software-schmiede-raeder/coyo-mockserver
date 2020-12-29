package db

import (
	"../structs"
	"context"
	"log"
	"strconv"
	"strings"
)

func SelectAllUsers(maxElements int, whereClause string) []structs.Content {
	query := "SELECT * FROM coyoUser "
	if whereClause != "" {
		query = query + whereClause + " "
	}
	if maxElements != -1 {
		query = query + "LIMIT " + strconv.Itoa(maxElements)
	}
	rows, err := DB.QueryContext(context.Background(), query)
	if err != nil {
		log.Println(err)
	}
	tenant := "Tenant"

	var users []structs.Content

	for rows.Next() {
		var id string
		var email string
		var loginName string
		var firstName string
		var lastName string
		var active bool
		var superadmin bool
		var password string
		err := rows.Scan(&id, &email, &loginName, &firstName, &lastName, &active, &superadmin, &password)

		if err != nil {
			log.Println(err)
		} else {
			status := "ACTIVE"
			if !active {
				status = "IN" + status
			}
			users = append(users, structs.Content{
				Tenant:        tenant,
				LoginName:     loginName,
				LoginNameAlt:  "",
				ModeratorMode: false,
				Status:        status,
				Active:        active,
				Anonymized:    false,
				Firstname:     firstName,
				Lastname:      lastName,
				Email:         email,
				Language:      "DE",
				Timezone:      "UTC",
				Properties:    structs.Properties{Unknown: ""},
				Manager:       "Stephan",
				UpdatedID:     id,
				ID:            id,
				EntityID: structs.EntityID{
					ID:       id,
					TypeName: "user",
				},
				Slug:                strings.ToLower(loginName),
				TypeName:            "user",
				DisplayName:         firstName + " " + lastName,
				Color:               "#BEEFFF",
				DisplayNameInitials: string([]rune(firstName)[0]) + string([]rune(lastName)[0]),
				Target: structs.Target{
					Name: "user",
					Params: structs.Params{
						ID:   id,
						Slug: strings.ToLower(loginName),
					},
				},
				ExternalWorkspaceMember: false,
				ImageUrls: structs.ImageUrls{
					Cover:  "https://software-schmiede-raeder.de/",
					Avatar: "https://software-schmiede-raeder.de/img/crown.svg",
				},
				Public: true,
			})
		}
	}
	defer rows.Close()
	return users
}

func SelectUserByUsername(username string) int {
	query := "SELECT cu.id FROM coyoUser cu WHERE cu.loginName LIKE '" + username + "'"
	rows, err := DB.QueryContext(context.Background(), query)
	if err != nil {
		log.Println(err)
	}
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
	defer rows.Close()
	return id
}

func SelectGroupByName(name string) int {
	return selectGroupOrRoleByName("coyoGroup", name)
}

func SelectRoleByName(name string) int {
	return selectGroupOrRoleByName("coyoRole", name)
}

func selectGroupOrRoleByName(tableName string, name string) int {
	query := "SELECT gr.id FROM " + tableName + " gr WHERE gr.name LIKE '" + name + "'"
	rows, err := DB.QueryContext(context.Background(), query)
	if err != nil {
		log.Println(err)
	}
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
	defer rows.Close()
	return id
}

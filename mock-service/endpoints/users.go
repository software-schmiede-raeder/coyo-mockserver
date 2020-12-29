package endpoints

import (
	"../db"
	"../structs"
	"net/http"
	"strconv"
	"strings"
)

func PostPushappinstallations(name string, appType string, token string, id string) (int, string) {
	if name == "fail" || name == "" || (appType != "COYO" && appType != "MESSENGER") || token == "" {
		return http.StatusNotFound, "Mobile app installation could not be created."
	}
	return http.StatusCreated, "Created"
}

func PostCreateUser(createUser structs.CreateUsersRequest) structs.CreateUsersResponse {
	db.InsertUser(createUser.Email, createUser.LoginName, createUser.Firstname, createUser.Lastname, createUser.Active, createUser.Superadmin, createUser.Password)
	return structs.CreateUsersResponse{
		Tenant:        nil,
		LoginName:     createUser.LoginName,
		LoginNameAlt:  nil,
		ModeratorMode: false,
		Status:        "INACTIVE",
		Active:        false,
		Anonymized:    false,
		Firstname:     createUser.Firstname,
		Lastname:      createUser.Lastname,
		Email:         createUser.Email,
		Language:      nil,
		Timezone:      nil,
		Properties:    nil,
		Manager:       nil,
		UpdatedID:     nil,
		ID:            nil,
		EntityID: structs.EntityID{
			ID:       nil,
			TypeName: "user",
		},
		Slug:                nil,
		TypeName:            "user",
		DisplayName:         createUser.Firstname + " " + createUser.Lastname,
		DisplayNameInitials: string([]rune(createUser.Firstname)[0]) + string([]rune(createUser.Lastname)[0]),
		Color:               "#BEEEEF",
		Target: structs.Target{
			Name: "user",
			Params: structs.Params{
				ID:   nil,
				Slug: nil,
			},
		},
		ExternalWorkspaceMember: false,
		ImageUrls: structs.ImageUrls{
			Cover:  "https://software-schmiede-raeder.de/",
			Avatar: "https://software-schmiede-raeder.de/img/crown.svg",
		},
		Public: false,
	}
}

func GetSearchAllUsers(maxElements string, orderBy string, displayName string, status string) structs.SearchAllResponse {
	maxElementsPuffer := setMaxElements(maxElements)

	whereClause := ""
	if displayName != "" {
		whereClause = "WHERE firstName LIKE '" + strings.Split(displayName, " ")[0] + "' "
		whereClause = whereClause + "AND lastName LIKE '" + strings.Split(displayName, " ")[1] + "'"
	}

	if status != "" {
		statusActive := "1"
		if status == "INACTIVE" {
			statusActive = "0"
		}
		if whereClause != "" {
			whereClause = whereClause + " AND "
		} else {
			whereClause = whereClause + "WHERE "
		}
		whereClause = whereClause + "active = " + statusActive
	}

	users := db.SelectAllUsers(maxElementsPuffer, whereClause)
	userSize := len(users)
	empty := false
	if userSize == 0 {
		users = nil
		empty = true
	}
	response := structs.SearchAllResponse{
		Content:          users,
		Pageable:         "INSTANCE",
		TotalPages:       1,
		TotalElements:    userSize,
		Last:             false,
		Number:           0,
		Sort:             structs.Sort{
			Sorted:   false,
			Unsorted: false,
			Empty:    false,
		},
		Size:             userSize,
		NumberOfElements: userSize,
		First:            false,
		Empty:            empty,
	}
	return response
}

func setMaxElements(maxElements string) int {
	if maxElements != "" {
		maxElementsPuffer, _ := strconv.Atoi(maxElements)
		return maxElementsPuffer
	}
	return 20 //Default
}

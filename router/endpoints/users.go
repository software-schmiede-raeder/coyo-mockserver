package endpoints

import (
	user "../../mock-service/endpoints"
	userStructs "../../mock-service/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddRoutes(ginRest *gin.Engine) {

	userRoutesGroup := ginRest.Group("/api")

	//Add a new mobile app installation for a user
	userRoutesGroup.POST("/users/:userId/pushappinstallations", func(c *gin.Context) {
		userId := c.Param("userId")

		//Required POST-Form
		name := c.PostForm("name")
		appType := c.PostForm("appType")
		token := c.PostForm("token")
		//Optional POST-Form
		id := c.PostForm("id")

		code, resultString := user.PostPushappinstallations(name, appType, token, id)
		c.JSON(code, resultString+" "+userId)
	})

	//Creates a user
	userRoutesGroup.POST("/users", func(c *gin.Context) {

		active, _ := strconv.ParseBool(c.PostForm("active"))
		superadmin, _ := strconv.ParseBool(c.PostForm("superadmin"))

		c.JSON(http.StatusOK, user.PostCreateUser(userStructs.CreateUsersRequest{
			Email:                c.PostForm("email"),
			LoginName:            c.PostForm("loginName"),
			LoginNameAlt:         c.PostForm("loginNameAlt"),
			Firstname:            c.PostForm("firstname"),
			Lastname:             c.PostForm("lastname"),
			Active:               active,
			Superadmin:           superadmin,
			GroupIds:             nil,
			RoleIds:              nil,
			RemoteLogonName:      nil,
			PersistedDisplayName: nil,
			Password:             c.PostForm("password"),
			WelcomeMail:          false,
			GeneratePassword:     false,
			Language:             nil,
			TemporaryPassword:    false,
			InitialUser:          false,
		}))
	})

	//Get users
	userRoutesGroup.GET("/users", func(c *gin.Context) {

		//Temporary ignored
		_ = c.Query("_page")

		pageSize := c.Query("_pageSize")

		orderBy := c.Query("_orderBy")
		displayName := c.Query("displayName")

		//Search by status. Either of ACTIVE, INACTIVE or DELETED.
		status := c.Query("status")
		/*with := c.Param("with")
		blocked := c.Param("blocked")
		userIds := c.Param("userIds")
		minUpdatedId := c.Param("minUpdatedId")
		maxUpdatedId := c.Param("maxUpdatedId")
		includeInactiveUsers := c.Param("includeInactiveUsers")*/

		c.JSON(http.StatusOK, user.GetSearchAllUsers(pageSize, orderBy, displayName, status))
	})

	//Deactivate a user
	userRoutesGroup.PUT("/users/:userId/deactivate", func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "test" {
			c.JSON(http.StatusNoContent, "No Content")
		} else {
			c.JSON(http.StatusOK, "OK")
		}
	})

	//Delete a user
	userRoutesGroup.DELETE("/users/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "test" {
			c.JSON(http.StatusNoContent, "No Content")
		} else {
			c.JSON(http.StatusOK, "OK")
		}
	})

	//Deletes a specified mobile app installation of a user
	userRoutesGroup.DELETE("/users/:userId/pushappinstallations/:id", func(c *gin.Context) {
		userId := c.Param("userId")
		id := c.Param("id")
		if userId == "test" {
			c.JSON(http.StatusNoContent, "No Content")
		} else if id == "test" {
			c.JSON(http.StatusNotFound, "Mobile app installation could not be deleted")
		} else {
			c.JSON(http.StatusOK, "OK")
		}
	})

	//Activates a User
	userRoutesGroup.PUT("/users/:userId/activate", func(c *gin.Context) {
		id := c.Param("userId")
		if id == "test" {
			c.JSON(http.StatusNoContent, "No Content")
		} else {
			c.JSON(http.StatusOK, "OK")
		}
	})

	//Disables the moderator mode for a user who is in moderator mode
	userRoutesGroup.DELETE("/users/:userId/moderatorMode", func(c *gin.Context) {
		id := c.Param("userId")
		if id == "test" {
			c.JSON(http.StatusNoContent, "No Content")
		} else {
			c.JSON(http.StatusOK, "OK")
		}
	})

	//Downloads a zip with pre-fetched json data
	userRoutesGroup.GET("/users/download/zip", func(c *gin.Context) {
		//TODO
		//c.JSON(httpCode, resultString)
	})

	//Enables the moderator mode for a user
	userRoutesGroup.PUT("/users/:userId/moderatorMode", func(c *gin.Context) {
		id := c.Param("userId")
		if id == "test" {
			c.JSON(http.StatusNoContent, "No Content")
		} else {
			c.JSON(http.StatusOK, "OK")
		}
	})
	/*
		//Exports all users with administration information. Maximum page size is 1000 (default is 200).
		userRoutesGroup.GET("/users/admin/export", func(c *gin.Context) {
			//Required POST-Form
			page := c.PostForm("_page")
			pageSize := c.PostForm("_pageSize")
			token := c.PostForm("_orderBy")

			c.JSON(httpCode, resultString)
		})

		//Get the user currently logged in
		userRoutesGroup.GET("/users/me", func(c *gin.Context) {
			c.JSON(httpCode, resultString)
		})

		//Returns the notification settings for the given user
		userRoutesGroup.GET("/users/:userId/notification-settings", func(c *gin.Context) {
			userId := c.Param("userId")
			channel := c.PostForm("channel")
			c.JSON(httpCode, resultString)
		})

		//Get a user
		userRoutesGroup.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(httpCode, resultString)
		})

		userRoutesGroup.GET("/users/online/count", func(c *gin.Context) {
			c.JSON(httpCode, resultString)
		})

		userRoutesGroup.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(httpCode, resultString)
		})

	*/
}

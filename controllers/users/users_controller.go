package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jmlgh/bookstore_users-api/domain/users"
	"github.com/jmlgh/bookstore_users-api/services"
	"github.com/jmlgh/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle the error
		return
	}
	if err = json.Unmarshal(bytes, &user); err != nil {
		// TODO: Handle the json error
		fmt.Println(err.Error())
		return
	}*/
	// same as the above code but using Gin
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

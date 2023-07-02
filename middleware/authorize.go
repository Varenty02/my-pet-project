package middleware

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorize(ctx appctx.AppContext, roles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(commons.CurrentUser).(commons.Requester)
		uRole := u.GetUserRole()
		authorize := false
		for _, role := range roles {
			if uRole == role {
				authorize = true
				break
			}
		}
		if authorize {
			c.Set(commons.CurrentUser, u)
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
		}
	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	prefix := "Bearer"

	if token == "" || !strings.HasPrefix(token, prefix) {
		c.JSON(http.StatusUnauthorized, helper.ApiReturn(common.CodeError, "token不存在", nil))
		c.Abort()
		return
	}

	token = token[len(prefix)+1:]

	UserId, err := helper.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeExpires, "权限不足", nil))
		c.Abort()
		return
	}
	_, err = model.GetUserById(UserId)
	if err != nil {
		log.Errorf("Invalid User_id %+v", errors.WithStack(err))
		c.Abort()
		return
	}
	c.Set("user_id", UserId)
	c.Next()
}

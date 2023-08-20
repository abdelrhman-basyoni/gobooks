package middlewares

import (
	"net/http"
	"strings"

	"github.com/abdelrhman-basyoni/gobooks/dto"
	customErrors "github.com/abdelrhman-basyoni/gobooks/errors"
	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			errorMassage := err.Error()
			switch err.Err.(type) {
			case *customErrors.DataBaseError:
				if strings.Contains(errorMassage, "E11000 duplicate key") {
					errorMassage = "Duplicate key"

				}
				errorMassage = "DataBaseError"
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response[any]{ErrorMessage: &errorMassage})
		}

	}
}

package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/abdelrhman-basyoni/gobooks/dto"
	customErrors "github.com/abdelrhman-basyoni/gobooks/errors"
	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var errorsList []string
		for _, err := range c.Errors {
			var errorMassage string
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)

			switch err.Err.(type) {
			case *customErrors.DataBaseError:
				if strings.Contains(errorMassage, "E11000 duplicate key") {

					errorMassage = "Duplicate key"

				} else {
					errorMassage = "DataBaseError"
				}
				break
			default:
				errorMassage = err.Error()

			}
			errorsList = append(errorsList, errorMassage)

		}
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response[any]{ErrorMessage: &errorsList})
	}
}

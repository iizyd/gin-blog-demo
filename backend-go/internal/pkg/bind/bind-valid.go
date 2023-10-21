package bind

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, param any) error {
	if err := c.ShouldBind(param); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(param); err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `json:"name" form:"name" binding:"required,myvalidator"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Age   int    `json:"age" form:"age" binding:"required,gte=1,lte=100"`
}

func MyValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	fmt.Println(value)
	return len(value) >= 5 && len(value) < 10
}

func Sv(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)
	if user.Name == "admin" && user.Age < 18 {
		sl.ReportError(user.Age, "age", "Age", "age", "admin must be older than 18")
	}
}
func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("myvalidator", MyValidator)
		v.RegisterStructValidation(Sv, User{})
	}

	r := gin.Default()

	r.POST("/abc", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"msg": err,
			})
			fmt.Println(err.Error())
			return
		}

		c.JSON(200, gin.H{
			"msg": "successful",
		})
	})
	r.Run(":8080")
}

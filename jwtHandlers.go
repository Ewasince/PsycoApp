package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*User); ok {
			return jwt.MapClaims{
				identityKey: v.id,
				usernameKey: v.username,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)

		fmt.Printf("identityHandler user_id0=%v\n", claims[identityKey])
		user_id := claims[identityKey].(float64)
		user := usersByIds[user_id]
		//if err {
		//	panic("Cannot find user")
		//}
		fmt.Printf("identityHandler user_id=%v\n", user_id)
		fmt.Printf("identityHandler user=%v\n", user)
		return &User{
			id:        user.id,
			username:  user.username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		username := loginVals.Username
		password := loginVals.Password

		fmt.Printf("usersCreds=%v\n", usersCreds)
		fmt.Printf("usersByIds=%v\n", usersByIds)

		user, ok := usersCreds[username]

		if !ok {
			return nil, jwt.ErrFailedAuthentication
		}

		if user.password != password {
			return nil, jwt.ErrFailedAuthentication
		}

		return &user, nil
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		//fmt.Printf("authorizator %s\n", data.(*User))
		//if v, ok := data.(*User); ok && v.UserName == "admin" {
		//	return true
		//}
		//return false.
		return true
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func handleNoRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
}

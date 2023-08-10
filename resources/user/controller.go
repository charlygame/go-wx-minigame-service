package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Get(c *gin.Context) {
	userID := c.Param("user_id")
	s := NewUserService()
	result, err := s.Get(userID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Insert(c *gin.Context) {
	var body User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := NewUserService()
	id, err := s.Insert(body)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func Update(c *gin.Context) {
	userID := c.Param("user_id")
	var body User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := NewUserService()
	result, err := s.Update(userID, body)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func WXLogin(c *gin.Context) {
	code := c.Param("code")
	s := NewUserService()
	result, err := s.WXLogin(code)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	// 查找库里是否存在OpenID
	count, err := s.Count(bson.M{"wx_open_id": result.OpenID})
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	if count <= 0 {
		userId, err := s.WXRegister(&User{
			Score:        0,
			Username:     "",
			WxOpenId:     result.OpenID,
			WxSessionKey: result.SessionKey,
		})
		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": userId})

		return
	} else {
		// 根据OpenID查找用户
		user, err := s.FindOne(bson.M{"wx_open_id": result.OpenID})
		// 更新SessionKey
		s.Update(user.ID, User{
			WxSessionKey: result.SessionKey,
		})
		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": user.ID})
	}
}

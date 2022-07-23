package controllers

import "github.com/gin-gonic/gin"

type Account interface {
	Get(c *gin.Context)
}

type Auth interface {
	RefreshToken(c *gin.Context)
}

type Guild interface {
	Get(c *gin.Context)
	GetUserOne(c *gin.Context)
	GetUserAll(c *gin.Context)
}

type GuildConfig interface {
	Overwrite(c *gin.Context)
	Get(c *gin.Context)
	GetPreset(c *gin.Context)
}

type Oauth2 interface {
	NewURL(c *gin.Context)
	NewInviteBotURL(c *gin.Context)
	DiscordCallback(c *gin.Context)
}

type Controllers struct {
	Account
	Auth
	Guild
	GuildConfig
	Oauth2
}

func errorResponse(err error) gin.H {
	return gin.H{"message": err.Error()}
}

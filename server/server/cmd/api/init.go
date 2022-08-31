package api

import (
	"server/cmd/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// The database can be accessed within the Handler
// because the DBController is owned within the structure.
type APIController struct {
	DBc *database.DBController
}

func NewAPIController(DBc *database.DBController) *APIController {
	return &APIController{
		DBc: DBc,
	}
}

func Setup(APIc APIController) *gin.Engine {
	router := gin.Default()
	corsConfig(router)

	router.POST("/user/create", APIc.UserHandler)
	router.PUT("/user/update", APIc.UserHandler)
	router.GET("/user/get", APIc.UserHandler)

	router.POST("/groupy/create", APIc.GroupHandler)
	router.PUT("/groupy/update", APIc.GroupHandler)
	router.DELETE("/groupy/delete", APIc.GroupHandler)

	router.POST("/uig/create", APIc.UserInGroupyHandler)
	router.DELETE("/uig/delete", APIc.UserInGroupyHandler)

	router.POST("/boarditem/create", APIc.BoardItemHandler)
	router.PUT("/boarditem/update", APIc.BoardItemHandler)
	router.DELETE("/boarditem/delete", APIc.BoardItemHandler)

	router.POST("/gbi/create", APIc.GroupyBoardItemHandler)
	router.DELETE("/gbi/delete", APIc.GroupyBoardItemHandler)

	router.POST("/ulbi/create", APIc.UserLikeBoardItemHandler)
	router.DELETE("/ulbi/delete", APIc.UserLikeBoardItemHandler)

	return router
}

func corsConfig(router *gin.Engine) *gin.Engine {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
			"PUT",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
	}))

	return router
}

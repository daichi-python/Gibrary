package api

import (
	"fmt"
	"net/http"
	"server/cmd/converter"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (APIc *APIController) UserHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var user converter.User
		if err := c.ShouldBindJSON(&user); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of User struct and json request failed.")
			return
		}

		params := converter.ToInsertQueryParams(user)
		result, err := APIc.DBc.InsertUser(params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "InsertUser",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		id, _ := result.LastInsertId()
		user.ID = fmt.Sprint(id)

		c.IndentedJSON(http.StatusOK, user)
		logrus.WithField("ID", id).Debugln("UserHandler has completed successfully.")

	case http.MethodPut:
		var user converter.User
		if err := c.ShouldBindJSON(&user); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of user struct and json request failed.")
			return
		}

		id, params := converter.ToUpdateQueryParams(user)
		_, err := APIc.DBc.UpdateUser(id, params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "UpdateUser",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, user)
		logrus.WithField("ID", id).Debugln("UserHandler has completed successfully.")

	case http.MethodGet:
		var user converter.User
		if err := c.ShouldBindJSON(&user); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of user struct and json request failed.")
			return
		}

		params := converter.ToSelectQueryParams(user)
		row := APIc.DBc.SelectUser(params)
		res_user, err := converter.UserRowToStruct(row)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Error":      err,
				"Function":   "UserRowToStruct",
				"RequestURL": c.Request.URL,
			}).Errorln("An error occurred during the execution of UserRowToStruct")
		}

		c.IndentedJSON(http.StatusOK, res_user)
		logrus.WithFields(logrus.Fields{
			"ID":     res_user.ID,
			"Method": http.MethodGet,
		}).Debugln("UserHandler has completed successfully.")

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Bad Request",
		})

		logrus.WithField("Method", c.Request.Method).Errorln(
			"This Handler accept POST or PUT Method.",
		)
		return
	}
}

func (APIc *APIController) GroupHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var groupy converter.Groupy
		if err := c.ShouldBindJSON(&groupy); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of groupy struct and json request failed.")
			return
		}

		params := converter.ToInsertQueryParams(groupy)
		result, err := APIc.DBc.InsertGroupy(params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "InsertGroupy",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		id, _ := result.LastInsertId()
		groupy.ID = fmt.Sprint(id)
		c.IndentedJSON(http.StatusOK, groupy)
		logrus.WithField("ID", id).Debugln("GroupyHandler has completed successfully.")

	case http.MethodPut:
		var groupy converter.Groupy
		if err := c.ShouldBindJSON(&groupy); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of groupy struct and json request failed.")
			return
		}

		id, params := converter.ToUpdateQueryParams(groupy)
		_, err := APIc.DBc.UpdateGroupy(id, params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "UpdateGroupy",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, groupy)
		logrus.WithField("ID", id).Debugln("GroupyHandler has completed successfully.")

	case http.MethodDelete:
		var groupy converter.Groupy
		if err := c.ShouldBindJSON(&groupy); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of groupy struct and json request failed.")
			return
		}

		_, err := APIc.DBc.DeleteGroupy(groupy.ID)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "DeleteteGroupy",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, groupy)
		logrus.WithField("ID", groupy.ID).Debugln("GroupyHandler has completed successfully.")

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Bad Request",
		})

		logrus.WithField("Method", c.Request.Method).Errorln(
			"This Handler accept POST or PUT or DELETE Method.",
		)
		return
	}
}

func (APIc *APIController) UserInGroupyHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var uig converter.UserInGroupy
		if err := c.ShouldBindJSON(&uig); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of UserInGroupy struct and json request failed.")
			return
		}

		params := converter.ToInsertQueryParams(uig)
		result, err := APIc.DBc.InsertUserInGroupy(params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "InsertUserInGroupy",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		id, _ := result.LastInsertId()
		uig.ID = fmt.Sprint(id)
		c.IndentedJSON(http.StatusOK, uig)

		logrus.WithField("ID", uig.ID).Debugln("UserInGroupyHandler has completed successfully.")

	case http.MethodDelete:
		var uig converter.UserInGroupy
		if err := c.ShouldBindJSON(&uig); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of UserInGroupy struct and json request failed.")
			return
		}

		_, err := APIc.DBc.DeleteUserInGroupy(uig.ID)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "DeleteUserInGroupy",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, uig)
		logrus.WithField("ID", uig.ID).Debugln("GroupyHandler has completed successfully.")

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Bad Request",
		})

		logrus.WithField("Method", c.Request.Method).Errorln(
			"This Handler accept POST or DELETE Method.",
		)
		return
	}
}

func (APIc *APIController) BoardItemHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var board_item converter.BoardItem
		if err := c.ShouldBindJSON(&board_item); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of BoardItem struct and json request failed.")
			return
		}

		params := converter.ToInsertQueryParams(board_item)
		result, err := APIc.DBc.InsertBoardItem(params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "InsertBoardItem",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		id, _ := result.LastInsertId()
		board_item.ID = fmt.Sprint(id)
		c.IndentedJSON(http.StatusOK, board_item)
		logrus.WithField("ID", board_item.ID).Debugln("BoardItemHandler has completed successfully.")

	case http.MethodPut:
		var board_item converter.BoardItem
		if err := c.ShouldBindJSON(&board_item); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of BoardItem struct and json request failed.")
			return
		}

		id, params := converter.ToUpdateQueryParams(board_item)
		_, err := APIc.DBc.UpdateBoardItem(id, params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "UpdateBoardItem",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, board_item)
		logrus.WithField("ID", board_item.ID).Debugln("BoardItemHandler has completed successfully.")

	case http.MethodDelete:
		var board_item converter.BoardItem
		if err := c.ShouldBindJSON(&board_item); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of BoardItem struct and json request failed.")
			return
		}

		_, err := APIc.DBc.DeleteBoardItem(board_item.ID)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "DeleteBoardItem",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, board_item)
		logrus.WithField("ID", board_item.ID).Debugln("BoardItemHandler has completed successfully.")

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Bad Request",
		})

		logrus.WithField("Method", c.Request.Method).Errorln(
			"This Handler accept POST or PUT or DELETE Method.",
		)
		return
	}
}

func (APIc *APIController) GroupyBoardItemHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var groupy_board_item converter.GroupyBoardItem
		if err := c.ShouldBindJSON(&groupy_board_item); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of GroupyBoardItem struct and json request failed.")
			return
		}

		params := converter.ToInsertQueryParams(groupy_board_item)
		result, err := APIc.DBc.InsertGroupyBoardItem(params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "InsertGroupyBoardItem",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		id, _ := result.LastInsertId()
		groupy_board_item.ID = fmt.Sprint(id)
		c.IndentedJSON(http.StatusOK, groupy_board_item)
		logrus.WithField("ID", groupy_board_item.ID).Debugln("GroupyBoardItemHandler has completed successfully.")

	case http.MethodDelete:
		var groupy_board_item converter.GroupyBoardItem
		if err := c.ShouldBindJSON(&groupy_board_item); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of GroupyBoardItem struct and json request failed.")
			return
		}

		_, err := APIc.DBc.DeleteBoardItem(groupy_board_item.ID)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "DeleteBoardItem",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, groupy_board_item)
		logrus.WithField("ID", groupy_board_item.ID).Debugln("GroupyBoardItemHandler has completed successfully.")

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Bad Request",
		})

		logrus.WithField("Method", c.Request.Method).Errorln(
			"This Handler accept POST or DELETE Method.",
		)
		return
	}
}

func (APIc *APIController) UserLikeBoardItemHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		var user_like_board_item converter.UserLikeBoardItem
		if err := c.ShouldBindJSON(&user_like_board_item); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of UserLikeBoardItem struct and json request failed.")
			return
		}

		params := converter.ToInsertQueryParams(user_like_board_item)
		result, err := APIc.DBc.InsertUserLikeBoardItem(params)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "InsertUserLikeBoardItem",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		id, _ := result.LastInsertId()
		user_like_board_item.ID = fmt.Sprint(id)
		c.IndentedJSON(http.StatusOK, user_like_board_item)
		logrus.WithField("ID", user_like_board_item.ID).Debugln("UserLikeBoardItemHandler has completed successfully.")

	case http.MethodDelete:
		var user_like_board_item converter.UserLikeBoardItem
		if err := c.ShouldBindJSON(&user_like_board_item); err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Method":     c.Request.Method,
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Bind of UserLikeBoardItem struct and json request failed.")
			return
		}

		_, err := APIc.DBc.DeleteUserLikeBoardItem(user_like_board_item.ID)
		if err != nil {
			InternalServerError(c)
			logrus.WithFields(logrus.Fields{
				"Function":   "DeleteUserLikeBoardItem",
				"Error":      err,
				"RequestURL": c.Request.URL,
			}).Errorln("Query execution failed.")
			return
		}

		c.IndentedJSON(http.StatusOK, user_like_board_item)
		logrus.WithField("ID", user_like_board_item.ID).Debugln("UserLikeBoardItemHandler has completed successfully.")

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Bad Request",
		})

		logrus.WithField("Method", c.Request.Method).Errorln(
			"This Handler accept POST or DELETE Method.",
		)
		return
	}
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"Message": "Please allow time for access.",
	})
}

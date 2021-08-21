package controller

import (
	"github.com/gin-gonic/gin"
	"web_app/dao/mysql"
	"web_app/model"
)

func Text(c *gin.Context) {
	param := make([]string, 5)
	param[0] = "xg"
	param[1] = "wm"

	inter := model.Interview{Questions: param, Scores: param}

	m := "insert into interviews(user_id,questions,scores) values(?,?,?)"

	db := mysql.DB()
	db.Exec(m, 1, inter.Questions, inter.Scores)

	m2 := "select questions from interviews where user_id=?"
	inter2 := model.Interview{}
	db.Get(&inter2, m2, 1)

	c.JSON(400, &inter2)

}

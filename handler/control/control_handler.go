package handler_control

import (
	"cmdb_backend/control"
	"cmdb_backend/logger"
	"cmdb_backend/property"

	"github.com/gin-gonic/gin"
)

func Send_Cmd(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	s := control.Send_Cmd_Struct{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	//err = p.Create_Bu(db)
	succ, fail, err := s.Send_Cmd(db)
	if err == nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{"code": 200,
			"success": succ,
			"fail":    fail,
		})
	}

	return
}

func Send_File(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	s := control.Send_File_Struct{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	//err = p.Create_Bu(db)
	_, fail, err := s.Send_File(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{"code": 200,
			"fail": fail,
		})
	}

	return
}

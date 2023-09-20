package handler_property

import (
	"cmdb_backend/logger"
	"cmdb_backend/property"

	"github.com/gin-gonic/gin"
)

func Delete_Depart(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.Department{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_Depart_By_Name(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Create_Depart(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.Department{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Create_Department(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

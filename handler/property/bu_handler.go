package handler_property

import (
	"cmdb_backend/logger"
	"cmdb_backend/property"

	"github.com/gin-gonic/gin"
)

func Create_Bu(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.BU{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Create_Bu(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Delete_Bu(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.BU{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_BU_By_Name(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

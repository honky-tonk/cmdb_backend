package handler_property

import (
	"cmdb_backend/logger"
	"cmdb_backend/property"

	"github.com/gin-gonic/gin"
)

func Update_Team_Departbelong(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.Team_Update_Departbelong{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Update_depart_belong(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Update_Team_Name(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.Team_Update_Teamname{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Update_depart_name(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Create_Team(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.Team{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Create_Team(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Delete_Team(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.Team{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_Team_By_Name(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

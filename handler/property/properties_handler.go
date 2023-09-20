package handler_property

import (
	"cmdb_backend/logger"
	"cmdb_backend/property"

	"github.com/gin-gonic/gin"
)

func Create_Properties(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()

	p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Create_Property(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Delete_Properties_Resource(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Identify{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_Property_Identify(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Delete_Properties_Userpri(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Delete_Userpriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_Property_Userpriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Delete_Properties_Teampri(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Delete_Teampriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_Property_Teampriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Delete_Properties_Departpri(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Delete_Departpriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_Property_Departpriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Delete_Properties_Bupri(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Delete_Bupriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.Delete_Property_Bupriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
	}
	return
}

func Update_Properties_Hostname(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Update_Hostname{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.UPDATE_Property_Hostname(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
		logger.Info.Println(err)
	}
	return
}

func Update_Properties_Userpriv(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Update_Userpriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.UPDATE_Property_Userpriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
		logger.Info.Println(err)
	}
	return
}

func Update_Properties_Bupriv(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Update_Bupriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.UPDATE_Property_Bupartpriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
		logger.Info.Println(err)
	}
	return
}

func Update_Properties_Teampriv(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Update_Teampriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.UPDATE_Property_Teampriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
		logger.Info.Println(err)
	}
	return
}

func Update_Properties_Departpriv(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Update_Departpriv{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	err = p.UPDATE_Property_Departpriv(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
		logger.Info.Println(err)
	}
	return
}

func Show_Properties_Userpri(c *gin.Context) {
	db := property.Db_main()
	defer db.Close()
	p := property.Properties_Identify{}
	//p := property.Properties{}
	//fmt.Println("request body is ", c.Request.Body)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": "please input correct format"})
		logger.Info.Println(err)
		return
	}
	//fmt.Println("p is ", p)
	res, err := p.SHOW_Property_Userpri(db)

	if err != nil {
		c.JSON(400, gin.H{"code": 400,
			"message": err.Error()})
		logger.Info.Println(err)
	}

	logger.Info.Println("res is ", res)

	return
}

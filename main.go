package main

import (
	"cmdb_backend/global"
	handler_control "cmdb_backend/handler/control"
	handler_property "cmdb_backend/handler/property"

	"github.com/gin-gonic/gin"
)

//var Env_file map[string]string

func main() {
	global.Init()

	//创建route
	router := gin.New() //创建一个router engine
	Api_route := router.Group("/api")

	//properties api
	properties_api := Api_route.Group("/properties")
	properties_api.POST("/create", handler_property.Create_Properties)

	properties_api.POST("/delete", handler_property.Delete_Properties_Resource)
	properties_api.POST("/delete/userprivilege", handler_property.Delete_Properties_Userpri)
	properties_api.POST("/delete/teamprivilege", handler_property.Delete_Properties_Teampri)
	properties_api.POST("/delete/buprivilege", handler_property.Delete_Properties_Bupri)
	properties_api.POST("/delete/departmentprivilege", handler_property.Delete_Properties_Departpri)

	properties_api.POST("/update/hostname", handler_property.Update_Properties_Hostname)
	properties_api.POST("/update/userprivilege", handler_property.Update_Properties_Userpriv)
	properties_api.POST("/update/teamprivilege", handler_property.Update_Properties_Teampriv)
	properties_api.POST("/update/buprivilege", handler_property.Update_Properties_Bupriv)
	properties_api.POST("/update/departmentprivilege", handler_property.Update_Properties_Departpriv)

	properties_api.POST("/show/test", handler_property.Show_Properties_Userpri)

	//bu api
	bu := Api_route.Group("/bu")
	bu.POST("/create", handler_property.Create_Bu)
	bu.POST("/delete", handler_property.Delete_Bu)

	//department api
	de := Api_route.Group("/department")
	de.POST("/create", handler_property.Create_Depart)
	de.POST("/delete", handler_property.Delete_Depart)

	//team api
	team := Api_route.Group("/team")
	team.POST("create", handler_property.Create_Team)
	team.POST("/delete", handler_property.Delete_Team)

	//control api
	control_api := Api_route.Group("/control")
	control_api.POST("/send_cmd", handler_control.Send_Cmd)
	control_api.POST("/send_file", handler_control.Send_File)

	router.Run("localhost:8080")
}

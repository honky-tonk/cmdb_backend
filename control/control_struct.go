package control

import "cmdb_backend/property"

type Send_Cmd_Struct struct {
	Pi    property.Properties_Identify `json:"property_identify"`
	Hosts []string                     `json:"hosts"`
	Cmd   string                       `json:"cmd"`
}

type Send_File_Struct struct {
	Pi        property.Properties_Identify `json:"property_identify"`
	Hosts     []string                     `json:"hosts"`
	File_Name string                       `json:"file_name"`
	Work_Dir  string                       `json:"work_dir"`
}

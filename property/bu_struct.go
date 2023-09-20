package property

//"database/sql"
//"cmdb_backend/database"

type BU struct {
	BU_Name string `json:"bu_name"` //在序列化放入json存储的时候使用这个key名字存储
	//BU_ID       int    `json:"bu_id"`
	Departments []string `json:"departments"`
}

func CreateBU(name string, de []string) *BU {
	BU_Ins := BU{
		BU_Name: name,
		//BU_ID:       id,
		Departments: de,
	}

	return &BU_Ins
}

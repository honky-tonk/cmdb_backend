package property

type Department struct {
	BU_Belong_to    string `json:"bu_belong_to"`
	Department_Name string `json:"department_name"`
	//Department_ID   int    `json:"department_id"`
	Teams []string `json:"teams"`
}

type Insert_Team_Into_Department struct {
	D Department `json:"department"`
}

func CreateDepartment(bu string, name string, t []string) *Department {
	De_Ins := Department{
		BU_Belong_to:    bu,
		Department_Name: name,
		//Department_ID:   department_id,
		Teams: t,
	}

	return &De_Ins
}

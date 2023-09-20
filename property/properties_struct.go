package property

import "fmt"

//"errors"

const (
	//properties type
	VPS int = iota
	Switch
	Router
)

const (
	//Privilege
	Privilege_For_Users = 100 + iota
	Privilege_For_Departments
	Privilege_For_BUs
)

type Properties struct {
	//belong to
	//if public for some department Team_Belong_to field is -1,same as public for some BU ,if public some BU,Department_Belong_to && Team_Belong_to field is -1
	BU_Belong_to         string `json:"bu_belong_to"`
	Department_Belong_to string `json:"department_belong_to"`
	Team_Belong_to       string `json:"team_belong_to"`

	//Is_Public == 1 mean public for BU's everyone, Is_Public == 2 mean public for Department's everyone, -1 mean not public just belong to some team
	//Is_Public int `json:"is_public"`

	//properties_type
	Properties_type int    `json:"properties_type"`
	Ipv4_addr       string `json:"ipv4_addr"`
	Hostname        string `json:"hostname"`

	//privilege of guest
	Privilege_For_Users       []string `json:"privilege_for_users"`       //user ids
	Privilege_For_Teams       []string `json:"privilege_for_Teams"`       //team ids
	Privilege_For_Departments []string `json:"privilege_for_departments"` //Department ids
	Privilege_For_BUs         []string `json:"privilege_for_bus"`         //BU ids

}

type Properties_Identify struct {
	//在POST请求进来的时候必须对比这三条,只有这3条相等才能删除,换句话说这三条信息相等才能辨别一个资产
	Properties_type int    `json:"properties_type"`
	Ipv4_addr       string `json:"ipv4_addr"`
	Hostname        string `json:"hostname"`
}

type Properties_Delete_Userpriv struct {
	//use to update hostname for post
	Pi              Properties_Identify `json:"properties_identify"`
	Userpriv_Delete []string            `json:"userpriv_delete"`
}

type Properties_Delete_Teampriv struct {
	//use to update hostname for post
	Pi              Properties_Identify `json:"properties_identify"`
	Teampriv_Delete []string            `json:"teampriv_delete"`
}

type Properties_Delete_Departpriv struct {
	//use to update hostname for post
	Pi                Properties_Identify `json:"properties_identify"`
	Departpriv_Delete []string            `json:"departpriv_delete"`
}

type Properties_Delete_Bupriv struct {
	//use to update hostname for post
	Pi            Properties_Identify `json:"properties_identify"`
	Bupriv_Delete []string            `json:"bupriv_delete"`
}

type Properties_Update_Hostname struct {
	//use to update hostname for post
	Pi                     Properties_Identify `json:"properties_identify"`
	Hostname_Before_Update string              `json:"hostname_before_update"`
}

type Properties_Update_Userpriv struct {
	//use to update user privilege for post
	Pi              Properties_Identify `json:"properties_identify"`
	Userpriv_Update []string            `json:"userpriv_update"`
}

type Properties_Update_Teampriv struct {
	//use to update team privilege for post
	Pi              Properties_Identify `json:"properties_identify"`
	Teampriv_Update []string            `json:"teampriv_update"`
}

type Properties_Update_Departpriv struct {
	//use to update department privilege for post
	Pi                Properties_Identify `json:"properties_identify"`
	Departpriv_Update []string            `json:"departpriv_update"`
}

type Properties_Update_Bupriv struct {
	//use to update Bu privilege for post
	Pi            Properties_Identify `json:"properties_identify"`
	Bupriv_Update []string            `json:"Bupriv_update"`
}

type Show_Properties_Userpriv struct {
	//use to update user privilege for post
	Pi       Properties_Identify `json:"properties_identify"`
	Userpriv []string            `json:"userpriv"`
}

type Show_Properties_Teampriv struct {
	//use to update user privilege for post
	Pi       Properties_Identify `json:"properties_identify"`
	Teampriv []string            `json:"teampriv"`
}

type Show_Properties_Departpriv struct {
	//use to update user privilege for post
	Pi         Properties_Identify `json:"properties_identify"`
	Departpriv []string            `json:"departpriv"`
}

type Show_Properties_Bupriv struct {
	//use to update user privilege for post
	Pi     Properties_Identify `json:"properties_identify"`
	Bupriv []string            `json:"bupriv"`
}

func CreateProperties(bu_name string, department_name string, team_name string,
	/*is_public int,*/
	properties_type int, ipaddr string, hostname string,
	/*privilege_for_users []int, priviliges_for_teams []int, priviliges_for_departments []int, priviliges_for_bus []int,*/
) *Properties {

	properties_ins := Properties{
		BU_Belong_to:         bu_name,
		Department_Belong_to: department_name,
		Team_Belong_to:       team_name,

		//Is_Public: -1,

		Properties_type: properties_type,
		Ipv4_addr:       ipaddr,
		Hostname:        hostname,

		Privilege_For_Users:       nil,
		Privilege_For_Teams:       nil,
		Privilege_For_Departments: nil,
		Privilege_For_BUs:         nil,
	}

	return &properties_ins

}

//delete element of s1, which also exist in s2
func delete_pri(s1 []string, s2 []string) []string {
	//var v []int64
	for _, i := range s2 {
		for index_s1, j := range s1 {
			if i == j {
				s1 = append(s1[:index_s1], s1[index_s1+1:]...)
			}
		}
	}
	return s1
}

//check s2 if all exist in s1
func if_all_exist(s1 []string, s2 []string) bool {
	//s2 can not be nil
	if len(s2) == 0 || len(s1) == 0 {
		return false
	}
	fmt.Println("s2 is ", s2)
	if_all_exit := true

	for _, i := range s2 {
		exist := false
		for _, j := range s1 {
			if j == i {
				exist = true
			}
		}
		if exist == false {
			if_all_exit = false
			break
		}
	}

	return if_all_exit
}

// check two slice if element of slice exist in anthoer slice
func if_exist(s1 []string, s2 []string) bool {

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	if len(s1) == 0 {
		return false
	}

	//s2 can not be null
	if len(s2) == 0 {
		return true
	}

	if_exist := false

	for _, i := range s1 {
		for _, j := range s2 {
			if i == j {
				if_exist = true
				break
			}

		}
	}

	return if_exist
}

//unsafe,if value in []int64 big than int will error
/*
func convert_int64_int_slice([]int64)[]int{

}
*/

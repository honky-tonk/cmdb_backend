package property

import (
	"database/sql"
	"errors"
	"fmt"

	"cmdb_backend/logger"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

//only check ip
func Check_Property_Exist_By_IP(db *sql.DB, addr string) (bool, error) {
	var ret bool

	//rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM property WHERE hostname = $1) ;", hostname)
	show := `SELECT EXISTS(SELECT 1 FROM property WHERE  address = $1  ) ;`
	rows, err := db.Query(show, addr)
	if err != nil {
		logger.Warning.Println("can't EXEC sql: ", err)
		return false, err
	}

	for rows.Next() {
		err = rows.Scan(&ret)
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	//fmt.Println("rows is ", rows)
	return ret, nil

}

//hostname type ip_addr都不能出现重复
func Check_Property_Exist(db *sql.DB, ptype int, addr string, hostname string) (bool, error) {
	var ret bool

	//rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM property WHERE hostname = $1) ;", hostname)
	rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM property WHERE properties_type = $1 AND address = $2 AND hostname = $3 ) ;", ptype, addr, hostname)
	if err != nil {
		logger.Warning.Println("can't EXEC sql: ", err)
		return false, err
	}

	for rows.Next() {
		err = rows.Scan(&ret)
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	//fmt.Println("rows is ", rows)
	return ret, nil

}

func (prop *Properties) Create_Property(db *sql.DB) error {
	ret, err := Check_Property_Exist(db, prop.Properties_type, prop.Ipv4_addr, prop.Hostname)
	//fmt.Println("--------------*****-------------")
	if err != nil {
		return err
	}
	if ret == true {
		//exist
		return errors.New("Insert Property exist")
	}
	//fmt.Println("---------------------------")
	//check bu is exist
	exist := Check_Bu_Exist(db, prop.BU_Belong_to)
	if exist != true {
		//bu not exist
		//logger.Info.Println("bu not exist when ");
		return errors.New("error: BU not exist ")
	}

	//check department is exist
	exist = Check_Depart_Exist(db, prop.Department_Belong_to)
	if exist != true {
		//depart not exist
		return errors.New("error: Department not exist ")
	}

	//check team is exist
	exist = Check_Team_Exist(db, prop.Team_Belong_to)
	if exist != true {
		//team not exist
		return errors.New("error: Team not exist ")
	}

	query := `INSERT INTO property (properties_type, address, hostname, bu_belong_to, department_belong_to, team_belong_to, privilege_for_users, privilege_for_teams, privilege_for_departments, privilege_for_bus) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	_, err = db.Query(query, prop.Properties_type, prop.Ipv4_addr, prop.Hostname, prop.BU_Belong_to, prop.Department_Belong_to, prop.Team_Belong_to, pq.Array(prop.Privilege_For_Users), pq.Array(prop.Privilege_For_Teams), pq.Array(prop.Privilege_For_Departments), pq.Array(prop.Privilege_For_BUs))
	if err != nil {
		logger.Error.Fatal("can't prepare database: ", err)
	}
	return nil

}

func (pd *Properties_Identify) Delete_Property_Identify(db *sql.DB) error {
	ret, err := Check_Property_Exist(db, pd.Properties_type, pd.Ipv4_addr, pd.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("Insert Property Non exist")
	}

	delete := `DELETE FROM property WHERE properties_type = $1 AND address = $2 AND hostname = $3 ;`
	_, err = db.Query(delete, pd.Properties_type, pd.Ipv4_addr, pd.Hostname)
	if err != nil {
		logger.Error.Fatal("can't delete database: ", err)
	}
	return nil
}

func (p *Properties_Delete_Userpriv) Delete_Property_Userpriv(db *sql.DB) error {
	// check if exist
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("delete Property Non exist")
	}

	//check if privilege ALL repeatly
	pri, err := p.Pi.SHOW_Property_Userpri(db)
	if err != nil {
		return err
	}

	exist := if_all_exist(pri, p.Userpriv_Delete)
	//fmt.Println("exist, pri, p.userpri_update", exist, pri, p.Userpriv_Delete)
	if exist != true {
		return errors.New("privilege not exist!!!")
	}

	new_pri := delete_pri(pri, p.Userpriv_Delete)
	//Update sql
	update := ` UPDATE property SET privilege_for_users = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Delete_Teampriv) Delete_Property_Teampriv(db *sql.DB) error {
	// check if exist
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("delete Property Non exist")
	}

	//check if privilege ALL repeatly
	pri, err := p.Pi.SHOW_Property_Teampri(db)
	if err != nil {
		return err
	}
	//fmt.Println("pri, p.Teampri_update", pri, p.Teampriv_Delete)
	exist := if_all_exist(pri, p.Teampriv_Delete)
	if exist != true {
		return errors.New("privilege not exist!!!")
	}

	new_pri := delete_pri(pri, p.Teampriv_Delete)
	//Update sql
	update := ` UPDATE property SET privilege_for_teams = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Delete_Departpriv) Delete_Property_Departpriv(db *sql.DB) error {
	// check if exist
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("delete Property Non exist")
	}

	//check if privilege ALL repeatly
	pri, err := p.Pi.SHOW_Property_Departpri(db)
	if err != nil {
		return err
	}
	//fmt.Println("pri, p.departpri_update", pri, p.Departpriv_Delete)
	exist := if_all_exist(pri, p.Departpriv_Delete)
	if exist != true {
		return errors.New("privilege not exist!!!")
	}

	new_pri := delete_pri(pri, p.Departpriv_Delete)
	//Update sql
	update := ` UPDATE property SET privilege_for_departments = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Delete_Bupriv) Delete_Property_Bupriv(db *sql.DB) error {
	// check if exist
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("delete Property Non exist")
	}

	//check if privilege ALL repeatly
	pri, err := p.Pi.SHOW_Property_Bupri(db)
	if err != nil {
		return err
	}
	//fmt.Println("pri, p.bupri_update", pri, p.Bupriv_Delete)
	exist := if_all_exist(pri, p.Bupriv_Delete)
	if exist != true {
		return errors.New("privilege not exist!!!")
	}

	new_pri := delete_pri(pri, p.Bupriv_Delete)
	//Update sql
	update := ` UPDATE property SET privilege_for_bus = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Update_Hostname) UPDATE_Property_Hostname(db *sql.DB) error {
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	fmt.Println("return is ", ret)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("Insert Property Non exist")
	}

	update := `UPDATE property SET hostname = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, p.Hostname_Before_Update, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Update_Userpriv) UPDATE_Property_Userpriv(db *sql.DB) error {
	// check if exist
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("Insert Property Non exist")
	}

	//check user if exist

	//check if privilege repeatly
	pri, err := p.Pi.SHOW_Property_Userpri(db)
	if err != nil {
		return err
	}
	fmt.Println("pri, p.userpri_update", pri, p.Userpriv_Update)
	exist := if_exist(pri, p.Userpriv_Update)
	if exist == true {
		return errors.New("privilege exist!!!")
	}

	new_pri := append(pri, p.Userpriv_Update...)
	//Update sql
	update := ` UPDATE property SET privilege_for_users = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Update_Teampriv) UPDATE_Property_Teampriv(db *sql.DB) error {
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("Insert Property Non exist")
	}

	//if privilege for team is nil just return

	pri, err := p.Pi.SHOW_Property_Teampri(db)
	fmt.Println("team pri is ", pri)
	if err != nil {
		return err
	}

	for _, i := range p.Teampriv_Update {
		exist := Check_Team_Exist(db, i)
		if exist != true {
			return errors.New("team not exist")
		}
	}

	//check if privilege repeatly

	exist := if_exist(pri, p.Teampriv_Update)
	if exist == true {
		return errors.New("privilege exist!!!")
	}

	new_pri := append(pri, p.Teampriv_Update...)
	//Update sql
	update := ` UPDATE property SET privilege_for_teams = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Update_Departpriv) UPDATE_Property_Departpriv(db *sql.DB) error {
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("Insert Property Non exist")
	}

	pri, err := p.Pi.SHOW_Property_Departpri(db)
	if err != nil {
		return err
	}
	for _, i := range p.Departpriv_Update {
		exist := Check_Depart_Exist(db, i)
		if exist != true {
			return errors.New("department not exist")
		}
	}

	//check if privilege repeatly
	exist := if_exist(pri, p.Departpriv_Update)
	if exist == true {
		return errors.New("privilege exist!!!")
	}

	new_pri := append(pri, p.Departpriv_Update...)

	//Update sql
	update := ` UPDATE property SET privilege_for_departments = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Update_Bupriv) UPDATE_Property_Bupartpriv(db *sql.DB) error {
	ret, err := Check_Property_Exist(db, p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		return err
	}
	if ret != true {
		//non exist
		return errors.New("Insert Property Non exist")
	}

	pri, err := p.Pi.SHOW_Property_Bupri(db)
	if err != nil {
		return err
	}

	for _, i := range p.Bupriv_Update {
		exist := Check_Bu_Exist(db, i)
		if exist != true {
			return errors.New("BU not exist")
		}
	}

	exist := if_exist(pri, p.Bupriv_Update)
	if exist == true {
		return errors.New("privilege exist!!!")
	}

	new_pri := append(pri, p.Bupriv_Update...)

	//Update sql

	update := ` UPDATE property SET privilege_for_bus = $1 WHERE properties_type = $2 AND address = $3 AND hostname = $4;`

	_, err = db.Query(update, pq.Array(new_pri), p.Pi.Properties_type, p.Pi.Ipv4_addr, p.Pi.Hostname)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (p *Properties_Identify) SHOW_Property_Userpri(db *sql.DB) ([]string, error) {
	var pri []string

	//rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM property WHERE hostname = $1) ;", hostname)
	show := `SELECT privilege_for_users FROM property WHERE  properties_type = $1 AND address = $2 AND hostname = $3;`
	rows, err := db.Query(show, p.Properties_type, p.Ipv4_addr, p.Hostname)
	if err != nil {
		logger.Warning.Println("can't EXEC sql: ", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(pq.Array(&pri))
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	//fmt.Println("rows is ", rows)
	return pri, nil

}

func (p *Properties_Identify) SHOW_Property_Teampri(db *sql.DB) ([]string, error) {
	var pri []string

	//rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM property WHERE hostname = $1) ;", hostname)
	show := `SELECT privilege_for_teams FROM property WHERE  properties_type = $1 AND address = $2 AND hostname = $3;`
	rows, err := db.Query(show, p.Properties_type, p.Ipv4_addr, p.Hostname)
	if err != nil {
		logger.Warning.Println("can't EXEC sql: ", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(pq.Array(&pri))
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	fmt.Println("pri is ", pri)
	return pri, nil

}

func (p *Properties_Identify) SHOW_Property_Departpri(db *sql.DB) ([]string, error) {
	var pri []string

	//rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM property WHERE hostname = $1) ;", hostname)
	show := `SELECT privilege_for_departments FROM property WHERE  properties_type = $1 AND address = $2 AND hostname = $3;`
	rows, err := db.Query(show, p.Properties_type, p.Ipv4_addr, p.Hostname)
	if err != nil {
		logger.Warning.Println("can't EXEC sql: ", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(pq.Array(&pri))
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	//fmt.Println("rows is ", rows)
	return pri, nil

}

func (p *Properties_Identify) SHOW_Property_Bupri(db *sql.DB) ([]string, error) {
	var pri []string

	//rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM property WHERE hostname = $1) ;", hostname)
	show := `SELECT privilege_for_bus FROM property WHERE  properties_type = $1 AND address = $2 AND hostname = $3;`
	rows, err := db.Query(show, p.Properties_type, p.Ipv4_addr, p.Hostname)
	if err != nil {
		logger.Warning.Println("can't EXEC sql: ", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(pq.Array(&pri))
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	//fmt.Println("rows is ", rows)
	return pri, nil

}

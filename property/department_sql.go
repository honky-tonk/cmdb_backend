package property

import (
	"database/sql"
	"errors"

	"cmdb_backend/logger"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func (de *Department) Create_Department(db *sql.DB) error {
	//check department exist
	exist := Check_Depart_Exist(db, de.Department_Name)
	if exist == true {
		//department exist
		return errors.New("Department exist")
	}

	//check bu exist
	exist = Check_Bu_Exist(db, de.BU_Belong_to)
	if exist != true {
		//Bu not exist
		return errors.New("BU Not exist")
	}

	insert := `INSERT INTO department (department_name, bu_belong_to, teams) VALUES ($1, $2, $3);`
	_, err := db.Query(insert, de.Department_Name, de.BU_Belong_to, pq.Array(de.Teams))

	if err != nil {
		logger.Error.Fatal("can't Query database: ", err)
	}
	return nil
}

func Show_Teams_In_Depart(db *sql.DB, department_name string) ([]string, error) {
	var s []string
	Show := `SELECT teams FROM department WHERE department_name = $1;`
	rows, err := db.Query(Show, department_name)
	if err != nil {
		logger.Error.Fatal("SQL can't EXEC ", err)
	}

	for rows.Next() {
		err = rows.Scan(pq.Array(&s))
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	//fmt.Println("rows is ", rows)
	return s, nil

}

func Check_Depart_Exist(db *sql.DB, department_name string) bool {
	var exist bool
	Show := `SELECT EXISTS(SELECT 1  FROM department WHERE department_name = $1);`
	rows, err := db.Query(Show, department_name)
	if err != nil {
		logger.Error.Fatal("SQL can't exec", err)
	}

	for rows.Next() {
		err = rows.Scan(&exist)
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
		//fmt.Println("hostnames is ", ret)
	}
	return exist
}

func (d *Department) Delete_Depart_By_Name(db *sql.DB) error {
	//check if bu exist
	exist := Check_Depart_Exist(db, d.Department_Name)

	if exist != true {
		//department not exist
		return errors.New("depart Not exist")
	}

	delete := `DELETE FROM department WHERE department_name = $1;`
	_, err := db.Query(delete, d.Department_Name)
	if err != nil {
		logger.Error.Fatal("can't Query database: ", err)
	}
	return nil
}

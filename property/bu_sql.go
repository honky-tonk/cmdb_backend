package property

import (
	"database/sql"
	"errors"

	"cmdb_backend/logger"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func Check_Bu_Exist(db *sql.DB, Bu_name string) bool {
	var exist bool
	Show := `SELECT EXISTS (SELECT 1 FROM bu WHERE bu_name = $1);`
	rows, err := db.Query(Show, Bu_name)
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

func (b *BU) Create_Bu(db *sql.DB) error {
	exist := Check_Bu_Exist(db, b.BU_Name)
	//fmt.Println("--------------*****-------------")
	if exist == true {
		//Bu exist
		return errors.New("Bu exist")
	}

	insert := `INSERT INTO bu (bu_name, departments) VALUES ($1, $2);`
	_, err := db.Query(insert, b.BU_Name, pq.Array(b.Departments))
	if err != nil {
		logger.Error.Fatal("can't Query database: ", err)
	}
	return nil
}

func (b *BU) Delete_BU_By_Name(db *sql.DB) error {
	//check if bu exist
	exist := Check_Bu_Exist(db, b.BU_Name)

	if exist != true {
		//Bu not exist
		return errors.New("Bu Not exist")
	}

	delete := `DELETE FROM bu WHERE bu_name = $1;`
	_, err := db.Query(delete, b.BU_Name)
	if err != nil {
		logger.Error.Fatal("can't Query database: ", err)
	}
	return nil
}

package property

import (
	"database/sql"
	"errors"
	"cmdb_backend/logger"

	_ "github.com/lib/pq"
)

func (t *Team) Create_Team(db *sql.DB) error {
	//if team_name exist
	exist := Check_Team_Exist(db, t.Team_Name)
	if exist == true {
		return errors.New("Create error Team exist!!! ")
	}

	//if department_belong to is exist
	exist = Check_Depart_Exist(db, t.Department_Belong_to)
	if exist != true {
		return errors.New("Create Team error, department_belong_to not exist!!! ")
	}

	create := `INSERT INTO team (team_name, department_belong_to) VALUES ($1, $2);`
	_, err := db.Query(create, t.Team_Name, t.Department_Belong_to)
	if err != nil {
		logger.Error.Fatal("can't INSERT: ", err)
	}
	return nil

}

func (t *Team) Delete_Team_By_Name(db *sql.DB) error {
	//if team_name exist
	exist := Check_Team_Exist(db, t.Team_Name)
	if exist != true {
		//not exist
		return errors.New("Delete error Team not exist!!! ")
	}

	delete := `DELETE FROM team WHERE team_name = $1;`
	_, err := db.Query(delete, t.Team_Name)
	if err != nil {
		logger.Error.Fatal("can't INSERT: ", err)
	}
	return nil

}

func (t *Team_Update_Departbelong) Update_depart_belong(db *sql.DB) error {
	//if team exist in department
	exist, err := Show_Teams_In_Depart(db, t.Team_Need_Update.Department_Belong_to)
	if len(exist) == 0 {
		//not exist
		return errors.New("team " + t.Team_Need_Update.Team_Name + " Not belong to " + t.Team_Need_Update.Department_Belong_to)
	}

	update := `UPDATE team SET department_belong_to = $1 WHERE team_name = $2;`

	_, err = db.Query(update, t.New_Departbelong, t.New_Departbelong)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func (t *Team_Update_Teamname) Update_depart_name(db *sql.DB) error {
	update := `UPDATE team SET team_name = $1 WHERE team_name = $2;`
	_, err := db.Query(update, t.New_Team_name, t.Team_Need_Update.Team_Name)
	if err != nil {
		logger.Error.Fatal("can't update database: ", err)
	}
	return nil
}

func Check_Team_Exist(db *sql.DB, Team_Name string) bool {
	var exist bool
	//check team name is exist
	show := `SELECT EXISTS(SELECT 1  FROM team WHERE team_name = $1);`
	rows, err := db.Query(show, Team_Name)
	if err != nil {
		logger.Error.Fatal("can't exec sql: ", err)
	}

	for rows.Next() {
		err := rows.Scan(&exist)
		if err != nil {
			logger.Error.Fatal("can't scan result: ", err)
		}
	}
	return exist

}

package property

type Team struct {
	Department_Belong_to string `json:"department_belong_to"`
	Team_Name            string `json:"team_name"`
	//Team_ID              int    `json:"team_id"`
}

type Team_Update_Departbelong struct {
	Team_Need_Update Team   `json:"team_need_update"`
	New_Departbelong string `json:"new_departbelong"`
}

type Team_Update_Teamname struct {
	Team_Need_Update Team   `json:"team_need_update"`
	New_Team_name    string `json:"new_team_name"`
}

func CreateTeam(depart_name string, name string, team_name string) *Team {
	team_Ins := Team{
		Department_Belong_to: depart_name,
		Team_Name:            team_name,
		//Team_ID:              team_id,
	}

	return &team_Ins
}

package mysql

type Group struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetGroupsByPhase(phaseId int64) ([]Group, error) {
	var groups []Group
	err := db.Table("groups").
		Where("phase_id =?", phaseId).
		Scan(&groups).Error
	return groups, err
}

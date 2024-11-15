package mysql

type Phase struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetPhasesByEvent(eventId int64) ([]Phase, error) {
	var phases []Phase
	err := db.Table("phase").
		Where("event_id =?", eventId).
		Scan(&phases).Error
	if err != nil {
		return nil, err
	}
	return phases, nil
}

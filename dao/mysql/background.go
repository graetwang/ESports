package mysql

type Background struct {
	Picture string `json:"picture"`
	Color   string `json:"color"`
}

func BackgroundList() ([]Background, error) {
	var background []Background
	err := db.Debug().
		Table("background").
		Select("picture, color").
		Scan(&background).Error
	return background, err
}

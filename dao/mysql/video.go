package mysql

type Video struct {
	Section    string `json:"section" gorm:"section"`
	Title      string `json:"title" gorm:"title"`
	Cover      string `json:"cover" gorm:"cover"`
	Duration   string `json:"duration" gorm:"duration"`
	ViewCounts int64  `json:"view_counts" gorm:"view_counts"`
	Likes      int64  `json:"likes" gorm:"likes"`
	Author     string `json:"author" gorm:"author"`
	Address    string `json:"address" gorm:"address"`
}

func GetVideoList() ([]Video, error) {
	var videos []Video
	err := db.Table("videos").
		Order("view_counts DESC").
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func GetVideoListBySection(section string) ([]Video, error) {
	var videos []Video
	err := db.Table("videos").
		Where("section=?", section).
		Order("view_counts DESC").
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func GetVideoListBySectionLimit(section string, pageSize, page int) ([]Video, error) {
	var videos []Video
	offset := (page - 1) * pageSize
	err := db.Table("videos").
		Where("section=?", section).
		Order("view_counts DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

package mysql

type Banner struct {
	Title   string `json:"title" gorm:"title"`
	Picture string `json:"picture" gorm:"picture"`
}

func BannerList() ([]Banner, error) {
	var banners []Banner
	err := db.Debug().
		Table("banner").
		Select("title, picture").
		Scan(&banners).Error
	return banners, err
}

package service

import (
	"server/dao/mysql"
)

func Banner() ([]mysql.Banner, error) {
	banners, err := mysql.BannerList()
	if err != nil {
		return nil, err
	}
	return banners, nil
}

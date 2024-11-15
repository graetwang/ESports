package service

import (
	"server/dao/mysql"
)

func GetVideosBySection(section string) ([]mysql.Video, error) {
	videos, err := mysql.GetVideoListBySection(section)
	return videos, err
}
func GetVideosBySectionLimit(section string, pageSize, page int) ([]mysql.Video, error) {
	videos, err := mysql.GetVideoListBySectionLimit(section, pageSize, page)
	return videos, err
}

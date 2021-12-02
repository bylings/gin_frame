package services

import "github.com/bylings/gin/models"

func GetRegionsDetail(where map[string]interface{}) models.Regions {
	var regions models.Regions
	models.DB().Debug().Where(where).Unscoped().Find(&regions)
	return regions
}

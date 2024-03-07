package restaurentstorage

import (
	"context"
	"food/common"
	"food/module/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db
	var result []restaurantmodel.Restaurant
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	db = db.Table(restaurantmodel.RestaurantCreate{}.
		TableName()).
		Where(condition).
		Where("status in (1)")

	if v := filter; v != nil {
		if v.Cityid > 0 {
			db = db.Where("city_id = ?", v.Cityid)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

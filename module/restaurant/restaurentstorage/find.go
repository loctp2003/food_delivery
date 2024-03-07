package restaurentstorage

import (
	"context"
	"food/module/restaurant/restaurantmodel"
)

func (s *sqlStore) FinDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

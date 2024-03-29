package restaurentstorage

import (
	"context"
	"food/module/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

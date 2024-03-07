package restaurantbiz

import (
	"context"
	"errors"
	"food/module/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
	FinDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}
type deteleRestaurant struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deteleRestaurant {
	return &deteleRestaurant{store: store}
}
func (biz *deteleRestaurant) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FinDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data deleted")
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}

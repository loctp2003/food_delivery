package uploadstorage

import (
	"context"
	"food/common"
)

func (s *sqlStore) CreateImage(context context.Context, data *common.Image) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

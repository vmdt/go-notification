package postgres

import (
	"context"

	"gorm.io/gorm"
)

type genericRepository[T any] struct {
	db *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) *genericRepository[T] {
	return &genericRepository[T]{db: db}
}

func (r *genericRepository[T]) Add(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *genericRepository[T]) AddAll(entity *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *genericRepository[T]) GetById(id int, ctx context.Context) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).FirstOrInit(&entity).Error

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *genericRepository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return &entities, nil
}

func (r *genericRepository[T]) Where(param *T, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Where(&param).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return &entities, nil
}

func (r *genericRepository[T]) Update(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *genericRepository[T]) UpdateAll(entities *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *genericRepository[T]) Delete(id int, ctx context.Context) error {
	var entity T
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity).Error
}

func (r *genericRepository[T]) SkipTake(skip int, take int, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return &entities, nil
}

func (r *genericRepository[T]) Count(ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count
}

func (r *genericRepository[T]) CountWhere(params *T, ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Where(&params).Count(&count)
	return count
}
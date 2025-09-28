package repo

import "gorm.io/gorm"

type CrudEntity[T any] interface {
	Create(entity T) (T, error)
	GetById(id string) (*T, error)
	Delete(id string) error
	Update(id string, entity *T) (T, error)
	GetAll() ([]T, error)
}

type BaseModel[T any] struct {
	Db *gorm.DB
}

// Constructor
func NewBaseModel[T any](db *gorm.DB) *BaseModel[T] {
	return &BaseModel[T]{Db: db}
}

// Create
func (b *BaseModel[T]) Create(entity T) (T, error) {
	if err := b.Db.Create(&entity).Error; err != nil {
		var zero T
		return zero, err
	}
	return entity, nil
}

// GetById
func (b *BaseModel[T]) GetById(id string) (*T, error) {
	var entity T
	if err := b.Db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete
func (b *BaseModel[T]) Delete(id string) error {
	var entity T
	return b.Db.Delete(&entity, "id = ?", id).Error
}

// Update
func (b *BaseModel[T]) Update(id string, entity *T) (T, error) {
	var existing T

	// Check if record exists
	if err := b.Db.First(&existing, "id = ?", id).Error; err != nil {
		var zero T
		return zero, err
	}

	// Update fields
	if err := b.Db.Model(&existing).Updates(entity).Error; err != nil {
		var zero T
		return zero, err
	}

	return *entity, nil
}

// GetAll
func (b *BaseModel[T]) GetAll() ([]T, error) {
	var entities []T
	if err := b.Db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

package repo

import "gorm.io/gorm"

type CrudEntity[T any] interface {
	Create(entity T) error
	GetById(id string) (*T, error)
	Delete(id string) error
	Update(id string, entity *T) (T, error)
	GetAll() ([]T, error)
}

type BaseModel[T any] struct {
	Db *gorm.DB
}

func (b *BaseModel[T]) Create(entity T) error {
	return b.Db.Create(&entity).Error
}

func (b *BaseModel[T]) GetById(id string) (*T, error) {
	var entity T
	if err := b.Db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (b *BaseModel[T]) Delete(id string) error {
	var entity T
	return b.Db.Delete(&entity, "id = ?", id).Error
}

func (b *BaseModel[T]) Update(id string, entity *T) (T, error) {
	var existing T

	// Pehle check karo ki record exist karta hai ya nahi
	if err := b.Db.First(&existing, "id = ?", id).Error; err != nil {
		var zero T
		return zero, err
	}

	// Ab update karo
	if err := b.Db.Model(&existing).Updates(entity).Error; err != nil {
		var zero T
		return zero, err
	}

	return *entity, nil
}

func (b *BaseModel[T]) GetAll() ([]T, error) {
	var entities []T
	if err := b.Db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

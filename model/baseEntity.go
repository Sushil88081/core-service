package baseEntity

import "time"

type IEntity interface {
	TableName() string
}
type BaseEntity struct {
	ID        string     `json:"id" bson:"_id" gorm:"primaryKey;column:id"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" bson:"deletedAt"`
	CreatedBy string     `json:"createdBy" bson:"createdBy"`
	UpdatedBy string     `json:"updatedBy" bson:"updatedBy"`
	DeletedBy *string    `json:"deletedBy,omitempty" bson:"deletedBy"`
	ProjectID string     `json:"projectId" bson:"projectId"`
}

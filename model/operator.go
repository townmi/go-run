package model

import (
	"time"
)

// Operaotr Structure reflect MYSQL
type Operaotr struct {
	ID          string     `json:"id" gorm:"column:id; type:char(36); primary_key; not null"`
	DisplayName *string    `json:"displayName" gorm:"column:displayName; type:varchar(255);"`
	Password    *string    `json:"password, omitempty" gorm:"column:password; type:varchar(255);"`
	Role        *string    `json:"role" gorm:"column:role; type:varchar(255);"`
	Email       *string    `json:"email" gorm:"column:email; type:varchar(255);"`
	Mobile      string     `json:"mobile" gorm:"column:mobile; type:varchar(255); not null"`
	LocationLat *string    `json:"locationLat" gorm:"column:locationLat; type:varchar(255)"`
	LocationLon *string    `json:"locationLon" gorm:"column:locationLon; type:varchar(255)"`
	CityID      *string    `json:"cityId" gorm:"column:cityId; type:varchar(255);"`
	Secret      *string    `json:"secret" gorm:"column:secret; type:varchar(255);"`
	WechatID    *string    `json:"wechatId" gorm:"column:WechatId; type:varchar(36);"`
	Level       *int       `json:"level" gorm:"column:level; type:int(10);"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:createdAt; type:datetime;"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt; type:datetime;"`
}

package model

import (
	"time"
)

// Venues Structure reflect MYSQL
type Venues struct {
	ID              string     `json:"id" gorm:"column:id; type:char(36); primary_key; not null"`
	DisplayName     *string    `json:"displayName" gorm:"column:displayName; type:varchar(255);"`
	Password        *string    `json:"password, omitempty" gorm:"column:password; type:varchar(255);"`
	Role            *string    `json:"role" gorm:"column:role; type:varchar(255);"`
	Email           *string    `json:"email" gorm:"column:email; type:varchar(255);"`
	Mobile          string     `json:"mobile" gorm:"column:mobile; type:varchar(255); not null"`
	LocationLat     *string    `json:"locationLat" gorm:"column:locationLat; type:varchar(255)"`
	LocationLon     *string    `json:"locationLon" gorm:"column:locationLon; type:varchar(255)"`
	Subscribe       *int       `json:"subscribe" gorm:"column:subscribe; type:int(255);"`
	CityID          *string    `json:"cityId" gorm:"column:cityId; type:varchar(255);"`
	Secret          *string    `json:"secret" gorm:"column:secret; type:varchar(255);"`
	MiniWechatID    *string    `json:"miniWechatId" gorm:"column:miniWechatId; type:char(36);"`
	WechatID        *string    `json:"wechatId" gorm:"column:WechatId; type:varchar(36);"`
	Level           *int       `json:"level" gorm:"column:level; type:int(11);"`
	CreatedAt       *time.Time `json:"createdAt" gorm:"column:createdAt; type:datetime;"`
	BindMobileAt    *time.Time `json:"bindMobileAt" gorm:"column:bindMobileAt; type:datetime;"`
	UpdatedAt       *time.Time `json:"updatedAt" gorm:"column:updatedAt; type:datetime;"`
	Consignee       *string    `json:"consignee" gorm:"column:consignee; type:varchar(255);"`
	ShippingAddress *string    `json:"shippingAddress" gorm:"column:shippingAddress; type:varchar(255);"`
	ReceivingPhone  *string    `json:"receivingPhone" gorm:"column:receivingPhone; type:varchar(255);"`
	IsPolicy        *int8      `json:"isPolicy" gorm:"column:isPolicy; type:tinyint(4);"`
	Postcode        *string    `json:"postcode" gorm:"column:Postcode; type:varchar(255);"`
	Unionid         *int       `json:"unionid" gorm:"column:unionid; type:int(11);"`
}

package models

type Order struct{
	GormModel
	CustomerName  string 		`gorm:"not null;" json:"customer_name" form:"customer_name" valid:"required"`
	Items					[]Items		`gorm:"foreignKey:OrderID"`
}

type Items struct{
	GormModel
	ItemCode 		string	`gorm:"not null;" json:"itemcode"`
	Description string	`gorm:"not null;" json:"description"`
	Quantity		int			`gorm:"not null;" json:"quantity"`
	OrderID			uint		`gorm:"not null;" json:"orderid"`
}
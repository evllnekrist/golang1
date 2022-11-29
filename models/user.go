package models

import (
	// "database/sql/driver"
	"photo-app/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type levelmodel string

const (
	Superadmin  levelmodel = "superadmin"
	Admin 			levelmodel = "admin"
	Users 			levelmodel = "user"
)

type User struct {
	GormModel
	Username 			string    		`gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string    		`gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string    		`gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products  []Product   		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Age       uint      		`gorm:"not null;" json:"age" form:"age" valid:"required~Your age is required"`
	Level 	  levelmodel		`sql:"type:levelmodel" gorm:"default:user"` 
	Status	  []Status			`gorm:"foreignKey:UserID"`
}
type Status struct{
	GormModel
	Status string 		`gorm:"not null" json:"status"`
	UserID uint				`gorm:"not null;" json:"userid"`
}

type UserResp struct{
	Username string    		`gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string    		`gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Age      uint      		`gorm:"not null;" json:"age" form:"age" valid:"required~Your age is required"`
	Level 	 levelmodel		`sql:"type:levelmodel" gorm:"default:user"` 
	Status	string				`gorm:"foreignKey;UserID" json:"status" `
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return

}

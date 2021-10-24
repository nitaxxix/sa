package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"uniqueIndex"`
	Pass     string
	Appoints []Appoint `gorm:"foreignKey:UserID"`
	RoleID   *uint
	Role     Role `gorm:"foreignKey:RoleID"`
}

type Role struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Users []User `gorm:"foreignKey:RoleID"`
}

type Patient struct {
	gorm.Model
	Name     string
	Sex      string
	Age      int
	IDcard   string `gorm:"uniqueIndex"`
	Time     time.Time
	Tel      string
	Appoints []Appoint `gorm:"foreignKey:PatientID"`
}

type RemedyType struct {
	gorm.Model
	Name     string
	Appoints []Appoint `gorm:"foreignKey:RemedyTypeID"`
}

type Appoint struct {
	gorm.Model
	AppointTime  time.Time
	Todo         string
	UserID       *uint
	User         User
	PatientID    *uint
	Patient      Patient
	RemedyTypeID *uint
	RemedyType   RemedyType
}

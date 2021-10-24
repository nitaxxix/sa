package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	database.AutoMigrate(
		&User{}, &Role{}, &Patient{}, &Appoint{}, &RemedyType{},
	)
	db = database

	role1 := Role{
		Name: "ทันตแพทย์",
	}
	db.Model(&Role{}).Create(&role1)

	role2 := Role{
		Name: "ผู้ช่วยทันตแพทย์",
	}
	db.Model(&Role{}).Create(&role2)

	role3 := Role{
		Name: "พยาบาล",
	}
	db.Model(&Role{}).Create(&role3)

	role4 := Role{
		Name: "เภสัชกร",
	}
	db.Model(&Role{}).Create(&role4)

	role5 := Role{
		Name: "พนักงานการเงิน",
	}
	db.Model(&Role{}).Create(&role5)

	dentist1 := User{
		Name:     "กอเอ๋ย กอไก่",
		Username: "nita",
		Pass:     "5678",
		Role:     role1,
	}
	db.Model(&User{}).Create(&dentist1)

	dentist2 := User{
		Name:     "ขอไข่ ในเล้า",
		Username: "name",
		Pass:     "1234",
		Role:     role1,
	}
	db.Model(&User{}).Create(&dentist2)

	remedy1 := RemedyType{
		Name: "อุดฟัน",
	}
	db.Model(&RemedyType{}).Create(&remedy1)

	remedy2 := RemedyType{
		Name: "ขูดหินปูน",
	}
	db.Model(&RemedyType{}).Create(&remedy2)

	remedy3 := RemedyType{
		Name: "เอ็กซ์เรย์",
	}
	db.Model(&RemedyType{}).Create(&remedy3)

	patient1 := Patient{
		Name:   "บอใบไม้ ทับถม",
		Sex:    "ชาย",
		Age:    30,
		IDcard: "1302153974578",
		Time:   time.Now(),
		Tel:    "0625789634",
	}
	db.Model(&Patient{}).Create(&patient1)

	patient2 := Patient{
		Name:   "ปอปลา ตากลม",
		Sex:    "หญิง",
		Age:    21,
		IDcard: "1197586345245",
		Time:   time.Now(),
		Tel:    "0852187233",
	}
	db.Model(&Patient{}).Create(&patient2)

	db.Model(&Appoint{}).Create(&Appoint{
		AppointTime: time.Now(),
		Todo:        "งดน้ำ 3 ชั่วโมง",
		User:        dentist1,
		Patient:     patient1,
		RemedyType:  remedy1,
	})

	db.Model(&Appoint{}).Create(&Appoint{
		AppointTime: time.Now(),
		Todo:        "-",
		User:        dentist1,
		Patient:     patient2,
		RemedyType:  remedy2,
	})

	db.Model(&Appoint{}).Create(&Appoint{
		AppointTime: time.Now(),
		Todo:        "งดอาหาร 12 ชั่วโมง",
		User:        dentist1,
		Patient:     patient1,
		RemedyType:  remedy3,
	})

}

package models

import (
	"dtby/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type H struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
	OccupationName string `gorm:"type:varchar(100);unique_index"`
	NormalName string `gorm:"type:varchar(100);unique_index"`
	HitPoint int
	ManaPoint int
	DamagePerSecond float64
	Damage float64
	AttackSpeed float64
	MoveSpeed float64
	AttackRange int
	MagicResist float64
	Armor float64
	RegenerationValue float64
	Level int
}

type HeroLevel struct {
	gorm.Model
	Level int
	Price int
	ValueDescribe string `gorm:"type:varchar(20)"`
	H []H `gorm:"ForeignKey:Level"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.Config.DSN)
	if err!=nil{
		fmt.Printf("GORM: open error: %v\n", err)
		return nil ,err
	}
	DB=db
	db.SingularTable(true)
	gorm.DefaultTableNameHandler(db, "dtby_")
	db.AutoMigrate(&H{}, &HeroLevel{})
	return db, err
}

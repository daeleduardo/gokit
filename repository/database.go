package database

import (
	Model "gokit/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//Get Database Connection
func GetConn() (*gorm.DB, error) {
	//Connection a Sqlite file for example purpose, in production the property config.GetConfig().DbConn, must have to used.
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	//Auto Migrase used for example purpose, in production the one must be evaluate.
	err = db.AutoMigrate(&Model.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Create(user Model.User) *gorm.DB {
	db, _ := GetConn()
	return db.Create(&user)
}

func ReadOne(user Model.User) (Model.User, error) {
	db, err := GetConn()
	if err != nil {
		return Model.User{}, err
	}
	var result Model.User
	db.Model(&Model.User{}).First(&result, user.Id)
	return result, nil
}

func ReadAll() ([]Model.User, error) {
	db, err := GetConn()
	if err != nil {
		return nil, err
	}
	var result []Model.User
	db.Find(&result)
	return result, nil
}

func Update(user Model.User) *gorm.DB {
	db, _ := GetConn()
	return db.Model(Model.User{}).Where("id = ?", user.Id).Updates(user)
}

func Delete(user Model.User) *gorm.DB {
	db, _ := GetConn()
	return db.Delete(&user, user.Id)
}

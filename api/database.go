package church

import (
    "gorm.io/gorm"
)

var db * gorm.DB
var err error = nil

func SignDB(con_str string, createDB func (string) (*gorm.DB, error)) (*gorm.DB, error) {
    db, err = createDB(con_str)

    db.AutoMigrate(&User{})
    return db, err
}

package church

import (
    "github.com/Coff3e/Api"
    "gorm.io/gorm"
)

var db * gorm.DB
var err error = nil
var _models [] interface{}

func SignDB(con_str string) (*gorm.DB, error) {
    _models = [] interface{} {
        &User{},
    }

    db, err = api.CreateDB(con_str)

    db.AutoMigrate(_models...)
    return db, err
}

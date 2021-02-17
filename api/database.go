package church

import (
    "github.com/Coff3e/Api"
    "gorm.io/gorm"
)

var db * gorm.DB
var err error = nil

func (router *Church) SignDB(con_str string, createDB func (string) (*gorm.DB, error)) (*gorm.DB, error) {
    db, err = createDB(con_str)
    if err == nil {
        models := []interface{} {
            &User{},
            &Token{},

            &Role{},
            &UserRole{},

            &Celule{},
            &UserCelule{},

            &Event{},
            &UserEvent{},
        }

        db.AutoMigrate(models...)

        if err == nil {
            return db, err
        } else {
            api.Die("Error on creation of tables on database")
        }
    }

    router.Database = db
    return router.Database, err
}

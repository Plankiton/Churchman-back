package main

import (
    "os"
    "time"

    "../api"
    "github.com/Coff3e/Api"
)

func main() {
    con_str := ":memory:"
    router := church.Church{}
    db, err := router.SignDB(con_str, api.Sqlite)
    if (err != nil) {
        os.Exit(1)
    }

    api.Log("Database connected with sucess")

    api.Log("Trying to create User on database")
    user := church.User{}
    user.ModelType = "User"
    user.Name = "Joao da Silva"
    user.Email = "joao@j.com"
    user.Phone = "99 8452 1107"
    user.Genre = "M"
    user.State = "sole"
    user.Born = time.Now()

    user.Create()
    _ = db.First(&user)

    api.Log("Trying to create Celule on database")
    celule := church.Celule{}
    celule.Name = "Anjos da morte"
    celule.Create()

    api.Log("Trying to Sign User to Celule")
    celule.Sign(user)

    api.Log("Trying to Unsign User to Celule")
    celule.Unsign(user)
}

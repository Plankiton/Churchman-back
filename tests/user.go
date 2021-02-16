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

    user := church.User{}
    user.ModelType = "User"
    user.Name = "Joao da Silva"
    user.Email = "joao@j.com"
    user.Phone = "99 8452 1107"
    user.Genre = "M"
    user.State = "sole"
    user.Born = time.Now()

    api.Log("Trying to create User on database")
    user.Create()

    api.Log("Trying to read User on database")
    r := db.First(&user)
    if (r.Error == nil) {
        api.Log("Readed", api.ToLabel(user.ID, user.ModelType))
    }

    api.Log("Trying to update User name from ", user.Name)
    user.Update(api.Dict {
            "name": "Joao Maria das Dores",
        })
    _ = db.First(&user, user.ID)
    api.Log("Updated to ", user.Name)

    api.Log("Trying to delete User on database")
    user.Delete()
}

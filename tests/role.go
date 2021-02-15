package main

import (
    "os"
    "time"

    "../api"
    "github.com/Coff3e/Api"
)

func main() {
    con_str := ":memory:"
    _, err := church.SignDB(con_str, api.Sqlite)
    if (err != nil) {
        os.Exit(1)
    }

    api.Log("Database connected with sucess")

    api.Log("Trying to create a Role")
    person := church.Role{}
    person.Name = "Person"
    person.Create()

    api.Log("Trying to create a User")
    user := church.User{}
    user.Name = "Joao da Silva"
    user.Email = "joana@j.com"
    user.Genre = "M"
    user.State = "married"
    user.Born = time.Now()
    user.Create()

    api.Log("Trying to link Role to User")
    person.Sign(user)

    api.Log("Trying to unlink Role to User")
    person.Unsign(user)
}

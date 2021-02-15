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

    api.Log("Trying to create Roles")
    person := church.Role{}
    person.Name = "Person"
    person.Create()

    poor := church.Role{}
    poor.Name = "Poor"
    poor.Create()

    api.Log("Trying to create a User")
    user := church.User{}
    user.Name = "Joao da Silva"
    user.Email = "joana@j.com"
    user.Genre = "M"
    user.State = "married"
    user.Born = time.Now()
    user.Create()

    api.Log("Trying to link Roles to User")
    person.Sign(user)
    poor.Sign(user)

    null_user := church.User{}
    null_user.Name = "NUll User"
    null_user.Create()
    person.Sign(null_user)

    api.Log("Trying to get User list from Role")
    print("\tPerson users -> [ ")
    for _, u := range person.GetUsers() {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")
    print("\tPoor users -> [ ")
    for _, u := range poor.GetUsers() {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")

    api.Log("Trying to get Role list from User")
    print("\t", user.Name, "-> [ ")
    for _, u := range user.GetRoles() {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")
    print("\t", null_user.Name, "-> [ ")
    for _, u := range null_user.GetRoles() {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")

    api.Log("Trying to unlink Role to User")
    person.Unsign(user)
}

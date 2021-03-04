package main

import (
    "os"
    "time"

    "github.com/Coff3e/Church-app/api"
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
    g12 := church.Role{}
    g12.Create()

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
    g12.Sign(user)

    api.Log("Trying to create Celule on database")
    celule := church.Celule{}
    celule.Leader = user.ID
    celule.Create()

    api.Log("Trying to Sign User to Celule")
    celule.Sign(user)

    api.Log("Trying to Get User list from Celule")
    print("\tPoor users -> [ ")
    for _, u := range celule.GetUsers(0, 0) {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")

    api.Log("Trying to get Celule list from User")
    print("\t", user.Name, "-> [ ")
    for _, u := range user.GetCelules(0, 0) {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")

    api.Log("Trying to Unsign User to Celule")
    celule.Unsign(user)

    api.SuperPut(celule.Name)

    celule_child := church.Celule{}
    celule_child.Leader = user.ID
    celule_child.Parent = celule.ID
    celule_child.Create()

    celule_child2 := church.Celule{}
    celule_child2.Leader = user.ID
    celule_child2.Parent = celule.ID
    celule_child2.Create()

    celule_child3 := church.Celule{}
    celule_child3.Leader = user.ID
    celule_child3.Parent = celule_child2.ID
    celule_child3.Create()
}

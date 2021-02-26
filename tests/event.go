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

    api.Log("Trying to create Event on database")
    event := church.Event{}
    event.Name = "Anjos da morte"
    event.Create()

    api.Log("Trying to Sign User to Event")
    event.Sign(user)

    api.Log("Trying to Get User list from Event")
    print("\tPoor users -> [ ")
    for _, u := range event.GetUsers(1, 5) {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")

    api.Log("Trying to get Event list from User")
    print("\t", user.Name, "-> [ ")
    for _, u := range user.GetEvents(1, 5) {
        print(api.ToLabel(u.ID, u.ModelType), ", ")
    }
    print("]\n")

    api.Log("Trying to Unsign User to Event")
    event.Unsign(user)

}

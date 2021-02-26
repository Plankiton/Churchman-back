package main

import (
    "encoding/json"
    "net/http"
    "bytes"
    "os"

    "github.com/Coff3e/Church-app/api"
    "github.com/Coff3e/Api"
)


func main() {
    con_str := ":memory:"
    r := church.Church{}
    _, err := r.SignDB(con_str, api.Sqlite)
    if (err != nil) {
        os.Exit(1)
    }

    if api.GetEnv("DB_URI", "") == "" {
    }

    r.
    Add(
        "get", "/user/", nil, church.GetUserList,
    ).
    Add(
        "post", "/user", nil, church.CreateUser,
    ).
    Add(
        "post", "/user/{id}/profile", nil, church.CreateUserProfile,
    )

    go r.Run("/", 8000)
    CreateUsers()

    //body := new(bytes.Buffer)
    url := "http://localhost:8000/user/?l=5&p=1"
    res, err := http.Get(url)

    process(res, err)
}

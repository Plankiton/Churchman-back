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
        CreateFakeDB()
    }


    r.
    Add(
        "delete", "/user/{id}", nil, church.DeleteUser,
    )

    go r.Run("/", 8000)

    uid := "1"

    //body := new(bytes.Buffer)
    url := "http://localhost:8000/user/" + uid
    req, err := http.NewRequest(http.MethodDelete, url, nil)
    client := &http.Client{}
    res, err := client.Do(req)

    // Creating User
    r_body := process(res, err)
    if  r_body.Type != "Error" && r_body.Data != nil {
        //data := r_body.Data.(map[string]interface{})
        url += "/profile"
        res, err := http.Get(url)

        process(res, err)
    }
}

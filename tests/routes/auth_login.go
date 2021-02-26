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

    CreateFakeDB()

    r.Add(
        "post", "/login", api.RouteConf {}, church.LogIn,
    )
    go r.Run("/", 8000)

    url := "http://localhost:8000/login"
    for _, d := range []interface{}{
        map[string]string{"email": "maria@j.com", "pass": "maria"},
        map[string]string{"email": "maria@j.com", "pass": "joao"},
        []map[string]string{{"j": "j", "pass": "joao"}},
        []map[interface{}]string{{"j": "j", "pass": "joao"},{0:"joao"}},
        map[string]string{"email": "maria@joao.com", "pass": "joao"},
    } {

        body := new(bytes.Buffer)
        json.NewEncoder(body).Encode(struct {Data interface{} `json:"data"`} {
            Data: d,
        })
        res, err := http.Post(url, "application/json", body)
        process(res, err)

    }
}

package main

import (
    "encoding/json"
    "net/http"
    "bytes"
    "os"

    "../api"
    "github.com/Coff3e/Api"
)

func main() {
    con_str := ":memory:"
    r := church.Church{}
    _, err := r.SignDB(con_str, api.Sqlite)
    if (err != nil) {
        os.Exit(1)
    }

    r.
    Add(
        "post", "/role", nil, church.CreateRole,
    )
    go r.Run("/", 8000)

    url := "http://localhost:8000/role"
    for _, d := range []interface{}{
        map[string]string{"email": "maria@j.com", "pass": "joao"},
        []map[string]string{{"j": "j", "pass": "joao"}},
        []map[interface{}]string{{"j": "j", "pass": "joao"},{0:"joao"}},
        map[string]string{"name": "chato"},
        map[string]string{
            "name": "professor",
        },
    } {

        body := new(bytes.Buffer)
        json.NewEncoder(body).Encode(struct {Data interface{} `json:"data"`} {
            Data: d,
        })
        res, err := http.Post(url, "application/json", body)

        process(res, err)
    }
}

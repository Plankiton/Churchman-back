package main

import (
    "encoding/json"
    "net/http"
    "time"
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
        "post", "/user/{id}", nil, church.UpdateUser,
    )

    go r.Run("/", 8000)

    born := time.Date(2002, time.March, 19, 0, 0, 0, 0, &time.Location{})
    url := "http://localhost:8000/user/1"
    d := map[string]string{
        "email": "maria_de_deus@joao.com",
        "pass": "maria",
        "name": "jose",
        "genre": "M",
        "born": born.Format(church.TimeLayout()),
    }

    body := new(bytes.Buffer)
    json.NewEncoder(body).Encode(struct {Data interface{} `json:"data"`} {
        Data: d,
    })

    res, err := http.Post(url, "application/json", body)
    process(res, err)
}

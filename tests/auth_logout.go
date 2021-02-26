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
    db, err := r.SignDB(con_str, api.Sqlite)
    if (err != nil) {
        os.Exit(1)
    }

    CreateFakeDB()

    valid_token := ""; {
        token := church.Token{}
        db.Take(&token)
        valid_token = token.ID
    }

    r.Add(
        "post", "/logout", api.RouteConf {
            "need-auth": true,
        }, church.LogOut,
    )
    go r.Run("/", 8000)

    url := "http://localhost:8000/logout"
    for _, t := range []interface{} {
        "8dadfa32800",
        0x8dadfa32800,
        valid_token,
    } {

        body := new(bytes.Buffer)
        json.NewEncoder(body).Encode(struct {Token interface{} `json:"auth"`} {
            Token: t,
        })
        res, err := http.Post(url, "application/json", body)
        process(res, err)

    }
}

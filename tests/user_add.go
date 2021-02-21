package main

import (
    "encoding/json"
    "net/http"
    "time"
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

    CreateFakeDB()

    r.Add(
        "post", "/user", api.RouteConf {}, church.CreateUser,
    )
    go r.Run("/", 8000)

    born := time.Date(2002, time.March, 19, 0, 0, 0, 0, &time.Location{})
    url := "http://localhost:8000/user"
    for _, d := range []interface{}{
        map[string]string{"email": "maria@j.com", "pass": "joao"},
        []map[string]string{{"j": "j", "pass": "joao"}},
        []map[interface{}]string{{"j": "j", "pass": "joao"},{0:"joao"}},
        map[string]string{"email": "maria@joao.com", "pass": "joao"},
        map[string]string{
            "email": "j@joao.com",
            "pass": "maria",
            "name": "jose",
            "genre": "M",
            "born": born.Format(church.TimeLayout()),
        },
    } {

        body := new(bytes.Buffer)
        json.NewEncoder(body).Encode(struct {Data interface{} `json:"data"`} {
            Data: d,
        })
        res, err := http.Post(url, "application/json", body)
        if (err == nil){
            r_body := api.Response{}
            json.NewDecoder(res.Body).Decode(&r_body)
            res.Body.Close()

            r_raw_body := new(bytes.Buffer)
            json.NewEncoder(r_raw_body).Encode(r_body)

            if  r_body.Type == "Error" {
                api.War("Sucessfull request with error\n\t-> Response:", r_raw_body)
            } else {
                api.Log("Sucessfull request\n\t-> Response:", r_raw_body)
            }
        } else {
            api.Err("Failure request sending, ", err)
        }

    }
}

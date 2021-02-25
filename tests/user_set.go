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

func process(res *http.Response, err error) api.Response {
    if err == nil {
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

        return r_body

    } else {
        api.Err("Failure request sending, ", err)
    }

    return api.Response {}
}

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

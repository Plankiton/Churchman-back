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
    r.Add("post", "/joao", api.RouteConf {}, func(r *http.Request) (api.Response, int) {
        return api.Response {
            Data: "JOAO É GAYZÃOOO!!",
        }, 200
    }).Add("get", "/name/{name}", api.RouteConf {}, func(r *http.Request) (api.Response, int) {
        vars, err := api.GetPathVars("/name/{name}", r.URL.Path)
        if err != nil {
            return api.Response {
                Data: "Request errada!!",
            }, 400
        }

        return api.Response {
            Data: vars["name"]+" É GAYZÃOOO!!",
        }, 200
    })
    go r.Run("/", 8000)

    for _, url := range []string{"http://localhost:8000/", "http://localhost:8000/joao", "http://localhost:8000/name/maria"} {
        body := new(bytes.Buffer)
        json.NewEncoder(body).Encode(api.Request {
            Token: "1290839028903809",
            Data: map[string]interface{}{
                "Joao": "Maria",
                "Marta": 90,
                "Cellbit": false,
            },
        })
        api.Log("Trying to make requests on", url)

        res, err := http.Get(url)
        if (err == nil) {
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

        res, err = http.Post(url, "application/json", body)
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

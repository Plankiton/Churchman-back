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
    r.Add("post", "/joao", api.RouteConf {}, func(r api.Request) (api.Response, int) {

        return api.Response {
            Data: "JOAO É GAYZÃOOO!!",
        }, 200

    }).Add("get", "/name/{name}", api.RouteConf {}, func(r api.Request) (api.Response, int) {

        return api.Response {
            Data: r.PathVars["name"]+" É GAYZÃOOO!!",
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
        process(res, err)

        res, err = http.Post(url, "application/json", body)
        process(res, err)

    }
}

package main

import (
	"net/http"
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
    })
    r.Run("/", 8000)
}

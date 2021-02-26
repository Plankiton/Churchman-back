package main

import (
    "encoding/json"
    "net/http"
    "time"
    "fmt"
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

    r.
    Add(
        "post", "/user", nil, church.CreateUser,
    ).
    Add(
        "post", "/user/{id}/profile", nil, church.CreateUserProfile,
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

        // Creating User
        r_body := process(res, err)

        if  r_body.Type != "Error" && r_body.Data != nil {
            data := r_body.Data.(map[string]interface{})

            // Sending him profile photo
            buf := new(bytes.Buffer )
            buf.Write([]byte(`--BOIOLAGE-DA-PORRA
Content-Disposition: form-data; name="data"; filename="j"
Content-Type: application/octet-stream

JOAO É GAYZÃO

--BOIOLAGE-DA-PORRA
Content-Disposition: form-data; name="description"

Imagem
--BOIOLAGE-DA-PORRA--`))
            mime := "multipart/form-data; boundary=BOIOLAGE-DA-PORRA"

            url := fmt.Sprint("http://localhost:8000/user/", data["id"].(float64), "/profile")
            res, err := http.Post(url, mime, buf)

            process(res, err)
        }
    }
}

package main

import (
    "bytes"
    "time"
    "fmt"

    "net/http"
    "encoding/json"

    "github.com/Coff3e/Church-app/api"
    "github.com/Coff3e/Api"
)

func CreateFakeDB() {
    joao := church.User{}
    joao.ModelType = "User"
    joao.Name = "Joao da Silva"
    joao.Email = "joao@j.com"
    joao.Phone = "99 8452 1107"
    joao.Genre = "M"
    joao.State = "married"
    joao.Born = time.Now()
    joao.SetPass("maria")

    maria := church.User{}
    maria.Name = "Maria da Silva"
    maria.Email = "maria@j.com"
    maria.Genre = "F"
    maria.Phone = "99 8452 1108"
    maria.State = "married"
    maria.Born = time.Now()
    maria.SetPass("joao")

    person := church.Role{}
    person.Name = "Person"

    poor := church.Role{}
    poor.Name = "Poor"

    maria.Create()
    joao.Create()

    poor.Create()
    person.Create()

    person.Sign(joao)
    poor.Sign(joao)

    person.Sign(maria)

    photo := church.File{}
    photo.AltText = "Image"
    photo.LoadString("JOAO É GAY")

    maria.SetProfile(photo)
    joao_login := church.Token{}
    joao_login.UserId = joao.ID
    joao_login.Create()
}

func CreateUsers() {
    born := time.Now()
    for i := 0 ; i < 20 ; i++ {
        u := church.User{}
        api.MapTo(map[string]interface{}{
            "email": fmt.Sprintf("jose%d@joao.com", i),
            "pass": "maria",
            "name": "jose",
            "genre": "M",
            "born": born.Format(church.TimeLayout()),
        }, &u)
        u.Create()

        photo := church.File{}
        photo.AltText = "Profile of "+ u.Name
        photo.LoadString("JOAO É GAY")
    }
}

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

package main

import (
    "encoding/json"
    "net/http"
    "bytes"
    "os"

    "../api"
    "github.com/Coff3e/Api"
)

func process(res *http.Response, err error) api.Response {
    if err == nil {
        r_body := api.Response{}
        r_raw_body := new(bytes.Buffer)
        r_raw_body.ReadFrom(res.Body)
        if json.Unmarshal(r_raw_body.Bytes(), &r_body) == nil {
            res.Body.Close()


            if  r_body.Type == "Error" {
                api.War("Sucessfull request with error\n\t-> Response:", r_raw_body)
            } else {
                api.Log("Sucessfull request\n\t-> Response:", r_raw_body)
            }
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
        "get", "/user/{id}", nil, church.GetUser,
    ).
    Add(
        "get", "/user/{id}/profile", nil, church.GetUserProfile,
    )

    go r.Run("/", 8000)

    uid := "1"

    //body := new(bytes.Buffer)
    url := "http://localhost:8000/user/" + uid
    res, err := http.Get(url)

    // Creating User
    r_body := process(res, err)
    if  r_body.Type != "Error" && r_body.Data != nil {
        //data := r_body.Data.(map[string]interface{})
        url += "/profile"
        res, err := http.Get(url)

        process(res, err)
    }
}

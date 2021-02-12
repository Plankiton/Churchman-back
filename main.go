package main

import (
    "os"
    "./api"
    "github.com/Coff3e/Api"
)

func main() {
    con_str := "host=localhost user=plankiton password=joaojoao dbname=church port=5432 sslmode=disable TimeZone=America/Araguaina"
    _, err := church.SignDB(con_str)
    if (err != nil) {
        os.Exit(1)
    }

    api.Log("Database connected with sucess")
}

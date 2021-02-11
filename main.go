package main

import (
    _ "./api"
    api "github.com/Coff3e/Api"
)

func main() {
    _, err := api.CreateDB("host=localhost user=plankiton password=joaojoao dbname=church port=5432 sslmode=disable TimeZone=America/Araguaina")

    if (err == nil) {
        print("DB loaded!!\n")
    }

    print("Hello World!!\n\n")
}

package church

import "github.com/Coff3e/Api"

type Address struct {
    api.Address
}

func GetFromCEP(r api.Request) (api.Response, int) {
    return api.Response{
        Type: "Sucess",
        Data: (&Address{}).FromPostalCode(r.PathVars["cep"]),
    }, 200
}

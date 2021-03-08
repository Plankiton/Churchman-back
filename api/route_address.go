package church

import "github.com/Coff3e/Api"

func GetFromCEP(r api.Request) (api.Response, int) {
    return api.Response {
        Type: "Sucess",
        Data: (&Address{}).FromPostalCode(r.PathVars["cep"]),
    }, 200
}

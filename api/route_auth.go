package church
import (
    "github.com/Coff3e/Api"
    "fmt"
)

func LogIn(r api.Request) (api.Response, int) {
    var data map[string]interface{}
    if api.ValidateData(r.Data, api.GenericJsonObj) {
        data = r.Data.(map[string]interface{})
    } else {
        msg := fmt.Sprint("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    user := User {}

    if db.First(&user, "email = ?", data["email"]).Error != nil {
        msg := fmt.Sprint("Não existe usuário com esse email")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }
    if _, ok := data["pass"];ok {
        if user.CheckPass(data["pass"].(string)) {
            token := Token {}
            token.UserId = user.ID
            token.Create()

            return api.Response {
                Type: "Sucess",
                Data: map[string]interface{} {
                    "token": token.ID,
                    "user": user,
                },
            }, 200
        }
    }

    msg := fmt.Sprint("Senha está errada")
    api.Err(msg)
    return api.Response {
        Message: msg,
        Type:    "Error",
    }, 405
}

func Verify(r api.Request) (api.Response, int) {
    token := Token {}
    token.ID = r.Token
    if u, ok := token.GetUser(); ok {
        return api.Response {
            Type: "Sucess",
            Data: u,
        }, 200
    }

    msg := fmt.Sprint("Fiel não existe")
    api.Err(msg)
    token.Delete()
    return api.Response {
        Message: msg,
        Type:    "Error",
    }, 404
}

func LogOut(r api.Request) (api.Response, int) {
    token := Token {}
    token.ID = r.Token

    db.First(&token)
    token.Delete()

    msg := fmt.Sprint("Dispositivo excluido com sucesso")
    api.Log(msg)
    return api.Response {
        Type: "Sucess",
        Message: msg,
    }, 200
}

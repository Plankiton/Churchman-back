package church
import (
    "github.com/Coff3e/Api"
    "time"
    "fmt"
)

func LogIn(r api.Request) (api.Response, int) {
    var data map[string]interface{}
    if api.ValidateData(r.Data, api.GenericJsonObj) {
        data = r.Data.(map[string]interface{})
    } else {
        msg := fmt.Sprint("Authentication fail, bad request, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    user := User {}

    if db.First(&user, "email = ?", data["email"]).Error != nil {
        msg := fmt.Sprint("Authentication fail, no users with \"", data["email"], "\" registered")
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

    msg := fmt.Sprint("Authentication fail, password from <", user.Name, ":", user.ID,"> is wrong, permission denied")
    api.Err(msg)
    return api.Response {
        Message: msg,
        Type:    "Error",
    }, 405
}

func Verify(r api.Request) (api.Response, int) {
    token := Token {}
    token.ID = r.Token
    if _, ok := token.GetUser(); ok {
        msg := fmt.Sprint("Token \"", r.Token, "\" Is valid")
        api.Log(msg)
        res := api.Response {
            Type: "Sucess",
            Message: msg,
        }

        res.AddCookie("token", r.Token, time.Hour*24*360*10)
        return res, 200
    }

    msg := fmt.Sprint("Authentication fail, user not found, permission denied")
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

    msg := fmt.Sprint("Token \"", r.Token, "\" removed")
    api.Log(msg)
    return api.Response {
        Type: "Sucess",
        Message: msg,
    }, 200
}

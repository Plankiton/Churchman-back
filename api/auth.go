package church

import (
    "github.com/Coff3e/Api"
    "reflect"
    "fmt"
)

type Token struct {
    api.Token
}

func (model *Token) Create() {
    model.ModelType = api.GetModelType(model)

    user := User{}
    user.ID = model.UserId

    e := db.First(&user)
    if e.Error == nil {
        var order int64
        db.Find(model).Count(&order)

        model.ID = api.ToHash(fmt.Sprintf(
            "%d;%d;%s;%s;%s", order, user.ID, user.Name, user.Email, user.Phone,
        ))

        db.Create(model)
        e = db.First(model)
        if e.Error == nil {
            ID := model.ID
            ModelType := model.ModelType
            api.Log("Created", api.ToLabel(ID, ModelType))
        }
    }
}

func Login(r api.Request) (api.Response, int) {
    var data map[string]interface{}
    switch reflect.TypeOf(r.Data).Kind() {
    case reflect.MapOf(reflect.TypeOf(generic_string), reflect.TypeOf(&generic_interface).Elem()).Kind():
        data = r.Data.(map[string]interface{})
    default:
        msg := fmt.Sprint("Authentication fail, bad request, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    user := User {}

    res := db.First(&user, "email = ?", data["email"])
    if res.Error != nil {
        msg := fmt.Sprint("Authentication fail, no users with \"", data["email"], "\" registered")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    if user.CheckPass(data["pass"].(string)) {
        token := Token {}
        token.UserId = user.ID
        token.Create()

        return api.Response {
            Type: "Sucess",
            Data: token,
        }, 200
    }

    msg := fmt.Sprint("Authentication fail, password from <", user.Name, ":", user.ID,"> is wrong, permission denied")
    api.Err(msg)
    return api.Response {
        Message: msg,
        Type:    "Error",
    }, 405
}

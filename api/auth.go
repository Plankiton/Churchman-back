package church

import (
    "github.com/Coff3e/Api"
    "fmt"
)

type Token struct {
    api.Token
}

func (token *Token) GetUser() (User, bool) {
    ok := false
    user := User{}
    if db.First(token, "id = ?", token.ID).Error == nil &&
    db.First(&user, "id = ?", token.UserId).Error == nil {
        ok = true
    }

    return user, ok
}

func (model *Token) Create() {
    model.ModelType = api.GetModelType(model)

    user := User{}
    user.ID = model.UserId

    if db.First(&user).Error == nil {
        var order int64
        db.Find(model).Count(&order)

        model.UserId = user.ID
        model.ID = api.ToHash(fmt.Sprintf(
            "%d;%d;%s;%s;%s", order, user.ID, user.Name, user.Email, user.Phone,
        ))

        if api.ModelCreate(model) == nil {
            ID := model.ID
            ModelType := model.ModelType
            api.Log("Created", api.ToLabel(ID, ModelType))
        }
    }
}


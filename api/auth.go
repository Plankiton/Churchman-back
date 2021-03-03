package church

import (
    "github.com/Coff3e/Api"
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


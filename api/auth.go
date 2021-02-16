package church

import (
    "github.com/Coff3e/Api"
)

type Token struct {
    api.Token
}

func (model *Token) Create() {
    model.ModelType = GetModelType(model)

    user := User{}
    user.ID = model.UserId

    e := _database.First(&user)
    if e.Error == nil {
        var order int64
        _database.Find(model).Count(&order)

        model.ID = api.ToHash(fmt.Sprintf(
            "%d;%d;%s;%s;%s", order, user.ID, user.Name, user.Email, user.Phone,
        ))

        _database.Create(model)
        e = _database.First(model)
        if e.Error == nil {
            ID := model.ID
            ModelType := model.ModelType
            Log("Created", ToLabel(ID, ModelType))
        }
    }
}



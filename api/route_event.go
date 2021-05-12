package church

import (
    "fmt"

    "time"
    "net/url"
    mp "mime/multipart"
    sc "strconv"

    "github.com/Coff3e/Api"
)

func GetEvent(r api.Request) (api.Response, int) {
    u := Event {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Evento não foi encontrado")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    return api.Response {
        Type: "Success",
        Data: u,
    }, 200
}

func RegectEventRequest(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprintf("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    u := EventPass {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Evento não foi encontrado")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    u.Confirmed = false
    api.ModelSave(u)

    return api.Response {
        Type: "Success",
        Message: "Usuario regeitado com sucesso",
    }, 200
}

func ApproveEventRequest(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprintf("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    u := EventPass {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Evento não foi encontrado")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    u.Confirmed = true
    api.ModelSave(u)

    return api.Response {
        Type: "Success",
        Message: "Usuario aprovado com sucesso",
    }, 200
}

func GetEventRequests(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprintf("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    var limit, page int

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    event_list := []EventPass{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Where("confirm <> true").Find(&event_list)

    if e.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Erro interno desconhecido",
        }, 500
    }

    custom_event_list := []map[string]interface{}{}
    for _, e := range event_list {
        event := map[string]interface{}{}
        event_name := ""
        user_name := ""

        db.Model(Event{}).Select("name").First(&event_name, "id = ?", e.EventId)
        db.Model(User{}).Select("name").First(&user_name, "id = ?", e.UserId)

        api.MapTo(e, &event)
        event["event"] = event_name
        event["user"] = user_name

        custom_event_list = append(custom_event_list, event)
    }

    return api.Response{
        Type: "Sucess",
        Data: custom_event_list,
    }, 200
}

func CreateEvent(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprintf("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    data := r.Data.(map[string]interface{})

    neededs := []string{
            "name", "begin", "end",
        }
    if (len(data)<len(neededs)){
        msg := "Campo"
        if (len(data)<len(neededs)-1) {
            msg += "s"
        }
        msg += " obrigatorio"
        if (len(data)<len(neededs)-1) {
            msg += "s"
        }
        msg += ": "
        for _, k := range neededs {
            if _, exist := data[k]; !exist {
                msg += fmt.Sprintf(`"%s", `, k)
            }
        }
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    event := Event {}

    api.MapTo(data, &event)

    begin_time, _ := time.Parse(TimeLayout(), data["begin"].(string))
    end_time, _ := time.Parse(TimeLayout(), data["begin"].(string))

    event.BeginAt = begin_time
    event.EndAt = end_time

    event.Create()

    return api.Response {
        Type: "Sucess",
        Data: event,
    }, 200
}

func UpdateEvent(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprintf("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprintf("Evento não foi encontrado")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    api.MapTo(data, &event)
    event.Save()

    return api.Response {
        Type: "Sucess",
        Data: event,
    }, 200
}

func DeleteEvent(r api.Request) (api.Response, int) {
    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprintf("Evento não foi encontrado")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    event.Delete()

    return api.Response {
        Type: "Sucess",
        Message: "Evento deleted",
    }, 200
}

func GetEventList(r api.Request) (api.Response, int) {
    var limit, page int

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    periodic, _ := sc.ParseBool(r.Conf["query"].(url.Values).Get("periodic"))

    event_list := []Event{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Where("periodic = ?", periodic).Find(&event_list)

    if e.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Erro interno desconhecido",
        }, 500
    }

    return api.Response{
        Type: "Sucess",
        Data: event_list,
    }, 200
}

func CreateEventCover(r api.Request) (api.Response, int) {
    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, event) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if r.Data == nil || !api.ValidateData(r.Data, api.GenericForm) {
        return api.Response{
            Type: "Error",
            Message: "Data must be a multipart-form",
        }, 400
    }

    data := r.Data.(*mp.Form)

    cover := File {}
    cover.Load(data)

    if db.First(&cover).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Erro interno desconhecido",
        }, 500
    }

    event.SetCover(cover)
    return api.Response {
        Type: "Sucess",
        Data: cover,
    }, 200
}

func GetEventCover(r api.Request) ([]byte, int) {
    u := Event {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Evento não foi encontrado")
        api.Err(msg)
        return []byte{}, 404
    }
    p := u.GetCover()

    return []byte(p.Render()), 200
}

func CreateEventAddr(r api.Request) (api.Response, int) {
    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, event) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprintf("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    addr := Address{}
    data := r.Data.(map[string]interface{})
    api.MapTo(data, &addr)
    addr.Create()

    event.SetAddress(addr)
    return api.Response {
        Type: "Sucess",
        Data: addr,
    }, 200
}

func GetEventAddr(r api.Request) (api.Response, int) {
    u := Event {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Evento não foi encontrado")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }
    addr := u.GetAddress()

    return api.Response {
        Type: "Sucess",
        Data: addr,
    }, 200
}

func EventUnsignUser(r api.Request) (api.Response, int) {
    user := User{}
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    event.Unsign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Unsigned to ", event.Name),
    }, 200
}

func EventSignUser(r api.Request) (api.Response, int) {
    user := User{}
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    event.Sign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Signed to ", event.Name),
    }, 200
}

func GetUserListByEvent(r api.Request) (api.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    var limit, page int

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    event := Event{}
    if (db.First(&event, "id = ?", r.PathVars["id"]).Error != nil) {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    user_list := event.GetUsers(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: user_list,
    }, 200
}


func GetEventListByUser(r api.Request) (api.Response, int) {
    var limit, page int

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    user := User{}
    if db.First(&user, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    event_list := user.GetEvents(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: event_list,
    }, 200
}


func EventUnsignCelule(r api.Request) (api.Response, int) {
    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, celule) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    event.Unsign(celule)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(celule.Name, " Unsigned to ", event.Name),
    }, 200
}

func EventSignCelule(r api.Request) (api.Response, int) {
    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, celule) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    event := Event{}
    if db.First(&event, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    event.Sign(celule)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(celule.Name, " Signed to ", event.Name),
    }, 200
}

func GetCeluleListByEvent(r api.Request) (api.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    var limit, page int

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    event := Event{}
    if (db.First(&event, "id = ?", r.PathVars["id"]).Error != nil) {
        return api.Response{
            Type: "Error",
            Message: "Evento não foi encontrado",
        }, 404
    }

    celule_list := event.GetCelules(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: celule_list,
    }, 200
}


func GetEventListByCelule(r api.Request) (api.Response, int) {
    limit, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ := sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    periodic, _ := sc.ParseBool(r.Conf["query"].(url.Values).Get("periodic"))

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    event_list := celule.GetEvents(page, limit, periodic)

    return api.Response{
        Type: "Sucess",
        Data: event_list,
    }, 200
}

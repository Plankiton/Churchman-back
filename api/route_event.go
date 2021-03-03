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
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Event not found")
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

func CreateEvent(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("Event create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    obrigatory_fields := []string{
            "name", "begin", "end",
        }
    if (len(data)<len(obrigatory_fields)){
        msg := "User create fail, Obrigatory field"
        if (len(data)==4) {
            msg += "s"
        }
        msg += " missing: "
        for _, k := range obrigatory_fields {
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
        msg := fmt.Sprint("Event create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    event := Event{}
    res := db.First(&event, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Event update fail, event not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
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
    res := db.First(&event, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Event delete fail, event not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    event.Delete()

    return api.Response {
        Type: "Sucess",
        Message: "Event deleted",
    }, 200
}

func EventUnsignUser(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["uid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    event := Event{}
    res = db.First(&event, "id = ?", r.PathVars["rid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Event not found",
        }, 404
    }

    user, event = event.Unsign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Unsigned to ", event.Name),
    }, 200
}

func EventSignUser(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["uid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    event := Event{}
    res = db.First(&event, "id = ?", r.PathVars["rid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Event not found",
        }, 404
    }

    user, event = event.Sign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Signed to ", event.Name),
    }, 200
}

func GetUserListByEvent(r api.Request) (api.Response, int) {
    var limit, page int
    var err error

    limit, err = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    if (err != nil) {
        return api.Response{
            Type: "Error",
            Message: "The query variable \"l\" is obrigatory and must be integer",
        }, 400
    }

    page, err = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    if (err != nil) {
        return api.Response{
            Type: "Error",
            Message: "The query variable \"p\" is obrigatory and must be integer",
        }, 400
    }

    event := Event{}
    if (db.First(&event, "id = ?", r.PathVars["id"]).Error != nil) {
        return api.Response{
            Type: "Error",
            Message: "Event not found",
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
    var err error

    limit, err = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    if (err != nil) {
        return api.Response{
            Type: "Error",
            Message: "The query variable \"l\" is obrigatory and must be integer",
        }, 400
    }

    page, err = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    if (err != nil) {
        return api.Response{
            Type: "Error",
            Message: "The query variable \"p\" is obrigatory and must be integer",
        }, 400
    }

    user := User{}
    if (db.First(&user, "id = ?", r.PathVars["id"]).Error != nil) {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    event_list := user.GetEvents(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: event_list,
    }, 200
}

func GetEventList(r api.Request) (api.Response, int) {
    var limit, page int
    var err error

    limit, err = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    if (err != nil) {
        return api.Response{
            Type: "Error",
            Message: "The query variable \"l\" is obrigatory and must be integer",
        }, 400
    }

    page, err = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))
    if (err != nil) {
        return api.Response{
            Type: "Error",
            Message: "The query variable \"p\" is obrigatory and must be integer",
        }, 400
    }

    event_list := []Event{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&event_list)

    if e.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    return api.Response{
        Type: "Sucess",
        Data: event_list,
    }, 200
}

func CreateEventCover(r api.Request) (api.Response, int) {
    user := Event{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Event not found",
        }, 404
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
            Message: "Error on creating of cover on database",
        }, 500
    }

    user.SetCover(cover)
    return api.Response {
        Type: "Sucess",
        Data: cover,
    }, 200
}

func GetEventCover(r api.Request) ([]byte, int) {
    u := Event {}
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Event not found")
        api.Err(msg)
        return []byte{}, 404
    }
    p := u.GetCover()

    return []byte(p.Render()), 200
}

func CreateEventAddr(r api.Request) (api.Response, int) {
    user := Event{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Event not found",
        }, 404
    }

    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("Event create fail, data need to be a object")
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

    user.SetAddress(addr)
    return api.Response {
        Type: "Sucess",
        Data: addr,
    }, 200
}

func GetEventAddr(r api.Request) (api.Response, int) {
    u := Event {}
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Event not found")
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

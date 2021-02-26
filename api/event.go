package church

import (
    "fmt"

    "time"
    "net/url"
    sc "strconv"

    "github.com/Coff3e/Api"
)

type UserEvent struct {
    api.UserGroup
}

func (model *UserEvent) TableName() string {
    return "user_groups"
}

type Event struct {
    api.Group
    Type       string       `json:"type,omitempty" gorm:"index"`

    AddrId     uint         `json:",empty" gorm:"index"`
    CoverId    uint         `json:",empty"`

    BeginAt    time.Time    `json:"begin,omitempty" gorm:"index"`
    EndAt      time.Time    `json:"end,omitempty" gorm:"index"`
}

func (model *Event) TableName() string {
    return "groups"
}

func (model *Event) Create() {
    model.ModelType = "Event"

    db.Create(model)

    e := db.First(model)
    if e.Error == nil {

        ID := model.ID
        ModelType := model.ModelType
        api.Log("Created", api.ToLabel(ID, ModelType))
    }
}

func (model *UserEvent) Create() {
    model.ModelType = "UserEvent"

    db.Create(model)

    e := db.First(model)
    if e.Error == nil {

        ID := model.ID
        ModelType := model.ModelType
        api.Log("Created", api.ToLabel(ID, ModelType))
    }
}

func (self UserEvent) Sign(user User, event Event) (User, Event) {
    self.UserId = user.ID
    self.GroupId = event.ID

    self.Create()
    api.Log("Linked", api.ToLabel(user.ID, user.ModelType), user.Name, "to", api.ToLabel(event.ID, event.ModelType), event.Name)

    return user, event
}

func (self UserEvent) Unsign(user User, event Event) (User, Event) {
    self.UserId = user.ID
    self.GroupId = event.ID

    self.Delete()
    api.Log("Unlinked", api.ToLabel(user.ID, user.ModelType), user.Name, "from", api.ToLabel(event.ID, event.ModelType), event.Name)

    return user, event
}

func (self Event) Sign(user User) (User, Event) {
    link := UserEvent{}
    user, self = link.Sign(user, self)

    return user, self
}

func (self Event) Unsign(user User) (User, Event) {
    link := UserEvent{}
    e := db.Where("user_id = ? AND group_id = ?", user.ID, self.ID).First(&link)
    if e.Error == nil {
        user, self = link.Unsign(user, self)
    }

    return user, self
}

func (self *Event) SetCover(cover File) {
    {
        tmp_cover := File{}
        e := db.First(&tmp_cover, "id = ?", self.CoverId)
        if e.Error == nil {
            tmp_cover.Delete()
        }
    }

    self.CoverId = cover.ID
    self.Save()
}

func (self *Event) GetCover() File {
    cover := File{}
    cover.ID = self.CoverId
    e := db.First(&cover)
    if e.Error == nil {
        return cover
    }

    return File{}
}

func (self *Event) SetAddress(addr File) {
    {
        tmp_addr := File{}
        e := db.First(&tmp_addr, "id = ?", self.AddrId)
        if e.Error == nil {
            tmp_addr.Delete()
        }
    }

    self.AddrId = addr.ID
    self.Save()
}

func (self *Event) GetAddress() File {
    addr := File{}
    addr.ID = self.AddrId
    e := db.First(&addr)
    if e.Error == nil {
        return addr
    }

    return File{}
}

func (self *Event) GetUsers(page int, limit int) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_groups ur INNER JOIN groups r ON ur.group_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID).Offset((page-1)*limit).Limit(limit).Find(&user_list)
        if e.Error == nil {
            e := db.Find(&users, "id in ?", user_list)
            if e.Error == nil {
                return users
            }
        }
    }

    return []User{}
}

func (self *User) GetEvents(page int, limit int) []Event {
    e := db.First(self)
    if e.Error == nil {
        event_list := []uint{}
        events := []Event{}
        e := db.Raw("SELECT r.id FROM groups r INNER JOIN user_groups ur INNER JOIN users u ON ur.group_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID).Offset((page-1)*limit).Limit(limit).Find(&event_list)
        if e.Error == nil {
            e := db.Find(&events, "id in ?", event_list)
            if e.Error == nil {
                return events
            }
        }
    }

    return []Event{}
}

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
    if !validData(r.Data, generic_json_obj) {
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
    if !validData(r.Data, generic_json_obj) {
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

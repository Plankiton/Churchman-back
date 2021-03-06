package church

import (
    "github.com/Coff3e/Api"
)

type UserEvent struct {
    api.UserEvent
}

type CeluleEvent struct {
    api.Model

    CeluleId   uint
    EventId    uint
}

func (model *CeluleEvent) TableName() string {
    return "group_events"
}

type Event struct {
    api.Event

    Name        string `json:"name,omitempty"`
    Periodic    bool   `json:"periodic,omitempty" gorm:"default:false"`
    WeeklyDay   uint   `json:"weekly_day,omitempty"`

    NeedPass    bool   `json:"need_pass,omitempty"`
}

type EventPass struct {
    api.Model
    EventId     uint    `json:"-"`
    UserId      uint    `json:"-"`
    Confirmed   bool    `json:"confirmed,omitempty"`
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
    self.EventId = event.ID

    if event.NeedPass {
        event_pass := EventPass{}
        event_pass.UserId = user.ID
        event_pass.EventId = event.ID

        api.ModelCreate(event_pass)
    }

    self.Create()
    api.Log("Linked", api.ToLabel(user.ID, user.ModelType), user.Name, "to", api.ToLabel(event.ID, event.ModelType), event.Name)

    return user, event
}

func (self UserEvent) Unsign(user User, event Event) (User, Event) {
    self.UserId = user.ID
    self.EventId = event.ID

    if event.NeedPass {
        event_pass := EventPass{}
        if db.First(&event_pass, "user_id = ? AND event_id = ?", user.ID, event.ID).Error == nil {
            api.ModelDelete(event_pass)
        }
    }

    self.Delete()
    api.Log("Unlinked", api.ToLabel(user.ID, user.ModelType), user.Name, "from", api.ToLabel(event.ID, event.ModelType), event.Name)

    return user, event
}

func (self Event) Sign(obj interface{}) (interface{}, Event) {
    if user, ok := obj.(User);ok {
        link := UserEvent{}
        user, self = link.Sign(user, self)
        return user, self
    } else if celule, ok := obj.(Celule);ok {
        link := CeluleEvent{}
        celule, self = link.Sign(celule, self)
        return celule, self
    }

    return nil, self
}

func (self Event) Unsign(obj interface{}) (interface{}, Event) {
    if user, ok := obj.(User);ok {
        link := UserEvent{}
        user, self = link.Unsign(user, self)
        return user, self
    } else if celule, ok := obj.(Celule);ok {
        link := CeluleEvent{}
        celule, self = link.Unsign(celule, self)
        return celule, self
    }

    return nil, self
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

func (self *Event) SetAddress(addr Address) {
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

func (self *Event) GetAddress() Address {
    addr := Address{}
    addr.ID = self.AddrId
    e := db.First(&addr)
    if e.Error == nil {
        return addr
    }

    return Address{}
}

func (self *Event) GetUsers(page int, limit int) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_events ur INNER JOIN events r ON ur.event_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&user_list)

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
        e := db.Raw("SELECT r.id FROM events r INNER JOIN user_events ur INNER JOIN users u ON ur.event_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&event_list)

        if e.Error == nil {
            e := db.Find(&events, "id in ?", event_list)
            if e.Error == nil {
                return events
            }
        }
    }

    return []Event{}
}

func (self *Event) QueryUsers(page int, limit int, query ...interface{}) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_events ur INNER JOIN events r ON ur.event_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }

        e = e.Find(&user_list, query...)

        if e.Error == nil {
            e := db.Find(&users, "id in ?", user_list)
            if e.Error == nil {
                return users
            }
        }
    }

    return []User{}
}


func (self *User) QueryEvents(page int, limit int, query...interface{}) []Event {
    e := db.First(self)
    if e.Error == nil {
        event_list := []uint{}
        events := []Event{}
        e := db.Raw("SELECT r.id FROM events r INNER JOIN user_events ur INNER JOIN users u ON ur.event_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&event_list, query...)

        if e.Error == nil {
            e := db.Find(&events, "id in ?", event_list)
            if e.Error == nil {
                return events
            }
        }
    }

    return []Event{}
}

func (model *CeluleEvent) Create() {
    if api.ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        api.Log("Created", api.ToLabel(ID, ModelType))
    }
}

func (model *CeluleEvent) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    if api.ModelDelete(model) == nil {
        api.Log("Deleted", api.ToLabel(ID, ModelType))
    }
}

func (model *CeluleEvent) Save() {
    ID := model.ID
    ModelType := model.ModelType

    if api.ModelSave(model) == nil {
        api.Log("Updated", api.ToLabel(ID, ModelType))
    }
}

func (model *CeluleEvent) Update(columns api.Dict) {
    ID := model.ID
    ModelType := model.ModelType

    if api.ModelUpdate(model, columns) == nil {
        api.Log("Updated", api.ToLabel(ID, ModelType))
    }
}

func (self CeluleEvent) Sign(celule Celule, event Event) (Celule, Event) {
    self.CeluleId = celule.ID
    self.EventId = event.ID

    self.Create()
    api.Log("Linked", api.ToLabel(celule.ID, celule.ModelType), celule.Name, "to", api.ToLabel(event.ID, event.ModelType), event.Name)

    return celule, event
}

func (self CeluleEvent) Unsign(celule Celule, event Event) (Celule, Event) {
    self.CeluleId = celule.ID
    self.EventId = event.ID

    self.Delete()
    api.Log("Unlinked", api.ToLabel(celule.ID, celule.ModelType), celule.Name, "from", api.ToLabel(event.ID, event.ModelType), event.Name)

    return celule, event
}

func (self *Event) GetCelules(page int, limit int) []Celule {
    e := db.First(self)
    if e.Error == nil {
        celule_list := []uint{}
        celules := []Celule{}
        e := db.Raw("SELECT u.id FROM groups u INNER JOIN group_events ur INNER JOIN events r ON ur.event_id = r.id AND ur.celule_id = u.id AND r.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&celule_list)

        if e.Error == nil {
            e := db.Find(&celules, "id in ?", celule_list)
            if e.Error == nil {
                return celules
            }
        }
    }

    return []Celule{}
}

func (self *Celule) GetEvents(page int, limit int, periodic bool) []Event {
    e := db.First(self)
    if e.Error == nil {
        event_list := []uint{}
        events := []Event{}
        e := db.Raw("SELECT r.id FROM events r INNER JOIN group_events ur INNER JOIN groups u ON ur.event_id = r.id AND ur.celule_id = u.id AND u.id = ? AND r.periodic = ?", self.ID, periodic)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&event_list)

        if e.Error == nil {
            e := db.Find(&events, "id in ?", event_list)
            if e.Error == nil {
                return events
            }
        }
    }

    return []Event{}
}

func (self *Event) QueryCelules(page int, limit int, query ...interface{}) []Celule {
    e := db.First(self)
    if e.Error == nil {
        celule_list := []uint{}
        celules := []Celule{}
        e := db.Raw("SELECT u.id FROM groups u INNER JOIN group_events ur INNER JOIN events r ON ur.event_id = r.id AND ur.celule_id = u.id AND r.id = ? AND r.periodic = false", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }

        e = e.Find(&celule_list, query...)

        if e.Error == nil {
            e := db.Find(&celules, "id in ?", celule_list)
            if e.Error == nil {
                return celules
            }
        }
    }

    return []Celule{}
}


func (self *Celule) QueryEvents(page int, limit int, query...interface{}) []Event {
    e := db.First(self)
    if e.Error == nil {
        event_list := []uint{}
        events := []Event{}
        e := db.Raw("SELECT r.id FROM events r INNER JOIN group_events ur INNER JOIN groups u ON ur.event_id = r.id AND ur.celule_id = u.id AND u.id = ?", self.ID)
        if limit > 0 && page > 0 {
            e = e.Offset((page-1)*limit).Limit(limit)
        }
        e = e.Find(&event_list, query...)

        if e.Error == nil {
            e := db.Find(&events, "id in ?", event_list)
            if e.Error == nil {
                return events
            }
        }
    }

    return []Event{}
}


package church

import (
    "github.com/Coff3e/Api"
)

type UserEvent struct {
    api.UserEvent
}

type Event struct {
    api.Event
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

    self.Create()
    api.Log("Linked", api.ToLabel(user.ID, user.ModelType), user.Name, "to", api.ToLabel(event.ID, event.ModelType), event.Name)

    return user, event
}

func (self UserEvent) Unsign(user User, event Event) (User, Event) {
    self.UserId = user.ID
    self.EventId = event.ID

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
    e := db.Where("user_id = ? AND event_id = ?", user.ID, self.ID).First(&link)
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


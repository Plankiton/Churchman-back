package church

import (
    "github.com/Coff3e/Api"
    "time"
)

type UserEvent struct {
    api.UserGroup
}

type Event struct {
    api.Group
    Type       string       `json:"type,omitempty" gorm:"index"`

    AddrId     uint         `json:",empty" gorm:"index"`
    Leader     uint         `json:",empty" gorm:"index"`
    CoLeader   uint         `json:",empty" gorm:"index"`
    CoverId    uint         `json:",empty"`

    BeginAt    time.Time    `json:"begin,omitempty" gorm:"index"`
    EndAt      time.Time    `json:"end,omitempty" gorm:"index"`
}

func (self UserEvent) TableName() string {
    return "user_groups"
}
func (self Event) TableName() string {
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

func (self *Event) GetUsers() []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_groups ur INNER JOIN groups r ON ur.group_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID).Find(&user_list)
        if e.Error == nil {
            e := db.Find(&users, "id in ?", user_list)
            if e.Error == nil {
                return users
            }
        }
    }

    return []User{}
}

func (self *User) GetEvents() []Event {
    e := db.First(self)
    if e.Error == nil {
        event_list := []uint{}
        events := []Event{}
        e := db.Raw("SELECT r.id FROM groups r INNER JOIN user_groups ur INNER JOIN users u ON ur.group_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID).Find(&event_list)
        if e.Error == nil {
            e := db.Find(&events, "id in ?", event_list)
            if e.Error == nil {
                return events
            }
        }
    }

    return []Event{}
}

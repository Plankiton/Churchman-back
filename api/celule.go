package church

import (
    "github.com/Coff3e/Api"
)

type UserCelule struct {
    api.UserGroup
}

type Celule struct {
    api.Group
    Type       string `json:"type,omitempty" gorm:"index"`
    LocalType  string `json:"local_type,omitempty"`
    IID        uint   `json:"internal_id,omitempty" gorm:"index,column:internal_id"`
    Parent     uint   `json:"parent_id,omitempty" gorm:"index"`
    Addr       uint   `json:"address_id,omitempty" gorm:"index"`
    Leader     uint   `json:"leader_id,omitempty" gorm:"index"`
    CoLeader   uint   `json:"co_leader_id,omitempty" gorm:"index,column:co_leader"`
    CoverId    uint   `json:"-"`
}

func (self UserCelule) TableName() string {
    return "user_groups"
}
func (self Celule) TableName() string {
    return "groups"
}

func (model *Celule) Create() {
    model.ModelType = "Celule"
    last := Celule{}
    db.Order("created_at desc").Last(&last, "parent = ?", model.Parent)
    model.IID = last.IID+1

    if api.ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType

        model.Name = GetCeluleName(*model)
        api.ModelSave(model)

        api.Log("Created", api.ToLabel(ID, ModelType))
    }
}

func (model *UserCelule) Create() {
    model.ModelType = "UserCelule"

    if api.ModelCreate(model) == nil {

        ID := model.ID
        ModelType := model.ModelType
        api.Log("Created", api.ToLabel(ID, ModelType))
    }
}


func (self UserCelule) Sign(user User, celule Celule) (User, Celule) {
    self.UserId = user.ID
    self.GroupId = celule.ID

    self.Create()
    api.Log("Linked", api.ToLabel(user.ID, user.ModelType), user.Name, "to", api.ToLabel(celule.ID, celule.ModelType), celule.Name)

    return user, celule
}

func (self UserCelule) Unsign(user User, celule Celule) (User, Celule) {
    self.UserId = user.ID
    self.GroupId = celule.ID

    self.Delete()
    api.Log("Unlinked", api.ToLabel(user.ID, user.ModelType), user.Name, "from", api.ToLabel(celule.ID, celule.ModelType), celule.Name)

    return user, celule
}

func (self Celule) Sign(user User) (User, Celule) {
    link := UserCelule{}
    user, self = link.Sign(user, self)

    return user, self
}

func (self Celule) Unsign(user User) (User, Celule) {
    link := UserCelule{}
    e := db.Where("user_id = ? AND group_id = ?", user.ID, self.ID).First(&link)
    if e.Error == nil {
        user, self = link.Unsign(user, self)
    }

    return user, self
}

func (self *Celule) SetAddress(addr Address) {
    {
        tmp_addr := Address{}
        e := db.First(&tmp_addr, "id = ?", self.Addr)
        if e.Error == nil {
            tmp_addr.Delete()
        }
    }

    self.Addr = addr.ID
    self.Save()
}

func (self *Celule) GetAddress() Address {
    addr := Address{}
    addr.ID = self.Addr
    e := db.First(&addr)
    if e.Error == nil {
        return addr
    }

    return Address{}
}

func (self *Celule) SetParent(user User) {
    {
        tmp_user := User{}
        e := db.First(&tmp_user, "id = ?", self.Parent)
        if e.Error == nil {
            tmp_user.Delete()
        }
    }

    self.Parent = user.ID
    self.Save()
}

func (self *Celule) GetParent() User {
    user := User{}
    user.ID = self.Parent
    e := db.First(&user)
    if e.Error == nil {
        return user
    }

    return User{}
}

func (self *Celule) GetUsers(page int, limit int) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_groups ur INNER JOIN groups r ON ur.group_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID)
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

func (self *User) GetCelules(page int, limit int, church ...bool) []Celule {
    query_str := "SELECT r.id FROM groups r INNER JOIN user_groups ur INNER JOIN users u ON ur.group_id = r.id AND ur.user_id = u.id AND u.id = ?"
    if church != nil {
        if church[0] {
            query_str += "AND r.type = 'church'"
        } else {
            query_str += "AND r.type <> 'church'"
        }
    }

    e := db.First(self)
    if e.Error == nil {
        celule_list := []uint{}
        celules := []Celule{}
        e := db.Raw(query_str, self.ID)
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

func (self *Celule) QueryUsers(page int, limit int, query ...interface{}) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_groups ur INNER JOIN groups r ON ur.group_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID)
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


func (self *User) QueryCelules(page int, limit int, query...interface{}) []Celule {
    e := db.First(self)
    if e.Error == nil {
        celule_list := []uint{}
        celules := []Celule{}
        e := db.Raw("SELECT r.id FROM groups r INNER JOIN user_groups ur INNER JOIN users u ON ur.celule_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID)
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

func (self *Celule) SetCover(cover File) {
    {
        tmp_cover := File{}
        e := db.First(&tmp_cover, "id = ?", self.CoverId)
        if e.Error == nil {
            tmp_cover.Delete()
        }

        e = db.First(&tmp_cover, "filename = ? OR id = ?", cover.Filename, cover.ID)
        if e.Error != nil {
            tmp_cover.Create()
        }
    }

    self.Update(api.Dict { "cover_id": cover.ID })
}

func (self *Celule) GetCover() File {
    cover := File{}
    cover.ID = self.CoverId
    e := db.First(&cover)
    if e.Error == nil {
        return cover
    }

    return File{}
}


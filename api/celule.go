package church

import (
    "fmt"

    "net/url"
    sc "strconv"

    "github.com/Coff3e/Api"
)

type UserCelule struct {
    api.UserGroup
}

type Celule struct {
    api.Group
    Type       string `json:"type,omitempty" gorm:"index"`
    AddrId     uint   `json:",empty" gorm:"index"`
    Leader     uint   `json:",empty" gorm:"index"`
    CoLeader   uint   `json:",empty" gorm:"index"`
}

func (self UserCelule) TableName() string {
    return "user_groups"
}
func (self Celule) TableName() string {
    return "groups"
}

func (model *Celule) Create() {
    model.ModelType = "Celule"

    db.Create(model)

    e := db.First(model)
    if e.Error == nil {

        ID := model.ID
        ModelType := model.ModelType
        api.Log("Created", api.ToLabel(ID, ModelType))
    }
}

func (model *UserCelule) Create() {
    model.ModelType = "UserCelule"

    db.Create(model)

    e := db.First(model)
    if e.Error == nil {

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
        e := db.First(&tmp_addr, "id = ?", self.AddrId)
        if e.Error == nil {
            tmp_addr.Delete()
        }
    }

    self.AddrId = addr.ID
    self.Save()
}

func (self *Celule) GetAddress() Address {
    addr := Address{}
    addr.ID = self.AddrId
    e := db.First(&addr)
    if e.Error == nil {
        return addr
    }

    return Address{}
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

func (self *User) GetCelules(page int, limit int) []Celule {
    e := db.First(self)
    if e.Error == nil {
        celule_list := []uint{}
        celules := []Celule{}
        e := db.Raw("SELECT r.id FROM groups r INNER JOIN user_groups ur INNER JOIN users u ON ur.group_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID)
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


func GetCelule(r api.Request) (api.Response, int) {
    u := Celule {}
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Celule not found")
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

func CreateCelule(r api.Request) (api.Response, int) {
    if !validData(r.Data, generic_json_obj) {
        msg := fmt.Sprint("Celule create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    if _, e := data["name"]; !e {
        msg := "Celule create fail, Obrigatory field \"name\""
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    res := db.First(&Celule {}, "name = ?", data["name"])
    if res.Error == nil {
        msg := fmt.Sprint("Celule create fail, celule already registered")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    celule := Celule {}

    api.MapTo(data, &celule)
    celule.Create()

    return api.Response {
        Type: "Sucess",
        Data: celule,
    }, 200
}

func UpdateCelule(r api.Request) (api.Response, int) {
    if !validData(r.Data, generic_json_obj) {
        msg := fmt.Sprint("Celule create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    celule := Celule{}
    res := db.First(&celule, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Celule update fail, celule not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    api.MapTo(data, &celule)
    celule.Save()

    return api.Response {
        Type: "Sucess",
        Data: celule,
    }, 200
}

func DeleteCelule(r api.Request) (api.Response, int) {
    celule := Celule{}
    res := db.First(&celule, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Celule delete fail, celule not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    celule.Delete()

    return api.Response {
        Type: "Sucess",
        Message: "Celule deleted",
    }, 200
}

func CeluleUnsignUser(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["uid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    celule := Celule{}
    res = db.First(&celule, "id = ?", r.PathVars["rid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    user, celule = celule.Unsign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Unsigned to ", celule.Name),
    }, 200
}

func CeluleSignUser(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["uid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    celule := Celule{}
    res = db.First(&celule, "id = ?", r.PathVars["rid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    user, celule = celule.Sign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Signed to ", celule.Name),
    }, 200
}

func GetUserListByCelule(r api.Request) (api.Response, int) {
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

    celule := Celule{}
    if (db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil) {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    user_list := celule.GetUsers(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: user_list,
    }, 200
}


func GetCeluleListByUser(r api.Request) (api.Response, int) {
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

    celule_list := user.GetCelules(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: celule_list,
    }, 200
}

func GetCeluleList(r api.Request) (api.Response, int) {
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

    celule_list := []Celule{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&celule_list)

    if e.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    return api.Response{
        Type: "Sucess",
        Data: celule_list,
    }, 200
}

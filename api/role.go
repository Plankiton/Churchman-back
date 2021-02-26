package church

import (
    "fmt"

    "net/url"
    sc "strconv"

    "github.com/Coff3e/Api"
)

type Role struct {
    api.Role
}

type UserRole struct {
    api.UserRole
}

func (self UserRole) Sign(user User, role Role) (User, Role) {
    self.UserId = user.ID
    self.RoleId = role.ID

    self.Create()
    api.Log("Linked", api.ToLabel(user.ID, user.ModelType), user.Name, "to", api.ToLabel(role.ID, role.ModelType), role.Name)

    return user, role
}

func (self UserRole) Unsign(user User, role Role) (User, Role) {
    self.UserId = user.ID
    self.RoleId = role.ID

    self.Delete()
    api.Log("Unlinked", api.ToLabel(user.ID, user.ModelType), user.Name, "from", api.ToLabel(role.ID, role.ModelType), role.Name)

    return user, role
}

func (self Role) Sign(user User) (User, Role) {
    link := UserRole{}
    user, self = link.Sign(user, self)

    return user, self
}

func (self Role) Unsign(user User) (User, Role) {
    link := UserRole{}
    e := db.Where("user_id = ? AND role_id = ?", user.ID, self.ID).First(&link)
    if e.Error == nil {
        user, self = link.Unsign(user, self)
    }

    return user, self
}

func (self *Role) GetUsers(page int, limit int) []User {
    e := db.First(self)
    if e.Error == nil {
        user_list := []uint{}
        users := []User{}
        e := db.Raw("SELECT u.id FROM users u INNER JOIN user_roles ur INNER JOIN roles r ON ur.role_id = r.id AND ur.user_id = u.id AND r.id = ?", self.ID).Offset((page-1)*limit).Limit(limit).Find(&user_list)
        if e.Error == nil {
            e := db.Find(&users, "id in ?", user_list)
            if e.Error == nil {
                return users
            }
        }
    }

    return []User{}
}

func (self *User) GetRoles(page int, limit int) []Role {
    e := db.First(self)
    if e.Error == nil {
        role_list := []uint{}
        roles := []Role{}
        e := db.Raw("SELECT r.id FROM roles r INNER JOIN user_roles ur INNER JOIN users u ON ur.role_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID).Offset((page-1)*limit).Limit(limit).Find(&role_list)
        if e.Error == nil {
            e := db.Find(&roles, "id in ?", role_list)
            if e.Error == nil {
                return roles
            }
        }
    }

    return []Role{}
}

func GetRole(r api.Request) (api.Response, int) {
    u := Role {}
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Role not found")
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

func CreateRole(r api.Request) (api.Response, int) {
    if !validData(r.Data, generic_json_obj) {
        msg := fmt.Sprint("Role create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    if _, e := data["name"]; !e {
        msg := "Role create fail, Obrigatory field \"name\""
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    res := db.First(&Role {}, "name = ?", data["name"])
    if res.Error == nil {
        msg := fmt.Sprint("Role create fail, role already registered")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    role := Role {}

    api.MapTo(data, &role)
    role.Create()

    return api.Response {
        Type: "Sucess",
        Data: role,
    }, 200
}

func UpdateRole(r api.Request) (api.Response, int) {
    if !validData(r.Data, generic_json_obj) {
        msg := fmt.Sprint("Role create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    role := Role{}
    res := db.First(&role, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Role update fail, role not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    api.MapTo(data, &role)
    role.Save()

    return api.Response {
        Type: "Sucess",
        Data: role,
    }, 200
}

func DeleteRole(r api.Request) (api.Response, int) {
    role := Role{}
    res := db.First(&role, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Role delete fail, role not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    role.Delete()

    return api.Response {
        Type: "Sucess",
        Message: "Role deleted",
    }, 200
}

func RoleUnsignUser(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    if !validData(r.Data, generic_json_obj) {
        msg := fmt.Sprint("Role create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    role := Role{}
    res = db.First(&role, "id = ?", data["id"])

    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    user, role = role.Unsign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Unsigned to ", role.Name),
    }, 200
}

func RoleSignUser(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    if !validData(r.Data, generic_json_obj) {
        msg := fmt.Sprint("Role create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    role := Role{}
    res = db.First(&role, "id = ?", data["id"])

    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    user, role = role.Sign(user)
    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(user.Name, " Signed to ", role.Name),
    }, 200
}

func GetUserListByRole(r api.Request) (api.Response, int) {
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

    role := Role{}
    if (db.First(&role, "id = ?", r.PathVars["id"]) != nil) {
        return api.Response{
            Type: "Error",
            Message: "Role not found",
        }, 404
    }

    user_list := role.GetUsers(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: user_list,
    }, 200
}


func GetRoleListByUser(r api.Request) (api.Response, int) {
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
    if (db.First(&user, "id = ?", r.PathVars["id"]) != nil) {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    role_list := user.GetRoles(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: role_list,
    }, 200
}

func GetRoleList(r api.Request) (api.Response, int) {
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

    role_list := []Role{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&role_list)

    if e.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    return api.Response{
        Type: "Sucess",
        Data: role_list,
    }, 200
}

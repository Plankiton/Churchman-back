package church

import (
    "fmt"

    "net/url"
    sc "strconv"
    "github.com/Coff3e/Api"
)

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
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
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
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
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
    res := db.First(&user, "id = ?", r.PathVars["uid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    role := Role{}
    res = db.First(&role, "id = ?", r.PathVars["rid"])
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
    res := db.First(&user, "id = ?", r.PathVars["uid"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    role := Role{}
    res = db.First(&role, "id = ?", r.PathVars["rid"])
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

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    role := Role{}
    if (db.First(&role, "id = ?", r.PathVars["id"]).Error != nil) {
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

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    user := User{}
    if (db.First(&user, "id = ?", r.PathVars["id"]).Error != nil) {
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

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

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

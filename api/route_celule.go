package church

import (
	"fmt"

	"net/url"
	sc "strconv"

	"github.com/Coff3e/Api"
)

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
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("Celule create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})
    required := []string{
        "leader_id",
    }

    if (len(data)<len(required)){
        msg := "User create fail, Obrigatory field"
        if (len(data)==len(required)-1) {
            msg += "s"
        }
        msg += " missing: "
        for _, k := range required {
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

    celule := Celule {}
    api.MapTo(data, &celule)
    celule.Create()
    return api.Response {
        Type: "Sucess",
        Data: celule,
    }, 200
}

func UpdateCelule(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
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

    if _,e := data["name"];e {
        data["name"] = nil
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

func CeluleSetCoLeader(r api.Request) (api.Response, int) {
    co_leader := User{}
    if db.First(&co_leader, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    celule.CoLeader = co_leader.ID
    celule.Save()

    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(co_leader.Name, " Signed to ", celule.Name),
    }, 200
}

func CeluleGetCoLeader(r api.Request) (api.Response, int) {
    co_leader := User{}
    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    if db.First(&co_leader, "id = ?", celule.CoLeader).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "CoLeader not found",
        }, 404
    }
    return api.Response {
        Type: "Sucess",
        Data: co_leader,
    }, 200
}

func CeluleSetLeader(r api.Request) (api.Response, int) {
    leader := User{}
    if db.First(&leader, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    celule.Leader = leader.ID
    celule.Save()

    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(leader.Name, " Signed to ", celule.Name),
    }, 200
}

func CeluleGetLeader(r api.Request) (api.Response, int) {
    leader := User{}
    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    if db.First(&leader, "id = ?", celule.Leader).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Leader not found",
        }, 404
    }
    return api.Response {
        Type: "Sucess",
        Data: leader,
    }, 200
}

func CeluleSetParent(r api.Request) (api.Response, int) {
    parent := Celule{}
    if db.First(&parent, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    celule.Parent = parent.ID
    celule.Save()

    return api.Response {
        Type: "Sucess",
        Message: fmt.Sprint(parent.Name, " Signed to ", celule.Name),
    }, 200
}

func CeluleGetParent(r api.Request) (api.Response, int) {
    parent := Celule{}
    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    if db.First(&parent, "id = ?", celule.Parent).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Parent not found",
        }, 404
    }
    return api.Response {
        Type: "Sucess",
        Data: parent,
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

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

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

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

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

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

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

func CreateCeluleAddr(r api.Request) (api.Response, int) {
    user := Celule{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celule not found",
        }, 404
    }

    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("Celule create fail, data need to be a object")
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

func GetCeluleAddr(r api.Request) (api.Response, int) {
    u := Celule {}
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("Celule not found")
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

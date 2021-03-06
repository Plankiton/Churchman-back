package church

import (
	"fmt"
  str "strings"

	"net/url"
	sc "strconv"
  mp "mime/multipart"

	"github.com/Coff3e/Api"
)

func GetCelule(r api.Request) (api.Response, int) {
    u := Celule {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Celula não encontrada")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }
    church := false
    if u.Type == "church" {
        church = true
    }

    token := Token {}
    token.ID = r.Token
    if !church {
        if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, u) {
            msg := "Você não tem permissão para acessar isso"
            api.Err(msg)
            return api.Response {
                Message: msg,
                Type:    "Error",
            }, 405
        }
    }

    return api.Response {
        Type: "Success",
        Data: u,
    }, 200
}

func CreateCelule(r api.Request) (api.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})
    neededs := []string{
        "leader_id",
    }

    if (len(data)<len(neededs)){
        msg := "Campo"
        if (len(data)<len(neededs)-1) {
            msg += "s"
        }
        msg += " obrigatorio"
        if (len(data)<len(neededs)-1) {
            msg += "s"
        }
        msg += ": "

        for _, k := range neededs {
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

    alt_name, _ := sc.ParseBool(r.Conf["query"].(url.Values).Get("alt_id"))
    if alt_name {
        celule.Name = str.Replace(celule.Name, "m", "h", -1)
        celule.Name = str.Replace(celule.Name, "f", "m", -1)
    }

    celule.Save()

    return api.Response {
        Type: "Sucess",
        Data: celule,
    }, 200
}

func UpdateCelule(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("Dados enviados são inválidos")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Celula não encontrada")
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
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Celula não encontrada")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    celule.Delete()

    return api.Response {
        Type: "Sucess",
        Message: "Celula deleted",
    }, 200
}

func CeluleSetCoLeader(r api.Request) (api.Response, int) {
    co_leader := User{}
    if db.First(&co_leader, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Usuário não encontrado",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celula não encontrada",
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
            Message: "Celula não encontrada",
        }, 404
    }

    if db.First(&co_leader, "id = ?", celule.CoLeader).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Timóteo não encontrado",
        }, 404
    }

    celule.CoLeader = co_leader.ID
    celule.Save()

    celule.Sign(co_leader)

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
            Message: "Usuário não encontrado",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celula não encontrada",
        }, 404
    }

    celule.Leader = leader.ID
    celule.Save()

    celule.Sign(leader)

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
            Message: "Celula não encontrada",
        }, 404
    }

    if db.First(&leader, "id = ?", celule.Leader).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Líder não encontrado",
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
            Message: "Celula superior não encontrada",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celula não encontrada",
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
            Message: "Celula não encontrada",
        }, 404
    }

    if db.First(&parent, "id = ?", celule.Parent).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celula Superior não encontrada",
        }, 404
    }
    return api.Response {
        Type: "Sucess",
        Data: parent,
    }, 200
}

func CeluleUnsignUser(r api.Request) (api.Response, int) {
    user := User{}
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celula não foi encontrado",
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
    if db.First(&user, "id = ?", r.PathVars["uid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["rid"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celula não foi encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, celule) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
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
            Message: "Celula não foi encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, celule) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    user_list := celule.GetUsers(page, limit)

    return api.Response{
        Type: "Sucess",
        Data: user_list,
    }, 200
}


func GetCeluleListByUser(r api.Request) (api.Response, int) {
    var limit, page int
    var church bool

    church, _ = sc.ParseBool(r.Conf["query"].(url.Values).Get("church"))
    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    user := User{}
    if (db.First(&user, "id = ?", r.PathVars["id"]).Error != nil) {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, user) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    celule_list := user.GetCelules(page, limit, church)

    return api.Response{
        Type: "Sucess",
        Data: celule_list,
    }, 200
}

func GetCeluleList(r api.Request) (api.Response, int) {
    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, nil) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    var limit, page int

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    celule_list := []Celule{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&celule_list)

    if e.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Error on creating of cover on database",
        }, 500
    }

    return api.Response{
        Type: "Sucess",
        Data: celule_list,
    }, 200
}

func CreateCeluleAddr(r api.Request) (api.Response, int) {
    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Celula não foi encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || !CheckPermissions(curr, celule) {
        msg := "Você não tem permissão para acessar isso"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprintf("Dados enviados são inválidos")
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

    celule.SetAddress(addr)
    return api.Response {
        Type: "Sucess",
        Data: addr,
    }, 200
}

func GetCeluleAddr(r api.Request) (api.Response, int) {
    u := Celule {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Celula não foi encontrado")
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

func CreateCeluleCover(r api.Request) (api.Response, int) {
    celule := Celule{}
    if db.First(&celule, "id = ?", r.PathVars["id"]).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Fiel não encontrado",
        }, 404
    }

    token := Token{}
    token.ID = r.Token

    if curr, ok := (token).GetUser();!ok {
        if !CheckPermissions(curr, celule) {
            api.SuperPut(curr)
            api.SuperPut(CheckPermissions(curr, celule))
            msg := "Você não tem permissão para acessar isso"
            api.Err(msg)
            return api.Response {
                Message: msg,
                Type:    "Error",
            }, 405
        }
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
            Message: "Erro interno desconhecido",
        }, 500
    }

    celule.SetCover(cover)
    return api.Response {
        Type: "Sucess",
        Data: cover,
    }, 200
}

func GetCeluleCover(r api.Request) ([]byte, int) {
    u := Celule {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("Fiel não encontrado")
        api.Err(msg)
        return []byte{}, 404
    }
    p := u.GetCover()
    return []byte(p.Render()), 200
}


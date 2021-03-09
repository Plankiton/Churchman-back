package church

import (
	"fmt"
	"time"

  "net/url"
  sc "strconv"
  mp "mime/multipart"

    "github.com/Coff3e/Api"
)

func GetUser(r api.Request) (api.Response, int) {
    u := User {}
    if db.First(&u, "id = ?", r.PathVars["id"]).Error != nil {
        msg := fmt.Sprint("User not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    token := Token{}
    token.ID = r.Token
    if curr, ok := (token).GetUser();!ok || CheckPermissions(curr, u) {
        msg := "Authentication fail, permission denied"
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 405
    }

    return api.Response {
        Type: "Success",
        Data: u,
    }, 200
}

func CreateUser(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("User create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    if (len(data)<5){
        msg := "User create fail, Obrigatory field"
        if (len(data)==4) {
            msg += "s"
        }
        msg += " missing: "
        for _, k := range []string{
            "email", "name", "pass", "born", "genre",
        } {
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

    res := db.First(&User {}, "email = ?", data["email"])
    if res.Error == nil {
        msg := fmt.Sprint("User create fail, user already registered")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 500
    }

    user := User {}

    api.MapTo(data, &user)
    born_time, _ := time.Parse(TimeLayout(), data["born"].(string))

    user.Born = born_time
    user.SetPass(data["pass"].(string))
    user.Create()

    return api.Response {
        Type: "Sucess",
        Data: user,
    }, 200
}

func UpdateUser(r api.Request) (api.Response, int) {
    if !api.ValidateData(r.Data, api.GenericJsonObj) {
        msg := fmt.Sprint("User create fail, data need to be a object")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 400
    }

    data := r.Data.(map[string]interface{})

    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("User delete fail, user not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    api.MapTo(data, &user)
    if _, e := data["born"];e {
        born_time, _ := time.Parse(TimeLayout(), data["born"].(string))
        user.Born = born_time
    }

    if _, e := data["pass"];e {
        user.SetPass(data["pass"].(string))
    }

    user.Save()

    return api.Response {
        Type: "Sucess",
        Data: user,
    }, 200
}

func DeleteUser(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("User delete fail, user not found")
        api.Err(msg)
        return api.Response {
            Message: msg,
            Type:    "Error",
        }, 404
    }

    user.Delete()

    return api.Response {
        Type: "Sucess",
        Message: "User deleted",
    }, 200
}

func CreateUserProfile(r api.Request) (api.Response, int) {
    user := User{}
    res := db.First(&user, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "User not found",
        }, 404
    }

    if r.Data == nil || !api.ValidateData(r.Data, api.GenericForm) {
        return api.Response{
            Type: "Error",
            Message: "Data must be a multipart-form",
        }, 400
    }

    data := r.Data.(*mp.Form)

    profile := File {}
    profile.Load(data)

    if db.First(&profile).Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    user.SetProfile(profile)
    return api.Response {
        Type: "Sucess",
        Data: profile,
    }, 200
}

func GetUserProfile(r api.Request) ([]byte, int) {
    u := User {}
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("User not found")
        api.Err(msg)
        return []byte{}, 404
    }
    p := u.GetProfile()

    return []byte(p.Render()), 200
}

func GetUserList(r api.Request) (api.Response, int) {
    var limit, page int

    limit, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("l"))
    page, _ = sc.Atoi(r.Conf["query"].(url.Values).Get("p"))

    user_list := []User{}
    offset := (page - 1) * limit
    e := db.Offset(offset).Limit(limit).Order("created_at desc, updated_at, id").Find(&user_list)

    if e.Error != nil {
        return api.Response{
            Type: "Error",
            Message: "Error on creating of profile on database",
        }, 500
    }

    query_response := []map[string]interface{}{}
    for _, u := range user_list {
        item := map[string]interface{}{}

        p := u.GetProfile()

        item["name"] = u.Name
        item["id"] = u.ID

        if p.ID != 0 {
            item["profile"] = map[string]string{
                "data": p.Render(),
                "alt_text": p.AltText,
            }
        }

        query_response = append(query_response, item)
    }

    return api.Response{
        Type: "Sucess",
        Data: query_response,
    }, 200
}

package church

import (
	"fmt"
	"time"

  mp "mime/multipart"
	"github.com/Coff3e/Api"
)

type User struct {
    api.User
    State   string   `json:"civil_state,omitempty" gorm:"default:sole"`
    ProfileId uint     `json:"-"`
}

func (self *User) SetProfile(profile File) {
    {
        tmp_profile := File{}
        e := db.First(&tmp_profile, "id = ?", self.ProfileId)
        if e.Error == nil {
            tmp_profile.Delete()
        }
    }

    self.ProfileId = profile.ID
    self.Save()
}

func (self *User) GetProfile() File {
    profile := File{}
    profile.ID = self.ProfileId
    e := db.First(&profile)
    if e.Error == nil {
        return profile
    }

    return File{}
}

func (self *User) GetRoles() []Role {
    e := db.First(self)
    if e.Error == nil {
        role_list := []uint{}
        roles := []Role{}
        e := db.Raw("SELECT r.id FROM roles r INNER JOIN user_roles ur INNER JOIN users u ON ur.role_id = r.id AND ur.user_id = u.id AND u.id = ?", self.ID).Find(&role_list)
        if e.Error == nil {
            e := db.Find(&roles, "id in ?", role_list)
            if e.Error == nil {
                return roles
            }
        }
    }

    return []Role{}
}

func GetUser(r api.Request) (api.Response, int) {
    u := User {}
    res := db.First(&u, "id = ?", r.PathVars["id"])
    if res.Error != nil {
        msg := fmt.Sprint("User not found")
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

func CreateUser(r api.Request) (api.Response, int) {
    if !validData(r.Data, generic_json_obj) {
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
    if !validData(r.Data, generic_json_obj) {
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

    if r.Data == nil || validData(r.Data, generic_form) {
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

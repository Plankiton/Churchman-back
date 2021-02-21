package church

import (
	"fmt"
	"time"

	"github.com/Coff3e/Api"
)

type User struct {
    api.User
    State   string   `json:"civil_state,omitempty" gorm:"default:sole"`
    ProfileId uint     `json:",empty"`
}

func (self *User) SetProfile(profile Image) {
    {
        tmp_profile := Image{}
        e := db.First(&tmp_profile, "id = ?", self.ProfileId)
        if e.Error == nil {
            tmp_profile.Delete()
        }
    }

    self.ProfileId = profile.ID
    self.Save()

    profile.Create()
}

func (self *User) GetProfile() Image {
    profile := Image{}
    e := db.First(&profile, "id = ?", self.ProfileId)
    if e.Error == nil {
        return profile
    }

    return Image{}
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

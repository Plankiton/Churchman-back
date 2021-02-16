package church

import (
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

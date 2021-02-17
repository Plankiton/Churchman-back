package church

import (
    "github.com/Coff3e/Api"
)

type User struct {
    api.User
    State   string   `json:"civil_state,omitempty" gorm:"default:sole"`
    ProfileId uint   `json:",empty" gorm:"index"`
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

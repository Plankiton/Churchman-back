package church

import (
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

        e = db.First(&tmp_profile, "filename = ? OR id = ?", profile.Filename, profile.ID)
        if e.Error != nil {
            tmp_profile.Create()
        }
    }

    self.Update(api.Dict { "profile_id": profile.ID })
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


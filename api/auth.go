package church

import "github.com/Coff3e/Api"

type Auth struct {
    api.Auth
    User    User  `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

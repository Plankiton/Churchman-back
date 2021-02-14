package church

import (
    "github.com/Coff3e/Api"
)

type User struct {
    api.User
    State   string   `json:"civil_state,omitempty" gorm:"default:sole"`
    Roles []Role   `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

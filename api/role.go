package church

import (
    "reflect"
    "github.com/Coff3e/Api"
)

type Role struct {
    api.Role
    Users []User   `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

type UserRoles struct {
    api.Model
    User    User  `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Role    Role  `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (self UserRoles) Sign(user User, role Role) (User, Role) {
    self.User = user
    self.Role = role

    user.Roles = append(user.Roles, role)
    role.Users = append(role.Users, user)

    self.Create()
    api.Log("Linked", api.ToLabel(user.ID, user.ModelType), user.Name, "to", api.ToLabel(role.ID, role.ModelType), role.Name)

    return user, role
}

func (self UserRoles) Unsign(user User, role Role) (User, Role) {
    self.User = user
    self.Role = role

    for i, r := range user.Roles {
        if reflect.DeepEqual(r, role) {
            user.Roles = append(user.Roles[:i-1], user.Roles[i+1:]...)
            break
        }
    }

    for i, u := range user.Roles {
        if reflect.DeepEqual(u, user) {
            role.Users = append(role.Users[:i-1], role.Users[i+1:]...)
            break
        }
    }

    self.Delete()
    api.Log("Unlinked", api.ToLabel(user.ID, user.ModelType), user.Name, "from", api.ToLabel(role.ID, role.ModelType), role.Name)

    return user, role
}

func (self Role) Sign(user *User) {
    link := UserRoles{}
    *user, self = link.Sign(*user, self)
}

func (self Role) Unsign(user *User) {
    link := UserRoles{}
    e := db.First(&link, "user_id = ? AND role_id = ?", *user.ID, self.ID)
    if e.Error == nil {
        *user, self = link.Unsign(*user, self)
    }
}

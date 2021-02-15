package church

import (
    "github.com/Coff3e/Api"
)

type Role struct {
    api.Role
}

type UserRole struct {
    api.UserRole
}

func (self UserRole) Sign(user User, role Role) (User, Role) {
    self.UserId = user.ID
    self.RoleId = role.ID

    self.Create()
    api.Log("Linked", api.ToLabel(user.ID, user.ModelType), user.Name, "to", api.ToLabel(role.ID, role.ModelType), role.Name)

    return user, role
}

func (self UserRole) Unsign(user User, role Role) (User, Role) {
    self.UserId = user.ID
    self.RoleId = role.ID

    self.Delete()
    api.Log("Unlinked", api.ToLabel(user.ID, user.ModelType), user.Name, "from", api.ToLabel(role.ID, role.ModelType), role.Name)

    return user, role
}

func (self Role) Sign(user User) (User, Role) {
    link := UserRole{}
    user, self = link.Sign(user, self)

    return user, self
}

func (self Role) Unsign(user User) (User, Role) {
    link := UserRole{}
    e := db.Where("user_id = ? AND role_id = ?", user.ID, self.ID).First(&link)
    if e.Error == nil {
        user, self = link.Unsign(user, self)
    }

    return user, self
}

func (self *Role) GetPersons() []User {
    e := db.First(self)
    if e.Error == nil {
        users := []User{}
        e := db.Preload("user_roles").Find(&users)
        if e.Error == nil {
            return users
        }
    }

    return []User{}
}

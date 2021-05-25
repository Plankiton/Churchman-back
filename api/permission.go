package church

import (
    "github.com/Coff3e/Api"
    str "strings"
)

func CheckPermissions(curr User, object interface {}) bool {
    if object == nil {
        api.SuperPut(curr.GetRoles(0,0))
        return CheckPermissions(curr, curr.GetRoles(0,0))
    }

    switch api.GetModelType(object) {
    case "[]Celule":
        for _, c := range (object.([]Celule)) {
            if CheckPermissions(curr, c) {
                return true
            }
        }
    case "Celule":
        celule := object.(Celule)
        if curr.ID == celule.Leader || curr.ID == celule.CoLeader {
            return true
        }
    case "[]Role":
        for _, c := range (object.([]Role)) {
            if CheckPermissions(curr, c) {
                return true
            }
        }
    case "Role":
        role := object.(Role)
        if str.ToLower(role.Name) == "pastor" {
            return true
        }
    case "User":
        user := object.(User)
        if CheckPermissions(curr, user.GetCelules(0,0)) ||
           CheckPermissions(curr, curr.GetRoles(0,0)) ||
           curr.ID == user.ID {
            return true
        }
    }

    return false
}

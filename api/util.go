package church

import (
    "fmt"
    "strings"
    // "github.com/Coff3e/Api"
)

func TimeLayout() string {
    return "2006-01-02T15:04:05.000Z"
}

func GetCeluleName(c Celule) string {
    if c.Type == "church" {
        return c.Name
    }

    iid := ""
    if c.IID/10 < 1 {
        iid = "0"
    }
    iid += fmt.Sprint(c.IID)

    parent := Celule{}
    if c.Parent != 0 && db.First(&parent, "id = ?", c.Parent).Error == nil && parent.Type != "church" {
        return fmt.Sprint(GetCeluleName(parent), ".", iid)
    }

    leader := User {}
    if c.Leader != 0 && db.First(&leader, "id = ?", c.Leader).Error == nil {
        iid = string(strings.ToUpper(leader.Genre)[0]) + iid
    }

    return iid
}

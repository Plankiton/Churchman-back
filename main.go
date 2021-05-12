package main

import (
    "os"
    "time"

    "gorm.io/gorm"
    "github.com/Coff3e/Church-app/api"
    "github.com/Coff3e/Api"
)

func main() {
    con_str := "host=localhost user=plankiton password=joaojoao dbname=church port=5432 sslmode=disable TimeZone=America/Araguaina"
    r := church.Church{}

    var db *gorm.DB
    var err error
    if os.Getenv("DEBUG_MODE") == "true" {
        api.Log("Entering on Debug mode, using sqlite database")
        db, err = r.SignDB("/tmp/debug.db", api.Sqlite)
    } else {
        api.Log("Trying to connect to postgresql")
        db, err = r.SignDB(con_str, api.Postgres)
    }
    if (err != nil) {
        api.Err("Database is dowm")
        os.Exit(1)
    }
    api.Log("Database connected with sucess")

    if db.First(&church.Role{}).Error != nil {
        pastor := church.Role{}
        pastor.Name = "Pastor"
        pastor.Create()

        root := church.User{}
        root.Email = "root@joao.com"
        root.SetPass("joao")
        root.Create()

        pastor.Sign(root)

        igreja_do_deus := church.Celule{}
        igreja_do_deus.Name = "Igreja do deus"
        igreja_do_deus.Type = "church"
        igreja_do_deus.Leader = root.ID
        igreja_do_deus.Create()

        // Horarios de cultos
        {
            week := []string{ "Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sabado" }
            cloc := []time.Time{ time.Now(), time.Now(), time.Now(), time.Now(), time.Now(), time.Now(), time.Now() }
            for d, _ := range week {
                event := church.Event{}
                event.Periodic = true
                event.WeeklyDay = uint(d+1)
                event.BeginAt = cloc[d]
                event.EndAt = cloc[d].Add(2 * time.Hour)
                event.Create()
                event.Sign(igreja_do_deus)
            }
        }

        link := church.UserCelule{}
        link.UserId = root.ID
        link.GroupId = igreja_do_deus.ID
        link.Create()
    }

    r.
    Add(
        "get", "/address/{cep}", nil, church.GetFromCEP,
    ).
    Add(
        "post", "/login", api.RouteConf {
            "need-auth": false,
        }, church.LogIn,
    ).
    Add(
        "post", "/verify", nil, church.Verify,
    ).
    Add(
        "post", "/logout", nil, church.LogOut,
    ).


    Add(
        "get", "/user/", nil, church.GetUserList,
    ).
    Add(
        "post", "/user", api.RouteConf {
            "need-auth": false,
        }, church.CreateUser,
    ).
    Add(
        "post", "/user/{id}/profile", nil, church.CreateUserProfile,
    ).
    Add(
        "get", "/user/{id}", nil, church.GetUser,
    ).
    Add(
        "post", "/user/{id}", nil, church.UpdateUser,
    ).
    Add(
        "get", "/user/{id}/profile", nil, church.GetUserProfile,
    ).
    Add(
        "get", "/user/{id}/roles", nil, church.GetRoleListByUser,
    ).
    Add(
        "get", "/user/{id}/celules", nil, church.GetCeluleListByUser,
    ).
    Add(
        "get", "/user/{id}/events", api.RouteConf {
            "need-auth": false,
        },  church.GetEventListByUser,
    ).



    Add(
        "get", "/role/{id}", nil, church.GetRole,
    ).
    Add(
        "get", "/role/", nil, church.GetRoleList,
    ).
    Add(
        "post", "/role/", nil, church.CreateRole,
    ).
    Add(
        "post", "/role/{id}", nil, church.UpdateRole,
    ).
    Add(
        "post", "/role/{rid}/sign/{uid}", nil, church.RoleSignUser,
    ).
    Add(
        "post", "/role/{rid}/unsign/{uid}", nil, church.RoleUnsignUser,
    ).
    Add(
        "delete", "/role/{id}", nil, church.DeleteRole,
    ).
    Add(
        "get", "/role/{id}/users", nil, church.GetUserListByRole,
    ).



    Add(
        "get", "/event/{id}", nil, church.GetEvent,
    ).
    Add(
        "get", "/event/", nil, church.GetEventList,
    ).
    Add(
        "post", "/event/", nil, church.CreateEvent,
    ).
    Add(
        "post", "/event/{id}", nil, church.UpdateEvent,
    ).
    Add(
        "post", "/event/{rid}/sign/{uid}", nil, church.EventSignUser,
    ).
    Add(
        "post", "/event/{rid}/unsign/{uid}", nil, church.EventUnsignUser,
    ).
    Add(
        "delete", "/event/{id}", nil, church.DeleteEvent,
    ).
    Add(
        "get", "/event/{id}/users", nil, church.GetUserListByEvent,
    ).
    Add(
        "post", "/event/{id}/cover", nil, church.CreateEventCover,
    ).
    Add(
        "get", "/event/{id}/cover", nil, church.GetEventCover,
    ).
    Add(
        "post", "/event/{id}/address", nil, church.CreateEventAddr,
    ).
    Add(
        "get", "/event/{id}/address", nil, church.GetEventAddr,
    ).



    Add(
        "get", "/celule/{id}", nil, church.GetCelule,
    ).
    Add(
        "get", "/celule/", nil, church.GetCeluleList,
    ).
    Add(
        "post", "/celule/", nil, church.CreateCelule,
    ).
    Add(
        "post", "/celule/{id}", nil, church.UpdateCelule,
    ).
    Add(
        "post", "/celule/{rid}/sign/{uid}", nil, church.CeluleSignUser,
    ).
    Add(
        "post", "/celule/{rid}/unsign/{uid}", nil, church.CeluleUnsignUser,
    ).
    Add(
        "delete", "/celule/{id}", nil, church.DeleteCelule,
    ).
    Add(
        "get", "/celule/{id}/users", nil, church.GetUserListByCelule,
    ).
    Add(
        "get", "/celule/{id}/events", nil, church.GetEventListByCelule,
    ).
    Add(
        "post", "/celule/{id}/address", nil, church.CreateCeluleAddr,
    ).
    Add(
        "get", "/celule/{id}/address", nil, church.GetCeluleAddr,
    ).
    Add(
        "post", "/celule/{id}/co-leader/{uid}", nil, church.CeluleSetCoLeader,
    ).
    Add(
        "get", "/celule/{id}/co-leader", nil, church.CeluleGetCoLeader,
    ).
    Add(
        "post", "/celule/{id}/leader/{uid}", nil, church.CeluleSetLeader,
    ).
    Add(
        "get", "/celule/{id}/leader", nil, church.CeluleGetLeader,
    ).
    Add(
        "post", "/celule/{id}/parent/{uid}", nil, church.CeluleSetParent,
    ).
    Add(
        "get", "/celule/{id}/parent", nil, church.CeluleGetParent,
    )

    r.Run("/", 8000)
}

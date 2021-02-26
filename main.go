package main

import (
    "os"

    "./api"
    "github.com/Coff3e/Api"
)

func main() {
    con_str := "host=localhost user=plankiton password=joaojoao dbname=church port=5432 sslmode=disable TimeZone=America/Araguaina"
    r := church.Church{}

    var err error
    if os.Getenv("DEBUG_MODE") == "true" {
        _, err = r.SignDB("/tmp/debug.db", api.Sqlite)
    } else {
        _, err = r.SignDB(con_str, api.Postgres)
    }
    if (err != nil) {
        os.Exit(1)
    }
    api.Log("Database connected with sucess")


    r.
    Add(
        "post", "/login", api.RouteConf {
            "need-auth": false,
        }, church.LogIn,
    ).
    Add(
        "post", "/logout", nil, church.LogOut,
    ).
    Add(
        "post", "/verify", nil, church.Verify,
    ).


    Add(
        "get", "/user/", nil, church.GetUserList,
    ).
    Add(
        "post", "/user", nil, church.CreateUser,
    ).
    Add(
        "post", "/user/{id}/profile", nil, church.CreateUserProfile,
    ).
    Add(
        "get", "/user/{id}", nil, church.GetUser,
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
        "get", "/user/{id}/events", nil, church.GetEventListByUser,
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
    )



    r.Run("/", 8000)
}

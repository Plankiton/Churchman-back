package church

import (
    mp "mime/multipart"
    "github.com/Coff3e/Api"
    "reflect"
)

type Church struct {
    api.API
}

func generic_interface () reflect.Type {
    var i interface{}
    return reflect.TypeOf(&i).Elem()
}

func generic_string () reflect.Type {
    var i string
    return reflect.TypeOf(i)
}

func generic_json_obj() reflect.Type {
    return reflect.MapOf(generic_string(), generic_interface())
}

func generic_json_array() reflect.Type {
    return reflect.ArrayOf(-1, reflect.MapOf(generic_string(), generic_interface().Elem()))
}

func generic_form() reflect.Type {
    var i mp.Form
    return reflect.TypeOf(i)
}

func validData(data interface{}, t func()reflect.Type) bool {
    if data == nil {
        return false
    }

    if reflect.TypeOf(data).Kind() == t().Kind() {
        return true
    }

    return false
}

package main

import "pretty-print/prettyprint"

type Student struct {
    Name string
    Age int
    Addr address
}


type address struct {
    street string
    no string
}

func main() {

    v := map[string]interface{}{
        "str":   "foo",
        "num":   100,
        "bool":  false,
        "null":  nil,
        "array": []string{"foo", "bar", "baz"},
        "map": map[string]interface{}{
            "foo": "bar",
        },
    }

    a := address{"长安街", "18"}
    s := Student{"张三", 30, a}
    v["info"] = s

    prettyprint.P(v)
}

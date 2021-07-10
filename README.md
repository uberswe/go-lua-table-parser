# go-lua-table-parser
A simple go package which takes a lua table as input and parses it to a map[string]interface{} where interface{} is either more map[string]interface{} or map[string]string.

This package uses the [github.com/yuin/gopher-lua](https://github.com/yuin/gopher-lua) package.

Usage:
```go
import parser "github.com/uberswe/go-lua-table-parser"

luaString := `Vars =
{
    ["key"] = 
    {
        ["someOtherKey"] = "value"
    }
}`

func main() {
    parser.Parse(luaString, "Vars")
}
```

package parser

import (
	lua "github.com/yuin/gopher-lua"
)

var (
	result map[string]interface{}
	keys   []string
)

// Parse takes a lua table as a string (s) and then the variable name of the lua table as arguments and returns a map[string]interface{} where interface{} is either more map[string]interface{} or map[string]string
func Parse(s string, variable string) (map[string]interface{}, error) {
	keys = []string{}
	result = map[string]interface{}{}
	l := lua.NewState()
	defer l.Close()
	if err := l.DoString(s); err != nil {
		return result, err
	}
	lv := l.GetGlobal(variable)
	if tbl, ok := lv.(*lua.LTable); ok {
		tbl.ForEach(recursiveLoop)
	}
	return result, nil
}

func recursiveLoop(k lua.LValue, v lua.LValue) {
	key := ""
	if str, ok := k.(lua.LString); ok {
		key = str.String()
	}
	if n, ok := k.(lua.LNumber); ok {
		key = n.String()
	}
	if tbl, ok := v.(*lua.LTable); ok {
		keys = append(keys, key)
		tbl.ForEach(recursiveLoop)
		_, keys = keys[len(keys)-1], keys[:len(keys)-1]
	} else if str, ok := v.(lua.LString); ok {
		result = setKeyValue(&result, key, str.String(), keys)
	} else if n, ok := v.(lua.LNumber); ok {
		result = setKeyValue(&result, key, n.String(), keys)
	}
}

func setKeyValue(r *map[string]interface{}, k string, v string, ks []string) map[string]interface{} {
	t := *r
	if len(ks) > 0 {
		x, ks := ks[0], ks[1:]
		if t[x] == nil {
			t[x] = map[string]interface{}{}
		}
		tt := t[x].(map[string]interface{})
		t[x] = setKeyValue(&tt, k, v, ks)
		return t
	}
	t[k] = v
	return t
}

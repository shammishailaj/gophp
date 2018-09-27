package phpvariables

import (
	"fmt"
)

func Var_dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

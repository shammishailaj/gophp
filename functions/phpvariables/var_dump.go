package phpvariables

import (
	"fmt"
)

func var_dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

package b

import (
	"fmt"
	c "github.com/carolynvs/deptest-transcons-c"
)

func B() {
	fmt.Println("b did a thing")
	c.C()
}

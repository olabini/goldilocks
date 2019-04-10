package goldilocks

import (
	"fmt"

	"github.com/olabini/goldilocks/internal/field"
)

func printElement(name string, e *field.Element) {
	fmt.Printf("%s = %s\n", name, e.String())
}

func printPoint(name string, p *point) {
	fmt.Printf("%s = {\n", name)
	fmt.Printf("    x = %s\n", p.x.String())
	fmt.Printf("    y = %s\n", p.y.String())
	fmt.Printf("    z = %s\n", p.z.String())
	fmt.Printf("    t = %s\n", p.t.String())
	fmt.Printf("} // %s\n", name)
}

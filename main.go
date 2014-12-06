package main

import (
	"fmt"

	"github.com/ian-kent/purl/perl"
)

func main() {
	purl := &perl.Purl{}
	purl.Init()
	defer purl.Destroy()

	fmt.Println("Purl::Test =>", purl.Eval(`Purl::Test()`))
	fmt.Println("Assign expr =>", purl.Eval(`$a = "foo"`))
	fmt.Println("Get expr =>", purl.Eval(`$a`))
	fmt.Println("Get expr =>", purl.Eval(`$foo = 1`))

	purl.Eval(`
sub test {
	my ($foo, $bar) = @_;
	return "$bar:$foo";
}
`)
	fmt.Println("Invoke =>", purl.Eval(`test("foo", "bar")`))

	purl.RegisterXS("Purl::Go::Test", func(args ...string) string {
		fmt.Println("In Purl::Go::Test XS function!")
		return "hi!"
	})
	fmt.Println("Invoke custom XS =>", purl.Eval(`Purl::Go->Test()`))
	fmt.Println("Invoke custom XS =>", purl.Eval(`Purl::Go->Test()`))
}

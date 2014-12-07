package main

import (
	"testing"

	"github.com/ian-kent/purl/perl"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPurl(t *testing.T) {
	Convey("Basic perl should work", t, func() {
		purl := &perl.Purl{}
		purl.Init()
		defer purl.Destroy()

		So(purl.Eval(`Purl::Test()`), ShouldEqual, "")
		So(purl.Eval(`$a = "foo"`), ShouldEqual, "foo")
		So(purl.Eval(`$a`), ShouldEqual, "foo")
		So(purl.Eval(`$foo = 1`), ShouldEqual, "1")

		purl.Eval(`
sub test {
	my ($foo, $bar) = @_;
	return "$bar:$foo";
}
`)
		So(purl.Eval(`test("foo", "bar")`), ShouldEqual, "bar:foo")

	})

	Convey("XS functions should work", t, func() {
		purl := &perl.Purl{}
		purl.Init()
		defer purl.Destroy()

		purl.RegisterXS("Purl::Go::Test", func(args ...string) string {
			return "hi!"
		})

		So(purl.Eval(`Purl::Go->Test()`), ShouldEqual, "hi!")
		So(purl.Eval(`Purl::Go->Test()`), ShouldEqual, "hi!")
	})

	Convey("XS functions with parameters should work", t, func() {
		purl := &perl.Purl{}
		purl.Init()
		defer purl.Destroy()

		var cbArgs []string
		purl.RegisterXS("Purl::Go::Test", func(args ...string) string {
			cbArgs = args
			return "hi!"
		})

		So(purl.Eval(`Purl::Go->Test()`), ShouldEqual, "hi!")
		So(len(cbArgs), ShouldEqual, 1)
		So(cbArgs[0], ShouldEqual, "Purl::Go")

		So(purl.Eval(`Purl::Go->Test("foo", "bar")`), ShouldEqual, "hi!")
		So(len(cbArgs), ShouldEqual, 3)
		So(cbArgs[0], ShouldEqual, "Purl::Go")
		So(cbArgs[1], ShouldEqual, "foo")
		So(cbArgs[2], ShouldEqual, "bar")
	})

}

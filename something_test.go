package something

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDoSomething(t *testing.T) {
	Convey("hi", t, func(c C) {
		err := doSomething(c)
		So(err, ShouldBeNil)
	})
}

package something

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDoSomething(t *testing.T) {
	Convey("hello", t, func(c C) {
		Convey("world", func() {
			ctx := context.TODO()
			err := doSomething(ctx)
			So(err, ShouldBeNil)
		})
	})
}

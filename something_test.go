package something

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDoSomething(t *testing.T) {
	Convey("hi", t, func(c C) {
		ctx := context.TODO()
		err := doSomething(ctx)
		So(err, ShouldBeNil)
	})
}

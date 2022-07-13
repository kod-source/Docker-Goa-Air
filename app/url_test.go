package main

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/Docker-Goa-Air/app"
	"github.com/kod-source/Docker-Goa-Air/app/test"
	"github.com/shogo82148/goa-v1"
	"github.com/shogo82148/pointer"
)

func TestController_URLAdd(t *testing.T) {
	servic := goa.New("test")
	ctx := context.Background()
	ctrl := NewURLController(servic)

	t.Run("[OK] 正常動作", func(t *testing.T) {
		left := 100
		middle := 100
		right := 100
		_, got := test.URLAddURLOK(t, ctx, servic, ctrl, left, middle, right)

		want := &app.URL{
			Amount: pointer.Int(100),
			ID:     100,
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("User value is mismatch (-got +want):\n%s", diff)
		}
	})
}

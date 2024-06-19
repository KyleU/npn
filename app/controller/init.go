package controller

import (
	"context"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/user"
	"github.com/kyleu/npn/app/util"
)

// Initialize app-specific system dependencies.
func initApp(_ context.Context, _ *app.State, _ util.Logger) error {
	user.SetPermissions(false,
		user.Perm("/admin", "github:kyleu.com", true),
		user.Perm("/admin", "github:npn.dev", true),
		user.Perm("/admin", "*", false),
		user.Perm("/", "*", true),
	)
	return nil
}

// Configure app-specific data for each request.
func initAppRequest(_ *app.State, _ *cutil.PageState) error {
	return nil
}

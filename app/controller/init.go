package controller

import (
	"context"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
)

// Initialize app-specific system dependencies.
func initApp(_ context.Context, _ *app.State, _ util.Logger) error {
	return nil
}

// Configure app-specific data for each request.
func initAppRequest(_ *app.State, _ *cutil.PageState) error {
	return nil
}

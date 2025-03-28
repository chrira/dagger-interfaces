// A generated module for Example functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	//"dagger/example/internal/dagger"
	"fmt"
)

type Example struct{}


// implement Fooer interface
func (m *Example) Foo(
	ctx context.Context,
	bar int,
	// +optional
	name string,
	// +optional
	name2 string,
) (string, error) {
	return fmt.Sprintf("%s - %s is: %d", name, name2, bar), nil
}

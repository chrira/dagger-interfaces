// A generated module for Bar functions
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
	"dagger/bar/internal/dagger"
)

type Bar struct{}

// Returns a container that echoes whatever string argument is provided
func (m *Bar) ContainerEcho(
	ctx context.Context,
) (string, error) {
	// Because `Bar` implements `Fooer`, the conversion function `AsMyModuleFooer` has been generated.
	stringArg, err := dag.MyModule().Foo(ctx, dag.Example().AsMyModuleFooer())

	if err != nil {
		return "", err
	}
	return stringArg, nil
}

// Returns a container that echoes whatever string argument is provided
func (m *Bar) ContainerEcho0(
	ctx context.Context,
	bar int,
	// +optional
	name string,
) (string) {
	res, err := dag.MyModule().ContainerEcho0(ctx, bar, dagger.MyModuleContainerEcho0Opts{Name: name})
	if err != nil {
		return ""
	}
	return res
}


// Returns a container that echoes whatever string argument is provided
func (m *Bar) ContainerEcho1(
	ctx context.Context,
	bar int,
	// +optional
	name string,
	// +optional
	// +default="default"
	name2 string,
) (string) {
	// Because `Bar` implements `Fooer`, the conversion function `AsMyModuleFooer` has been generated.
	fooer := dag.Example().AsMyModuleFooer()
	res, err := dag.MyModule().ContainerEcho(ctx, bar, fooer, dagger.MyModuleContainerEchoOpts{Name: name, Name2: name2})
	if err != nil {
		return ""
	}
	return res
}

/*
// Returns a container that echoes whatever string argument is provided
func (m *Bar) ContainerEcho2(
	ctx context.Context,
	bar int,
	// +optional
	name string,
) (string, error) {
	// Because `Bar` implements `Fooer`, the conversion function `AsMyModuleFooer` has been generated.
	stringArg, err := dag.MyModule().ContainerEcho(ctx, dag.Example().AsMyModuleFooer())

	if err != nil {
		return "", err
	}
	return stringArg, nil
}
*/
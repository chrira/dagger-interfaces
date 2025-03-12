// A generated module for MyModule functions
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
	"dagger/my-module/internal/dagger"
)

type MyModule struct{}

type Fooer interface {
	DaggerObject
	Foo(
		ctx context.Context,
		bar int,
		// +optional
		name string,
		// +optional
		// +default="awesome"
		name2 string,
	) (string, error)
}

func (m *MyModule) Foo(
	ctx context.Context,
	fooer Fooer,
) (string, error) {
	return fooer.Foo(ctx, 42, "Puzzle", "")
}


// Returns a container that echoes whatever string argument is provided
func (m *MyModule) ContainerEcho(
	ctx context.Context,
	bar int,
	// +optional
	name string,
	fooer Fooer,
) (string, error) {
	stringArg, err := fooer.Foo(ctx, bar, name, "")
	if err != nil {
		return "", err
	}
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg}).Stdout(ctx)
}

// Returns a container that echoes whatever string argument is provided
func (m *MyModule) ContainerEcho0(
	ctx context.Context,
	bar int,
	// +optional
	name string,
) (string, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", name}).Stdout(ctx)
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *MyModule) GrepDir(
	ctx context.Context,
	directoryArg *dagger.Directory,
	pattern string,
) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

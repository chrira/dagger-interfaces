# Dagger Interfaces

https://docs.dagger.io/api/interfaces/

##  Setup

### Modules

```sh
dagger init --sdk go --name Bar --source dagger

mkdir fooer
cd fooer/
dagger init --sdk go --name MyModule --source dagger

mkdir impl
cd impl/
dagger init --sdk go --name Example --source dagger
```

### Tests

```sh
dagger install ./fooer/
dagger install ./impl/

```

## Test

### pre optional


run interfaced method:

```sh
dagger call container-echo
number is: 42
```

fooer infos:

```sh
dagger -m fooer/ functions
✔ connect 0.3s
✔ load module 3.0s

Name       Description
grep-dir   Returns lines that match a pattern in the files of the provided Directory

Skipped 2 function(s) with unsupported types: foo, container-echo
```
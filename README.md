# Test Vendoring

This repo contains a GOPATH with 3 copies of a package called `projA`.

The first copy is in the default GOPATH location under `$GOPATH/src/origA/userA/projA`.
The second copy is vendored under `$GOPATH/src/myhost/myuser/myproject/vendor/origA/userA/projA/`.
The third copy is vendored one level deeper under `GOPATH/src/myhost/myuser/myproject/vendor/origB/userB/projB/vendor/origA/userA/projA/`.

```
.
└── src
    ├── myhost
    │   └── myuser
    │       └── myproject
    │           ├── main.go
    │           └── vendor
    │               ├── origA
    │               │   └── userA
    │               │       └── projA
    │               │           └── main.go
    │               └── origB
    │                   └── userB
    │                       └── projB
    │                           ├── main.go
    │                           └── vendor
    │                               └── origA
    │                                   └── userA
    │                                       └── projA
    │                                           └── main.go
    └── origA
        └── userA
            └── projA
                └── main.go
```

The goal is to illustrate which version of `projA` is used based on the package
which imports it. The program at `$GOPATH/src/myhost/myuser/myproject/main.go`
first calls `projA` from its vendored location, then through `projB` which also
vendors it. The result is displayed by main.go:

```
$ go run src/myhost/myuser/myproject/main.go
A comes from $GOPATH/src/myhost/myuser/myproject/vendor/origA/userA/projA/
B.A comes from $GOPATH/src/myhost/myuser/myproject/vendor/origB/userB/projB/vendor/origA/userA/projA/
```

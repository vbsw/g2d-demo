# g2d-demo

[![Go Reference](https://pkg.go.dev/badge/github.com/vbsw/g2d-demo.svg)](https://pkg.go.dev/github.com/vbsw/g2d-demo) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/g2d-demo)](https://goreportcard.com/report/github.com/vbsw/g2d-demo) [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

## About
g2d-demo is a graphic application that uses g2d framework to test and to demonstrate basic idea of using g2d. It is published on <https://github.com/vbsw/g2d-demo>.

## Compile
Install Go (<https://golang.org/doc/install>). For Cgo install a C compiler (<https://jmeubank.github.io/tdm-gcc/>).

For Windows:
To compile an executable that doesn't open a console, use

	-ldflags -H=windowsgui

## Controls
	1 - 5     spawn 1, 10, 100, 1000, 10000 entities
	q, w      de-/increment size of entities by 1 pixel
	a, s      de-/increment size of entities by half/twice
	k, l      de-/increment movement speed of entities
	o         set original size
	r         toggle rotation
	m         toggle movement
	j         switch between no, auto and custom mipmaps
	v         toggle vsync
	b         toggle window borders on/off
	d         toggle window dragable on/off
	t         toggle anti-aliasing (offscreen buffer, only)
	i         print stats (UPS, FPS, ...)
	c         clear screen
	f         fullscreen
	n         show new window

## Copyright
See file COPYRIGHT.

## References
- https://go.dev/doc/install
- https://jmeubank.github.io/tdm-gcc/
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
- https://github.com/golang/go/wiki/cgo
- https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies
- https://pkg.go.dev/cmd/link

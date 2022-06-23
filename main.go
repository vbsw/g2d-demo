/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package main tests and demonstrates basic idea of using g2d.
package main

import (
	"fmt"
	"github.com/vbsw/g2d"
)

func main() {
	g2d.Start(new(DemoEngine))
}

type DemoEngine struct {
	g2d.Engine
	params parameters
}

func (demo *DemoEngine) ParseOSArgs() error {
	infoOnly, err := demo.params.parseOSArgs()
	demo.SetInfoOnly(infoOnly)
	return err
}

func (demo *DemoEngine) Info() {
	fmt.Println("...parameters not implemented, yet")
}

func (demo *DemoEngine) CreateWindow() {
	var builder demoWindowBuilder
	builder.CreateWindow()
}

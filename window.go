/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"github.com/vbsw/g2d"
)

type demoWindowBuilder struct {
	g2d.WindowBuilder
}

type demoWindow struct {
}

func (builder *demoWindowBuilder) CreateWindow() {
	builder.ClientWidth = 640
	builder.ClientHeight = 480
	builder.Resizable = true
	builder.Centered = true
	builder.Handler = new(demoWindow)
	builder.WindowBuilder.CreateWindow()
}

func (window *demoWindow) foo() {
}

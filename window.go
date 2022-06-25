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

type tDemoWindow struct {
	g2d.Window
	params *tParameters
}

func newDemoWindow(params *tParameters) *tDemoWindow {
	window := new(tDemoWindow)
	window.params = params
	return window
}

func (window *tDemoWindow) Close() (bool, error) {
	return true, nil
}

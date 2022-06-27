/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"fmt"
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

func (window *tDemoWindow) Create() error {
	fmt.Println("create")
	return nil
}

func (window *tDemoWindow) Show() error {
	fmt.Println("show")
	return nil
}

func (window *tDemoWindow) KeyDown(key int, repeated uint) error {
	if repeated == 0 {
		if key == 41 { // esc
			window.Cmd.CloseReq = true
		} else if key == 9 { // f
			window.Props.Fullscreen = !window.Props.Fullscreen
		} else if key == 23 { // t
			fmt.Println("time", window.Time.NanosEvent)
		} else {
			fmt.Println("key down", key)
		}
	} else {
	}
	return nil
}

func (window *tDemoWindow) KeyUp(key int) error {
	return nil
}

func (window *tDemoWindow) Close() (bool, error) {
	return true, nil
}

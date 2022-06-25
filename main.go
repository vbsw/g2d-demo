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
	var params tParameters
	infoOnly, err := params.parseOSArgs()
	if err == nil {
		if !infoOnly {
			g2d.Init(nil)
			g2d.Show(newDemoWindow(&params))
			g2d.ProcessEvents()
		} else {
			printInfo(&params)
		}
	}
	if err == nil {
		err = g2d.Err
	}
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

func printInfo(params *tParameters) {
	fmt.Println("...parameters not implemented, yet")
}

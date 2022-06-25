/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"github.com/vbsw/golib/osargs"
)

type tParameters struct {
}

func (params *tParameters) parseOSArgs() (bool, error) {
	args := osargs.New()
	infoOnly := len(args.Values) > 0
	return infoOnly, nil
}

/*
 *          Copyright 2025, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package main. This is g2d demo.
package main

import (
	"embed"
	"fmt"
	"github.com/vbsw/g2d"
	"image"
	"image/png"
	"math/rand"
	"runtime"
	"time"
)

const (
	padding = 3
	speed   = 0.09
)

var (
	//go:embed *.png
	fs            embed.FS
	windowCounter int
	imgNames      []string
	imgWidths     []int
	imgHeights    []int
	imgIncX       []float32
	imgIncY       []float32
	random        *rand.Rand
	scale         float32
	infoCount     int
	currWnd       int
)

func init() {
	// run main on main thread.
	runtime.LockOSThread()
	imgNames = []string{"chibi0.png", "chibi1.png", "chibi2.png", "chibi3.png", "chibi4.png"}
	imgWidths = []int{512, 399, 400, 347, 418}
	imgHeights = []int{443, 512, 512, 512, 512}
	imgIncX = []float32{1.0, 399.0 / 512.0, 400.0 / 512.0, 347.0 / 512.0, 418.0 / 512.0}
	imgIncY = []float32{443.0 / 512.0, 1.0, 1.0, 1.0, 1.0}
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	scale = 0.125
	currWnd = -1
}

func imageFromEmbededPNG(fileName string) (image.Image, error) {
	file, err := fs.Open(fileName)
	if err == nil {
		defer file.Close()
		return png.Decode(file)
	}
	return nil, err
}

func printUsage() {
	fmt.Println("CONTROLS")
	fmt.Println("  1 - 5     spawn 1, 10, 100, 1000, 10000 entities")
	fmt.Println("  q, w      de-/increment size of entities by 1 pixel")
	fmt.Println("  a, s      de-/increment size of entities by half/twice")
	fmt.Println("  k, l      de-/increment movement speed of entities")
	fmt.Println("  o         set original size")
	fmt.Println("  r         toggle rotation")
	fmt.Println("  m         toggle movement")
	fmt.Println("  j         switch between no, auto and custom mipmaps")
	fmt.Println("  v         toggle vsync")
	fmt.Println("  b         toggle window borders on/off")
	fmt.Println("  d         toggle window dragable on/off")
	fmt.Println("  t         toggle anti-aliasing (offscreen buffer, only)")
	fmt.Println("  i         print stats (UPS, FPS, ...)")
	fmt.Println("  c         clear screen")
	fmt.Println("  f         fullscreen")
	fmt.Println("  h         show new window")
}

func main() {
	g2d.Init()
	g2d.MainLoop(newDemoWindow())
	if g2d.Err != nil {
		fmt.Println("error:", g2d.Err.Error())
	}
}

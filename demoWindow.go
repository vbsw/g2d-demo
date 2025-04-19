/*
 *          Copyright 2025, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"fmt"
	"github.com/vbsw/g2d"
)

type demoWindow struct {
	g2d.WindowImpl
	id           int
	title        string
	clientWidth  int
	clientHeight int
	layer0       *g2d.RectanglesLayer
	entities     []*g2d.Rectangle
	entImgIdxs   []int
	counter      int
	rotating        bool
	rotSpeed     []float32
	moving        bool
	movX, movY   []float32
	incX, incY   float32
	mipMapOn     bool
}

func newDemoWindow() *demoWindow {
	wnd := new(demoWindow)
	wnd.id = windowCounter
	if wnd.id == 0 {
		wnd.title = "g2d Demo"
	} else {
		wnd.title = fmt.Sprintf("g2d Demo #%i", wnd.id)
	}
	wnd.clientWidth = 1024
	wnd.clientHeight = 768
	windowCounter++
	return wnd
}

func (wnd *demoWindow) OnConfig(config *g2d.Configuration) error {
	config.Title = wnd.title
	config.ClientWidth = wnd.clientWidth
	config.ClientHeight = wnd.clientHeight
	if wnd.id == 0 {
		printUsage()
	}
	return nil
}

func (wnd *demoWindow) OnCreate() error {
	wnd.Gfx.VSync = true
	for i := range imgNames {
		wnd.Gfx.LoadTexture(newTextureLoader(i, imgNames[i]))
	}
	return nil
}

func (wnd *demoWindow) OnShow() error {
	wnd.layer0 = new(g2d.RectanglesLayer)
	wnd.layer0.Enabled = true
	wnd.entities = make([]*g2d.Rectangle, 0, 50000)
	wnd.entImgIdxs = make([]int, 0, 50000)
	wnd.rotSpeed = make([]float32, 0, 50000)
	wnd.Gfx.Layers = append(wnd.Gfx.Layers, wnd.layer0)
	return nil
}

func (wnd *demoWindow) OnTextureLoaded(texture g2d.Texture) error {
	if texture.Id() < len(imgNames) {
		// load as a mipmap
		// (normaly you wouldn't load a texture twice; this for tests, only)
		wnd.Gfx.LoadTexture(texture)
	}
	wnd.layer0.UseTexture(texture.Id(), texture.Id())
	return nil
}

func (wnd *demoWindow) OnUpdate() error {
	for i, rect := range wnd.entities {
		if wnd.rotating {
			rotSpeed := wnd.rotSpeed[i]
			alpha := rect.RotAlpha + float32(wnd.Stats.DeltaTime)*rotSpeed
			for alpha > 360 {
				alpha -= 360
			}
			rect.RotAlpha = alpha
		}
		if wnd.moving {
			if rect.X < padding && wnd.movX[i] < 0 {
				wnd.movX[i] = -1*wnd.movX[i]
			} else if rect.X > float32(wnd.Props.ClientWidth)-(rect.Width+padding) && wnd.movX[i] > 0 {
				wnd.movX[i] = -1*wnd.movX[i]
			}
			if rect.Y < padding && wnd.movY[i] < 0 {
				wnd.movY[i] = -1*wnd.movY[i]
			} else if rect.Y > float32(wnd.Props.ClientHeight)-(rect.Height+padding) && wnd.movY[i] > 0 {
				wnd.movY[i] = -1*wnd.movY[i]
			}
			rect.X, rect.Y = wnd.movX[i]*float32(wnd.Stats.DeltaTime)*speed+rect.X, wnd.movY[i]*float32(wnd.Stats.DeltaTime)*speed+rect.Y
		}
		if wnd.rotating {
			rotSpeed := wnd.rotSpeed[i]
			alpha := rect.RotAlpha + float32(wnd.Stats.DeltaTime)*rotSpeed
			for alpha > 360 {
				alpha -= 360
			}
			rect.RotAlpha = alpha
		}
	}
	if len(wnd.entities) > 0 {
		wnd.Update()
	}
	return nil
}

func (wnd *demoWindow) OnResize() error {
	clW, clH := float32(wnd.Props.ClientWidth), float32(wnd.Props.ClientHeight)
	// speed up moving (when out of screen)
	for i, rect := range wnd.entities {
		if rect.X + rect.Width > clW {
			movXAbs, signX := wnd.movX[i], float32(1.0)
			if movXAbs < 0 {
				signX = -1.0
				movXAbs *= -1.0
			}
			if movXAbs < 0.4 {
				for movXAbs < 0.4 {
					movXAbs *= 2.0
				}
				wnd.movX[i] = movXAbs * signX
			}
		}
		if rect.Y + rect.Height > clH {
			movYAbs, signY := wnd.movY[i], float32(1.0)
			if movYAbs < 0 {
				signY = -1.0
				movYAbs *= -1.0
			}
			if movYAbs < 0.4 {
				for movYAbs < 0.4 {
					movYAbs *= 2.0
				}
				wnd.movY[i] = movYAbs * signY
			}
		}
	}
	return nil
}

func (wnd *demoWindow) OnKeyDown(keyCode int, repeated uint) error {
	if keyCode == 14 { // K
		for i := range wnd.entities {
			wnd.movX[i] /= 1.1
			wnd.movY[i] /= 1.1
		}
	} else if keyCode == 15 { // L
		for i := range wnd.entities {
			wnd.movX[i] *= 1.1
			wnd.movY[i] *= 1.1
		}
	} else if keyCode == 20 { // Q
		wnd.incX -= 1
		wnd.incY -= 1
		for i, rect := range wnd.entities {
			imgIndex := wnd.entImgIdxs[i]
			incX := imgIncX[imgIndex]
			incY := imgIncY[imgIndex]
			rect.Width -= incX
			rect.Height -= incY
			rect.RotX -= incX / 2.0
			rect.RotY -= incY / 2.0
		}
	} else if keyCode == 26 { // W
		wnd.incX += 1
		wnd.incY += 1
		for i, rect := range wnd.entities {
			imgIndex := wnd.entImgIdxs[i]
			incX := imgIncX[imgIndex]
			incY := imgIncY[imgIndex]
			rect.Width += incX
			rect.Height += incY
			rect.RotX += incX / 2.0
			rect.RotY += incY / 2.0
		}
	} else if keyCode == 30 { // 1
		wnd.spawn(1)
	} else if keyCode == 31 { // 2
		wnd.spawn(10)
	} else if keyCode == 32 { // 3
		wnd.spawn(100)
	} else if keyCode == 33 { // 4
		wnd.spawn(1000)
	} else if keyCode == 34 { // 5
		wnd.spawn(10000)
	} else if repeated == 0 {
		if keyCode == 4 { // A
			scale /= 2
			for i, rect := range wnd.entities {
				imgIndex := wnd.entImgIdxs[i]
				rect.Width = float32(imgWidths[imgIndex]) * scale
				rect.Height = float32(imgHeights[imgIndex]) * scale
				rect.RotX = rect.Width / 2
				rect.RotY = rect.Height / 2
			}
		} else if keyCode == 22 { // S
			scale *= 2
			for i, rect := range wnd.entities {
				imgIndex := wnd.entImgIdxs[i]
				rect.Width = float32(imgWidths[imgIndex]) * scale
				rect.Height = float32(imgHeights[imgIndex]) * scale
				rect.RotX = rect.Width / 2
				rect.RotY = rect.Height / 2
			}
		} else if keyCode == 6 { // C
			for _, rect := range wnd.entities {
				wnd.layer0.Release(rect)
			}
			wnd.entities = wnd.entities[:0]
			wnd.entImgIdxs = wnd.entImgIdxs[:0]
			wnd.rotSpeed = wnd.rotSpeed[:0]
			wnd.movX = wnd.movX[:0]
			wnd.movY = wnd.movY[:0]
			wnd.counter = 0
		} else if keyCode == 9 { // F
			wnd.Props.Fullscreen = !wnd.Props.Fullscreen
		} else if keyCode == 12 { // I
			fmt.Println(wnd.Stats.FPS, "FPS  ", wnd.Stats.UPS, "UPS  ", wnd.counter, "entities")
		} else if keyCode == 13 { // J
			wnd.mipMapOn = !wnd.mipMapOn
			texCount := len(imgNames)
			if wnd.mipMapOn {
				for _, rect := range wnd.entities {
					rect.TexRef += texCount
				}
			} else {
				for _, rect := range wnd.entities {
					rect.TexRef -= texCount
				}
			}
		} else if keyCode == 16 { // M
			if len(wnd.entities) > 0 {
				wnd.moving = !wnd.moving
			}
		} else if keyCode == 18 { // O
			scale = 0.125
			wnd.incX = 0
			wnd.incY = 0
			for i, rect := range wnd.entities {
				imgIndex := wnd.entImgIdxs[i]
				rect.Width = float32(imgWidths[imgIndex]) * scale
				rect.Height = float32(imgHeights[imgIndex]) * scale
				rect.RotX = rect.Width / 2
				rect.RotY = rect.Height / 2
			}
		} else if keyCode == 21 { // R
			if len(wnd.entities) > 0 {
				wnd.rotating = !wnd.rotating
			}
		} else if keyCode == 25 { // V
			wnd.Gfx.VSync = !wnd.Gfx.VSync
		} else if keyCode == 41 { // ESC
			wnd.Close()
		} else {
			println("key", keyCode)
		}
	}
	return nil
}

func (wnd *demoWindow) spawn(n int) {
	clW, clH := float32(wnd.Props.ClientWidth), float32(wnd.Props.ClientHeight)
	firstUpdate := bool(len(wnd.entities) == 0)
	texCount := len(imgNames)
	for i := 0; i < n; i++ {
		rndA, rndB, rndC, rndD := random.Float32(), random.Float32(), random.Float32(), random.Float32()
		imgIndex := int(rndA * 5)
		imgW, imgH := imgWidths[imgIndex], imgHeights[imgIndex]
		incX := imgIncX[imgIndex]
		incY := imgIncY[imgIndex]
		rectW, rectH := float32(imgW)*scale+wnd.incX*incX, float32(imgH)*scale+wnd.incY*incY
		rect := wnd.layer0.NewEntity()
		rect.X, rect.Y, rect.Width, rect.Height = rndB*(clW-rectW-2*padding)+padding, rndC*(clH-rectH-2*padding)+padding, rectW, rectH
		rect.TexRef, rect.TexX, rect.TexY, rect.TexWidth, rect.TexHeight = imgIndex, 0, 0, imgW, imgH
		if wnd.mipMapOn {
			rect.TexRef += texCount
		}
		rect.RotX, rect.RotY = rect.Width/2, rect.Height/2
		if random.Float32() < 0.5 {
			wnd.movX = append(wnd.movX, random.Float32()*0.5+0.3)
		} else {
			wnd.movX = append(wnd.movX, -random.Float32()*0.5+0.3)
		}
		if random.Float32() < 0.5 {
			wnd.movY = append(wnd.movY, random.Float32()*0.5+0.3)
		} else {
			wnd.movY = append(wnd.movY, -random.Float32()*0.5+0.3)
		}
		wnd.entities = append(wnd.entities, rect)
		wnd.entImgIdxs = append(wnd.entImgIdxs, imgIndex)
		wnd.rotSpeed = append(wnd.rotSpeed, rndD*0.25-0.125)
	}
	wnd.counter += n
	if firstUpdate {
		wnd.Update()
	}
}

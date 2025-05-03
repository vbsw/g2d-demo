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
	"github.com/vbsw/keycode"
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
	rotating     bool
	rotSpeed     []float32
	moving       bool
	movX, movY   []float32
	incX, incY   float32
	mipMapOn     int
	showRects    bool
	entFirstMov  bool
	shiftDown    bool
}

func newDemoWindow() *demoWindow {
	wnd := new(demoWindow)
	wnd.id = windowCounter
	wnd.title = fmt.Sprintf("g2d Demo - %d", wnd.id)
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
		wnd.Gfx.LoadTexture(newChibiTexture(i, imgNames[i]))
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
	if texture.Id() < len(imgNames)*2 {
		// load as a mipmap
		// (normaly you wouldn't load a texture twice; this is for tests, only)
		wnd.Gfx.LoadTexture(texture)
	}
	if texture.Id() < len(imgNames) {
		wnd.layer0.UseTexture(texture.Id(), texture.Id())
	}
	return nil
}

func (wnd *demoWindow) OnUpdate() error {
	if wnd.entFirstMov {
		wnd.Stats.DeltaTime = 0
	}
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
				wnd.movX[i] = -1 * wnd.movX[i]
			} else if rect.X > float32(wnd.Props.ClientWidth)-(rect.Width+padding) && wnd.movX[i] > 0 {
				wnd.movX[i] = -1 * wnd.movX[i]
			}
			if rect.Y < padding && wnd.movY[i] < 0 {
				wnd.movY[i] = -1 * wnd.movY[i]
			} else if rect.Y > float32(wnd.Props.ClientHeight)-(rect.Height+padding) && wnd.movY[i] > 0 {
				wnd.movY[i] = -1 * wnd.movY[i]
			}
			rect.X, rect.Y = wnd.movX[i]*float32(wnd.Stats.DeltaTime)*speed+rect.X, wnd.movY[i]*float32(wnd.Stats.DeltaTime)*speed+rect.Y
		}
	}
	if len(wnd.entities) > 0 {
		wnd.entFirstMov = false
		wnd.Update()
	}
	return nil
}

func (wnd *demoWindow) OnResize() error {
	clW, clH := float32(wnd.Props.ClientWidth), float32(wnd.Props.ClientHeight)
	// speed up moving (when out of screen)
	for i, rect := range wnd.entities {
		if rect.X+rect.Width > clW {
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
		if rect.Y+rect.Height > clH {
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

func (wnd *demoWindow) OnKeyDown(key int, repeated uint) error {
	if key == keycode.K {
		for i := range wnd.entities {
			wnd.movX[i] /= 1.1
			wnd.movY[i] /= 1.1
		}
	} else if key == keycode.L {
		for i := range wnd.entities {
			wnd.movX[i] *= 1.1
			wnd.movY[i] *= 1.1
		}
	} else if key == keycode.Q {
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
	} else if key == keycode.W {
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
	} else if key == keycode.One {
		wnd.spawn(1)
	} else if key == keycode.Two {
		wnd.spawn(10)
	} else if key == keycode.Three {
		wnd.spawn(100)
	} else if key == keycode.Four {
		wnd.spawn(1000)
	} else if key == keycode.Five {
		wnd.spawn(10000)
	} else if repeated == 0 {
		if key == keycode.B {
			wnd.Props.Borderless = !wnd.Props.Borderless
		} else if key == keycode.D {
			wnd.Props.Dragable = !wnd.Props.Dragable
		} else if key == keycode.A {
			scale /= 2
			for i, rect := range wnd.entities {
				imgIndex := wnd.entImgIdxs[i]
				rect.Width = float32(imgWidths[imgIndex]) * scale
				rect.Height = float32(imgHeights[imgIndex]) * scale
				rect.RotX = rect.Width / 2
				rect.RotY = rect.Height / 2
			}
		} else if key == keycode.S {
			scale *= 2
			for i, rect := range wnd.entities {
				imgIndex := wnd.entImgIdxs[i]
				rect.Width = float32(imgWidths[imgIndex]) * scale
				rect.Height = float32(imgHeights[imgIndex]) * scale
				rect.RotX = rect.Width / 2
				rect.RotY = rect.Height / 2
			}
		} else if key == keycode.C {
			for _, rect := range wnd.entities {
				wnd.layer0.Release(rect)
			}
			wnd.entities = wnd.entities[:0]
			wnd.entImgIdxs = wnd.entImgIdxs[:0]
			wnd.rotSpeed = wnd.rotSpeed[:0]
			wnd.movX = wnd.movX[:0]
			wnd.movY = wnd.movY[:0]
			wnd.counter = 0
			wnd.Stats.FPS, wnd.Stats.UPS = 0, 0
		} else if key == keycode.F {
			wnd.Props.Fullscreen = !wnd.Props.Fullscreen
		} else if key == keycode.G {
			wnd.showRects = !wnd.showRects
			if wnd.showRects {
				for _, rect := range wnd.entities {
					rect.TexRef = -1
				}
			} else {
				for i, rect := range wnd.entities {
					rect.TexRef = wnd.entImgIdxs[i]
				}
			}
		} else if key == keycode.N {
			wnd.Show(newDemoWindow())
		} else if key == keycode.I {
			if wnd.shiftDown {
				fmt.Println("")
				fmt.Println("------ Video Card ------")
				fmt.Println(fmt.Sprintf("%-6d MaxTexSize", g2d.MaxTexSize))
				fmt.Println(fmt.Sprintf("%-6d MaxTexUnits", g2d.MaxTexUnits))
				fmt.Println(fmt.Sprintf("%-6d MaxTextures", g2d.MaxTextures))
				fmt.Println(fmt.Sprintf("%-6t V-Sync", g2d.VSyncAvailable))
				fmt.Println(fmt.Sprintf("%-6t AV-Sync", g2d.AVSyncAvailable))
				currWnd = -1
				infoCount = 0
			} else {
				if len(wnd.entities) == 0 {
					wnd.Stats.FPS, wnd.Stats.UPS = 0, 0
				}
				if currWnd != wnd.id {
					currWnd = wnd.id
					fmt.Println("")
					fmt.Println(fmt.Sprintf("------ window %d ------", wnd.id))
					fmt.Println("FPS    UPS       entities")
				} else if infoCount%10 == 0 {
					fmt.Println("")
					fmt.Println("FPS    UPS       entities")
				}
				fmt.Println(fmt.Sprintf("%-4d   %-7d   %-8d", wnd.Stats.FPS, wnd.Stats.UPS, wnd.counter))
				infoCount++
			}
		} else if key == keycode.J {
			texCount := len(imgNames)
			wnd.mipMapOn = (wnd.mipMapOn + 1) % 3
			texOffset := texCount * wnd.mipMapOn
			for i := 0; i < texCount; i++ {
				wnd.layer0.UseTexture(i, i+texOffset)
			}
		} else if key == keycode.M {
			if len(wnd.entities) > 0 {
				wnd.moving = !wnd.moving
			}
		} else if key == keycode.O {
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
		} else if key == keycode.R {
			if len(wnd.entities) > 0 {
				wnd.rotating = !wnd.rotating
			}
		} else if key == keycode.V {
			wnd.Gfx.VSync = !wnd.Gfx.VSync
		} else if key == keycode.Escape {
			wnd.Close()
		} else if key == keycode.LeftShift {
			wnd.shiftDown = true
		} else {
			println("key", key)
		}
	}
	return nil
}

func (wnd *demoWindow) OnKeyUp(key int) error {
	if key == keycode.LeftShift {
		wnd.shiftDown = false
	}
	return nil
}

func (wnd *demoWindow) spawn(n int) {
	clW, clH := float32(wnd.Props.ClientWidth), float32(wnd.Props.ClientHeight)
	wnd.entFirstMov = len(wnd.entities) == 0
	for i := 0; i < n; i++ {
		rndA, rndB := random.Float32(), random.Float32()
		imgIndex := int(random.Float32() * 5)
		imgW, imgH := imgWidths[imgIndex], imgHeights[imgIndex]
		incX := imgIncX[imgIndex]
		incY := imgIncY[imgIndex]
		rectW, rectH := float32(imgW)*scale+wnd.incX*incX, float32(imgH)*scale+wnd.incY*incY
		rect := wnd.layer0.NewEntity()
		rect.X, rect.Y, rect.Width, rect.Height = rndA*(clW-rectW-2*padding)+padding, rndB*(clH-rectH-2*padding)+padding, rectW, rectH
		rect.R, rect.G, rect.B, rect.A = random.Float32(), random.Float32(), random.Float32(), random.Float32()*0.49+0.49
		rect.TexRef, rect.TexX, rect.TexY, rect.TexWidth, rect.TexHeight = imgIndex, 0, 0, imgW, imgH
		rect.RotX, rect.RotY = rect.Width/2, rect.Height/2
		if random.Float32() < 0.5 {
			wnd.movX = append(wnd.movX, random.Float32()*0.5+0.15)
		} else {
			wnd.movX = append(wnd.movX, -random.Float32()*0.5-0.15)
		}
		if random.Float32() < 0.5 {
			wnd.movY = append(wnd.movY, random.Float32()*0.5+0.15)
		} else {
			wnd.movY = append(wnd.movY, -random.Float32()*0.5-0.15)
		}
		wnd.entities = append(wnd.entities, rect)
		wnd.entImgIdxs = append(wnd.entImgIdxs, imgIndex)
		wnd.rotSpeed = append(wnd.rotSpeed, random.Float32()*0.25-0.125)
		if wnd.showRects {
			rect.TexRef = -1
		}
	}
	wnd.counter += n
	if wnd.entFirstMov {
		wnd.Stats.FPS, wnd.Stats.UPS = 0, 0
		wnd.Update()
	}
}

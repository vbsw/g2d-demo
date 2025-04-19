/*
 *          Copyright 2025, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"github.com/vbsw/g2d"
)

type textureLoader struct {
	fileName string
	id       int
	width    int
	height   int
	bytes    []byte
	mipMap   bool
}

func newTextureLoader(id int, fileName string) *textureLoader {
	texture := new(textureLoader)
	texture.id = id
	texture.fileName = fileName
	return texture
}

func (texture *textureLoader) RGBABytes() ([]byte, error) {
	if len(texture.bytes) == 0 {
		img, err := imageFromEmbededPNG(texture.fileName)
		if err == nil {
			texture.bytes = g2d.BytesFromImage(img)
			texture.width = img.Bounds().Max.X - img.Bounds().Min.X
			texture.height = img.Bounds().Max.Y - img.Bounds().Min.Y
			return texture.bytes, nil
		}
		return nil, err
	}
	texture.id += len(imgNames)
	texture.mipMap = true
	return texture.bytes, nil
}

func (texture *textureLoader) Id() int {
	return texture.id
}

func (texture *textureLoader) Dimensions() (int, int) {
	return texture.width, texture.height
}

func (texture *textureLoader) GenMipMap() bool {
	return texture.mipMap
}

func (texture *textureLoader) IsMipMap() bool {
	return texture.mipMap
}

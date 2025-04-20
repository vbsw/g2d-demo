/*
 *          Copyright 2025, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"github.com/vbsw/g2d"
	"image"
)

type chibiTexture struct {
	fileName string
	id       int
	width    int
	height   int
	bytes    []byte
	pass     int
}

func newChibiTexture(id int, fileName string) *chibiTexture {
	texture := new(chibiTexture)
	texture.id = id
	texture.fileName = fileName
	return texture
}

func (texture *chibiTexture) RGBABytes() ([]byte, error) {
	var err error
	if texture.pass == 0 {
		var img image.Image
		img, err = imageFromEmbededPNG(texture.fileName)
		if err == nil {
			var mipMapFileName string
			mipMapIndex := texture.id
			texture.width = img.Bounds().Max.X - img.Bounds().Min.X
			texture.height = img.Bounds().Max.Y - img.Bounds().Min.Y
			texture.bytes = g2d.BytesFromImage(img)
			if mipMapIndex < 3 {
				mipMapFileName = "mipmap0.png"
			} else {
				mipMapFileName = "mipmap1.png"
				mipMapIndex -= 3
			}
			img, err = imageFromEmbededPNG(mipMapFileName)
			if err == nil {
				var size int
				bytes0 := g2d.BytesFromImage(img)
				for i, j := texture.width, texture.height; i > 0 && j > 0; i, j = i/2, j/2 {
					size += i * j * 4
				}
				bytes1 := make([]byte, size, size)
				copy(bytes1, texture.bytes)
				texture.bytes = bytes1
				bytes1 = bytes1[texture.width*texture.height*4:]
				levW, levH := texture.width/2, texture.height/2
				fw, fh := mipMapIndex%2, mipMapIndex/2
				from := levW*fw*4 + levH*fh*texture.width*4
				for ; levW > 0 && levH > 0; levW, levH = levW/2, levH/2 {
					size = levW * 4
					for i := 0; i < levH; i++ {
						to := from + size
						copy(bytes1, bytes0[from:to])
						bytes1 = bytes1[size:]
						from += texture.width * 4
					}
					from += size - fw*size/2 - fh*texture.width*levW*2
				}
			}
		}
	} else {
		texture.id += len(imgNames)
	}
	texture.pass++
	return texture.bytes, err
}

func (texture *chibiTexture) Id() int {
	return texture.id
}

func (texture *chibiTexture) Dimensions() (int, int) {
	return texture.width, texture.height
}

func (texture *chibiTexture) GenMipMap() bool {
	return texture.pass == 2
}

func (texture *chibiTexture) IsMipMap() bool {
	return texture.pass > 1
}

func (texture *chibiTexture) FilterLinear() bool {
	return true
}

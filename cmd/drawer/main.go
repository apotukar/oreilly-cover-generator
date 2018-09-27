/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"image/png"
	"log"
	"os"

	"github.com/nanmu42/orly"
	"golang.org/x/image/colornames"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	c := orly.Cover{
		Width:        500,
		Height:       700,
		PrimaryColor: colornames.Red,
	}
	img, err := c.Draw()
	if err != nil {
		return
	}

	f, err := os.Create("out.png")
	if err != nil {
		return
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return
	}
}

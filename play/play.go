// /home/krylon/go/src/github.com/blicero/rndplay/play/play.go
// -*- mode: go; coding: utf-8; -*-
// Created on 17. 04. 2021 by Benjamin Walkenhorst
// (c) 2021 Benjamin Walkenhorst
// Time-stamp: <2021-04-21 20:01:42 krylon>

// Package play implements the handling of the media player.
package play

import "github.com/blicero/krylib"

// Player is a wrapper around the media player.
type Player struct {
	Bin             string
	Args            []string
	Files           []string
	HandlesMultiple bool
	idx             int
}

// Play plays the media file(s).
func (p *Player) Play() error {
	return krylib.ErrNotImplemented
} // func (p *Player) Play() error

#!/usr/bin/env bash

set -xe

fyne bundle --package ui --prefix embeddedRsc -o ui/bundled.go res/icons/cars32.png
fyne bundle --package ui --append --prefix embeddedRsc -o ui/bundled.go res/icons/drivers32.png
fyne bundle --package ui --append --prefix embeddedRsc -o ui/bundled.go res/icons/standings32.png
fyne bundle --package ui --append --prefix embeddedRsc -o ui/bundled.go res/icons/seasons32.png

#!/usr/bin/env bash

set -xeo pipefail

fyne bundle res/icons/cars32.png > bundled.go
fyne bundle -append res/icons/drivers32.png >> bundled.go

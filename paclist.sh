#!/usr/bin/env bash

comm -23 <(pacman -Qeq | sort) <(pacman -Qgq xorg | sort)

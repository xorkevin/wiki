#!/usr/bin/env bash

comm -23 <(pacman -Qeq | sort) <(pacman -Qgq base base-devel xorg | sort)

# first install

- xorg, xorg-xinit
- i3: window manager
- yay
- stow
- LUKS
- ntpd: time and date
- htop: process monitor
- moreutils
- lm_sensors: temp monitor
- tmux
- zprezto: zsh config
- base16-shell
- powerline-fonts
- ttf-google-fonts-git: fonts
- alsa-utils, alsa-plugins: sound
- pulseaudio, paprefs, pavucontrol
- xbindkeys
- maim, slop: screenshots
- to reset keyboard: setxkbmap -layout us
- xclip
- ufw
- imagemagick
- graphicsmagick
- feh
- rxvt-unicode
- lxappearance
- cronie

## tweaks

`/etc/makepkg.conf`:

```conf
MAKEFLAGS="-j$(nproc)" # parallel jobs
COMPRESSXZ=(xz -c -z - --threads=0) # multithreaded compression
```

`/etc/pacman.conf`:

```conf
Color
ParallelDownloads = 8
[multilib]
Include = /etc/pacman.d/mirrorlist
```

`/etc/sysctl.d/99-swappiness.conf`:
```conf
vm.swappiness=8
```

SSD wear-leveling
```
systemctl enable fstrim.timer
```

disable mouse acceleration

`/etc/X11/xorg.conf.d/50-mouse-acceleration.conf`:

```conf
Section "InputClass"
	Identifier "My Mouse"
	Driver "libinput"
	MatchIsPointer "yes"
	Option "AccelProfile" "flat"
	Option "AccelSpeed" "0"
EndSection
```

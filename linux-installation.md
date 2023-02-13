# Installation

## Installation media

* Archlinux

    ```sh
    dd bs=4M if=path/to/archlinux-version-x86_64.iso of=/dev/sdx conv=fsync oflag=direct status=progress && sync
    ```

* Windows
    * woeusb-ng

## Archlinux installation

* Confirm system is in uefi boot mode

    ```sh
    ls /sys/firmware/efi/efivars
    ```

    The directory should exist with contents
* Check network connection

    ```sh
    ip link
    ip addr
    ping 8.8.8.8
    ```

    * If using wifi, use `iwctl`
    * `systemd-networkd` and `systemd-resolved` should work out of the box for
      ethernet
* Check the system clock

    ```sh
    timedatectl status
    ```

    * `systemd-timesyncd` is enabled by default for NTP
* Partition disks
    * List block devices

        ```sh
        fdisk -l
        ```

    * Create a 1 GiB EFI system partition, 4 GiB swap space, and a root
      directory partition with the remainder of the space

        ```sh
        cfdisk /dev/nvmexn1
        ```

        * Partition types:
            * EFI System
            * Linux swap
            * Linux filesystem
    * Confirm partitions with `lsblk`
* Format partitions
    * root partition

        ```sh
        mkfs.ext4 /dev/nvmexn1px
        ```

    * swap partition

        ```sh
        mkswap /dev/nvmexn1px
        ```

    * EFI system partition

        ```sh
        mkfs.fat -F 32 /dev/nvmexn1px
        ```

* Mount file systems
    * root partition

        ```sh
        mount /dev/nvmexn1px /mnt
        ```

    * EFI system partition

        ```sh
        mount --mkdir /dev/nvmexn1px /mnt/boot
        ```

    * Swap partition

        ```sh
        swapon /dev/nvmexn1px
        ```
* Bootstrap

    ```sh
    pacstrap -K /mnt base base-devel linux linux-firmware
    ```

* Gen fstab

    ```sh
    genfstab -U /mnt >> /mnt/etc/fstab
    ```

* Chroot

    ```sh
    arch-chroot /mnt
    ```

* Install packages
    * vim
    * amd-ucode
* Configure install
    * Timezone
        * Set timezone

            ```sh
            ln -sf /usr/share/zoneinfo/Region/City /etc/localtime
            ```

        * Generate `/etc/adjtime`

            ```sh
            hwclock --systohc
            ```

    * Locale
        * Edit `/etc/locale.gen` and uncomment wanted locales
        * Generate locales

            ```sh
            locale-gen
            ```

        * Create `/etc/locale.conf` with:

            ```
            LANG=en_US.UTF-8
            ```

    * Hostname
        * Create `/etc/hostname` with:

            ```
            myhostname
            ```

    * Initramfs
        * Not usually needed, since run on pacstrap

            ```sh
            mkinitcpio -P
            ```

    * Root password

        ```sh
        passwd
        ```

* Boot loader `systemd-boot`

    ```sh
    bootctl install
    ```

    * Manual update

        ```sh
        bootctl update
        ```

    * Automatic update

        ```sh
        systemd-boot-update.service
        ```

    * Configure

        ```
        esp/loader/loader.conf

        default      arch.conf
        timeout      3
        console-mode max
        editor       no
        ```

        ```
        esp/loader/entries/arch.conf

        title   Arch Linux
        linux   /vmlinuz-linux
        initrd  /amd-ucode.img
        initrd  /initramfs-linux.img
        options root="LABEL=arch_os" rw
        ```

        ```
        esp/loader/entries/arch-fallback.conf

        title   Arch Linux (fallback initramfs)
        linux   /vmlinuz-linux
        initrd  /amd-ucode.img
        initrd  /initramfs-linux-fallback.img
        options root="LABEL=arch_os" rw
        ```

* Reboot

    ```sh
    exit
    umount -R /mnt
    reboot
    ```

### Notes

* If failing to boot, the GPU might be too new. Try using the iGPU instead.

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

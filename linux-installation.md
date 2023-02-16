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
        mkfs.ext4 -L arch_os /dev/nvmexn1px
        ```

    * swap partition

        ```sh
        mkswap -L swp /dev/nvmexn1px
        ```

    * EFI system partition

        ```sh
        mkfs.fat -F 32 -n ESP /dev/nvmexn1px
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
* If forgetting to label the root fs, use `e2label /dev/device my-label`
* If forgetting to label the esp, use `fatlabel /dev/device MY_LABEL`
* If forgetting to label the swap partition, use `swaplabel -L my-label /dev/device`

## Archlinux post install

### Network

* Enable `systemd-networkd.service`

    ```sh
    `systemctl enable --now systemd-networkd.service`
    ping 8.8.8.8
    ```

* Add wired interface

    ```
    /etc/systemd/network/20-wired.network

    [Match]
    Name=enp1s0

    [Link]
    RequiredForOnline=no

    [Network]
    DHCP=yes
    ```

    * Use `ip link` to list network interfaces
    * `systemctl restart systemd-networkd.service` to refresh configuration
    * `ip addr` to show dhcp assigned ip address
    * Change systemd-networkd startup behavior to only wait for one interface

        ```
        /etc/systemd/system/systemd-networkd-wait-online.service.d/wait-for-only-one-interface.conf

        [Service]
        ExecStart=
        ExecStart=/usr/lib/systemd/systemd-networkd-wait-online --any
        ```

* Enable `systemd-resolved.service`

    ```sh
    systemctl enable --now systemd-resolved.service
    resolvectl query xorkevin.com
    ```

* Stub `/etc/resolv.conf` with `systemd-resolved` for programs that read from
  `systemd-resolved` directly such as GPG

    ```sh
    ln -rsf /run/systemd/resolve/stub-resolv.conf /etc/resolv.conf
    ```

* For configuring custom DNS and fallback DNS servers if not using DHCP
  provided DNS

    ```
    /etc/systemd/resolved.conf.d/dns_servers.conf

    [Resolve]
    DNS=127.0.0.1 ::1
    Domains=~.
    ```

    ```
    /etc/systemd/resolved.conf.d/fallback_dns.conf

    [Resolve]
    FallbackDNS=127.0.0.1 ::1
    ```

### Time

* Enable `systemd-timesyncd` for NTP

    ```sh
    systemctl enable --now systemd-timesyncd.service
    ```

    ```sh
    timedatectl status
    ```

    ```
    /etc/systemd/timesyncd.conf

    NTP=time1.google.com time2.google.com time3.google.com time4.google.com
    ```

### Drives

* SSD wear-leveling

    ```
    systemctl enable fstrim.timer
    ```

* Favor using RAM instead of swap

    ```sh
    sysctl -w vm.swappiness=8
    ```

    ```
    /etc/sysctl.d/99-swappiness.conf

    vm.swappiness=8
    ```

### Pacman

```
/etc/makepkg.conf

MAKEFLAGS="-j$(nproc)" # parallel jobs
COMPRESSXZ=(xz -c -z --threads=0 -) # use all threads for compression
COMPRESSZST=(zstd -c -z -q --threads=0 -) # use all threads for compression
```

```
/etc/pacman.conf

Color
ParallelDownloads = 8
[multilib]
Include = /etc/pacman.d/mirrorlist
```

## Add a user

* Add wheel group to sudoers

    ```
    /etc/sudoers

    %wheel ALL=(ALL:ALL) ALL
    ```

* Install `zsh`
* Add a root user

    ```sh
    useradd -m -G users,wheel -s /bin/zsh myusername
    passwd myusername
    ```

    * List groups with `getent group`

## Base packages

* Install packages
    * `pacman-contrib`
    * `man-db`
    * `man-pages`
    * `openssh`
    * `git`
    * `polkit`

## Xorg

* Install `nvidia`
    * If there is an issue with `nouveau`:
        * remove `kms` from `HOOKS` in `/etc/mkinitcpio.conf` and run
          `mkinitcpio -P`
        * add `nomodeset nouveau.modeset=0 nvidia_drm.modeset=1` kernel
          parameters to options in bootloader entry
* Install packages
    * `xorg`
    * `xorg-xinit`
    * `arandr`
    * `i3`
    * `dmenu`
    * `feh`
    * `lxqt-policykit`
    * `rxvt-unicode`
    * `urxvt-perls`
    * `stow`
    * `starship`
    * `tmux`
    * `xclip`
    * `neovim`
    * `python-pynvim`
    * `fzf`
    * `ttf-nerd-fonts-symbols-2048-em`
    * `ttf-hack`
    * `ttf-roboto`
    * `inter-font`
    * `noto-fonts`
    * `maim`
    * `lxappearance`
    * `arc-solid-gtk-theme`
    * `arc-icon-theme`
    * `pipewire`
    * `pipewire-alsa`
    * `pipewire-pulse`
    * `pipewire-jack`
    * `pipewire-docs`
    * `wireplumber`
    * `pavucontrol`
    * `pass`
    * `pass-otp`
    * `qrencode`
    * `firefox`
    * `mpv`
    * `mpd`
    * `ncmpcpp`
    * `rustup`
    * `bat`
    * `ripgrep`
    * `fd`
* Install aur helper

    ```sh
    rustup default stable
    git clone https://aur.archlinux.org/paru.git
    cd paru
    makepkg -si
    ```

* Install packages
    * `paru-bin`
    * `yay-bin`
    * `antibody-bin`

* Generate ssh keys

    ```sh
    systemctl --user enable --now ssh-agent.service
    ssh-keygen -f filename -t ed25519
    ```

    ```
    ~/.ssh/config.d/hosts

    Host hades
      HostName hades1
      User kevin
      IdentityFile ~/.ssh/id_ed25519
    Host github.com
      IdentityFile ~/.ssh/id_ed25519_github
    ```

* Add key to github
* Clone gpgkeys

* Disable mouse acceleration
    * list props

        ```sh
        xinput --list
        xinput --list-props <id>
        xinput --set-prop <id> 'libinput Accel Profile Enabled' 0, 1
        ```

    ```
    /etc/X11/xorg.conf.d/50-mouse-acceleration.conf

    Section "InputClass"
      Identifier "My Mouse"
      Driver "libinput"
      MatchIsPointer "yes"
      Option "AccelProfile" "flat"
      Option "AccelSpeed" "0"
    EndSection
    ```

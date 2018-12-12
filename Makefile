pacman-list=pacman-list.md

.PHONY: paclist

paclist:
	rm -f $(pacman-list)
	pacman -Qeq > $(pacman-list)

reflector:
	reflector --verbose --country 'United States' --age 12 --sort rate --save /etc/pacman.d/mirrorlist

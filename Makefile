pacman-list=pacman-list.md

.PHONY: paclist

paclist:
	rm -f $(pacman-list)
	pacman -Qeq > $(pacman-list)

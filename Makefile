pacman-list=pacman-list.md

.PHONY: paclist

paclist:
	rm -f $(pacman-list)
	./paclist.sh > $(pacman-list)

reflector:
	reflector --verbose --country 'United States' --age 3 --sort rate --save /etc/pacman.d/mirrorlist

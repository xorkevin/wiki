pacman-list=pacman-list.txt

.PHONY: paclist

paclist:
	rm -f $(pacman-list)
	./paclist.sh > $(pacman-list)

reflector:
	reflector --verbose --country 'United States' --latest 32 --protocol https,http --sort rate --save /etc/pacman.d/mirrorlist

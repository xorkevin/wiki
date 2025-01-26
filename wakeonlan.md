# Wake On LAN

## Enable in UEFI

Enable Wake On LAN setting in UEFI. It may be named "Power On By from PCI-E".

## Create link file

```
/etc/systemd/network/50-wired.link

[Match]
MACAddress=aa:bb:cc:dd:ee:ff

# Copy contents from /usr/lib/systemd/network/99-default.link
# or the file pointed to by Link File from networkctl status INTERFACE
[Link]
NamePolicy=keep kernel database onboard slot path
AlternativeNamesPolicy=database onboard slot path mac
MACAddressPolicy=persistent
# Add this
WakeOnLan=magic
```

Network settings will take effect on reboot.

## Debug

```sh
sudo ethtool INTERFACE
# Look for output `Wake-on:` which should be `g` for magic packet
```

## Trigger wake

```sh
wol MAC_ADDRESS
```

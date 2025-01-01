# Secure Boot

## Prerequisites

- Backup secure boot keys

  ```sh
  for var in PK KEK db dbx ; do efi-readvar -v $var -o old_${var}.esl ; done
  ```

- Go into UEFI and clear secure boot keys
- Install `sbctl`
- Confirm `sbctl status` shows that sbctl is not installed and secure boot is
  disabled
- `bootctl` should show that secure boot is `disabled (setup)`

## Create and enroll secure boot keys

- Create secure boot keys

  ```sh
  sbctl create-keys
  ```

- Enroll secure boot keys in UEFI

  ```sh
  sbctl enroll-keys -m
  ```

- `sbctl status` should show that sbctl is installed but secure boot will still
  be disabled until secure boot in the UEFI has been enabled

## Sign files

- Check files that need to be signed

  ```sh
  sbctl verify
  ```

- Sign those files

  ```sh
  sbctl sign -s PATH/TO/FILE

  cat unsigned.txt | xargs -I{} sh -c 'echo "{}" && sbctl sign -s "{}"'
  ```

- Secure boot should now work and after a reboot, can confirm with `sbctl status`

## Future signing

With `systemd-boot-update.service`, the boot loader is only updated after a
reboot, and the sbctl pacman hook will not sign the new file. As a workaround,
sign the boot loader directly in `/usr/lib/`, because `bootctl install` and
`bootctl update` will automatically recognize and copy `.efi.signed` files to
the ESP if present, instead of the normal `.efi` file. See bootctl(1).

```
sbctl sign -s -o /usr/lib/systemd/boot/efi/systemd-bootx64.efi.signed /usr/lib/systemd/boot/efi/systemd-bootx64.efi
```

# Pacoloco

/etc/pacoloco.yaml
```
port: 9129
cache_dir: /var/cache/pacoloco
download_timeout: 8 # seconds
purge_files_after: 2592000 # seconds
repos:
  archlinux:
    urls:
      - https://mirrors.kernel.org/archlinux
prefetch:
  cron: '0 0 16 * * * *'
  ttl_unaccessed_in_days: 30
  # It deletes and stop prefetch packages (and db links) when not downloaded after ttl_unaccessed_in_days days that it had been updated.
  ttl_unupdated_in_days: 300
```

/etc/pacman.d/mirrorlist
```
Server = http://172.16.115.40:9129/repo/archlinux/$repo/os/$arch
```

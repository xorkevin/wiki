# Dnsmasq

/etc/resolvconf.conf
```
resolv_conf=/etc/resolv.conf
name_servers="::1 127.0.0.1"
dnsmasq_conf=/etc/dnsmasq-conf.conf
dnsmasq_resolv=/etc/dnsmasq-resolv.conf
```

/etc/dnsmasq.conf
```
conf-file=/etc/dnsmasq-conf.conf
resolv-file=/etc/dnsmasq-resolv.conf
local=/dev.localhost/
address=/dev.localhost/127.0.0.1
```

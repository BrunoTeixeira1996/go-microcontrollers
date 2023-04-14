
# Digispark

- Download latest release from here https://github.com/micronucleus/micronucleus/releases 

```console
$ unzip ~/Downloads/micronucleus-cli-master-882e7b4a
$ sudo cp micronucleus-cli-master-882e7b4a/micronucleus /usr/local/bin/micronucleus
$ export PATH=$PATH:/usr/local/bin/micronucleus
$ tinygo flash -target=digispark main.go
INSERT USB STICK
```


# XIAO

- Create new target in `/usr/local/lib/tinygo/targets/xiao-new.json`

```console
brun0@pop-os:/usr/local/lib/tinygo/targets
$ cat xiao-new.json
{
    "inherits": ["atsamd21g18a"],
    "build-tags": ["xiao"],
    "serial": "usb",
    "serial-port": ["2886:802f"],
    "flash-1200-bps-reset": "true",
    "flash-method": "msd",
    "msd-volume-name": "Seeed XIAO",
    "msd-firmware-name": "firmware.uf2"
}
```

- Write your main.go program and flash like this

```console
tinygo build -o firmware.uf2 -target=xiao-new main.go && sudo mkdir -p /media/xiao && sudo mount /dev/sda /media/xiao && sudo cp firmware.uf2 /media/xiao
```

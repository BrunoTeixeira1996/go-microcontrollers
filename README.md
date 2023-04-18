# Digispark

## Flashing

Download latest release from here https://github.com/micronucleus/micronucleus/releases

```console
$ unzip ~/Downloads/micronucleus-cli-master-882e7b4a
$ sudo cp micronucleus-cli-master-882e7b4a/micronucleus /usr/local/bin/micronucleus
$ export PATH=$PATH:/usr/local/bin/micronucleus
$ tinygo flash -target=digispark digispark/digispark.go
INSERT USB STICK
```

## Note

For now this is not working because tinygo does not provied yet the HID for keyboard emulator


# XIAO

## Flashing

Looking at this issue https://github.com/tinygo-org/tinygo/issues/3639 we can see that tinygo flash does not work on Seeed XIAO because _as recently as Dec 2022, the Seeed Studio XIAO M0 dev boards shipped with a bootloader from Nov 2019: UF2 Bootloader v3.7.0-33-g90ff611-dirty SFHWRO_.
To find the new MSD (Mass Storage Device) just plug the device into your computer, open the device in your file manager and check the files present in it.
After getting the new MSD name, create a new target with it.

Create new target in `/usr/local/lib/tinygo/targets/xiao-new.json`

```console
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

After this, the flash way still does not work (at least for me) so we need to build and move `.uf2` file to XIAO device

The following shows an example to emulate a keyboard by opening xcalc on my device (using i3wm with dmenu)

```console
$ tinygo build -o firmware.uf2 -target=xiao-new main.go && sudo mkdir -p /media/xiao && sudo mount /dev/sda /media/xiao && sudo cp firmware.uf2 /media/xiao
```

In order to reset the device we need to double TAP de RST pins like this https://wiki.seeedstudio.com/Seeeduino-XIAO/#enter-bootloader-mode

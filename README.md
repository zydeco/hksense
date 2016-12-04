# hksense

This is a HomeKit bridge for the [Sense](https://hello.is) sleep tracking device using [HomeControl](https://github.com/brutella/hc).

The Sense device is published as a HomeKit accessory providing readings for:

 - Temperature
 - Humidity
 - Light Level
 - Air Quality

# Getting Started

1. [Install Go](http://golang.org/doc/install)
2. [Setup Go workspace](http://golang.org/doc/code.html#Organization)
3. Install

        cd $GOPATH/src
        
        # Clone project
        git clone https://github.com/zydeco/hksense && cd hksense

4. Log In

        go run hksense.go login
    
    You will be prompted for your username and password. Save the Access Token from the output.

5. Set environment and run

        export SENSE_ACCESS_TOKEN=3.c7acec88ba2311e6b5016c4008a8f70a
        go run hksense.go

# Configuration

hksense is configured through environment variables:

**Environment Variables**

Required

- `SENSE_ACCESS_TOKEN`: Set this to the access token you get when logging in

Optional

- `SENSE_HOMEKIT_PIN`: 8-digit PIN to use when pairing with HomeKit. Defaults to 12300123
- `SENSE_DEVICE_NAME`: Name the device will appear as in HomeKit
- `SENSE_REFRESH_INTERVAL`: How often to refresh measurements. Signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "60s", "1m30s" or "1h". Default and minimum is "1m"

# Cross-compiling

It's easy to cross-compile hksense to run on a Raspberry Pi or something similar:

    > GOARCH=arm GOARM=6 GOOS=linux go build -v github.com/zydeco/hksense
    > file hksense
    hksense: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked, not stripped


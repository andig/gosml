# gosml

[![Build Status](https://travis-ci.org/andig/gosml.svg?branch=master)](https://travis-ci.org/andig/gosml)

Go port of [volkszaehler/libsml](https://github.com/volkszaehler/libsml)

## Usage

To install run

    go get github.com/andig/gosml

`gosml` uses `dep` for vendor management. Make sure you have `dep` installed and and run `dep ensure` to fetch vendor dependencies.

To include in your code use

````
import(
	"github.com/andig/gosml"
)
````
to import the `sml` package.

## Example

For an example see [cmd/server](https://github.com/andig/gosml/blob/master/cmd/server/main.go) and the [libsml](https://github.com/volkszaehler/libsml) documentation.

## Status

The implementation of this port is not complete and has not been extensively tested.

It is only intended for parsing OBIS codes. This has been validated against
[libsml-testing](https://github.com/devZer0/libsml-testing):

    libsml-testing/DrNeuhaus_SMARTY_ix-130.bin
    129-129:199.130.3*255 44 4e 54
    1-0:0.0.9*255         09 01 44 4e 54 01 00 00 20 30
    1-0:2.8.0*255               4373.0 Wh
    1-0:2.8.1*255               4373.0 Wh
    1-0:2.8.2*255                  0.0 Wh
    1-0:15.7.0*255                 0.0 W
    129-129:199.130.5*255 fe a5 d6 bc 65 8d 73 46 73 f5 79 00 51 d6 1c 2e bd 55 34 4a 62 85 68 6d e3 b6 45 ee 01 04 d6 5f 5a bb a6 cb 68 f1 bf cb 7f cf 94 f9 9c d4 d8 45

    libsml-testing/EMH_eHZ-GW8E2A500AK2.bin
    129-129:199.130.3*255 45 4d 48
    1-0:0.0.0*255         30 32 32 38 30 38 31 36
    1-0:1.8.1*255           14798112.9 Wh
    1-0:1.8.2*255               2012.4 Wh
    0-0:96.1.255*255      30 30 30 32 32 38 30 38 31 36
    1-0:1.7.0*255                 13.8 W

    libsml-testing/EMH_eHZ-HW8E2A5L0EK2P.bin
    129-129:199.130.3*255 45 4d 48
    1-0:0.0.9*255         06 45 4d 48 01 02 71 53 c8 c6
    1-0:1.8.0*255            8391648.8 Wh
    1-0:1.8.1*255            8391447.6 Wh
    1-0:1.8.2*255                201.2 Wh
    1-0:15.7.0*255               163.5 W
    129-129:199.130.5*255 28 ab db f4 41 a5 ef 63 35 ec 38 bd 9c 0e 0a 79 03 0d a5 f4 a4 d7 39 5f 69 52 ef 7a 0b 0f f3 fc d1 37 f8 66 d1 14 15 4c f0 a9 52 42 de aa b1 7d

    libsml-testing/EMH_eHZ-HW8E2A5L0EK2P_1.bin
    129-129:199.130.3*255 45 4d 48
    1-0:0.0.9*255         06 45 4d 48 01 00 1d 46 15 ca
    1-0:1.8.0*255           15141809.0 Wh
    1-0:1.8.1*255           15141809.0 Wh
    1-0:1.8.2*255                  0.0 Wh
    1-0:15.7.0*255              1213.3 W
    129-129:199.130.5*255 2a 27 8d ad f3 a5 d8 01 1e ea 7c 1f 60 cf 20 2d 4e da da 88 98 7e b3 8e e0 f4 e4 ce 46 e8 6d 1a 2b e8 6b 2b 40 65 51 ad 6f 9e 93 67 aa d2 81 9d

    libsml-testing/EMH_eHZ-HW8E2AWL0EK2P.bin
    129-129:199.130.3*255 45 4d 48
    1-0:0.0.9*255         01 a8 15 98 64 80 02 01 02
    1-0:1.8.0*255            5378499.0 Wh
    1-0:1.8.1*255            5378499.0 Wh
    1-0:1.8.2*255                  0.0 Wh
    1-0:15.7.0*255               191.9 W
    129-129:199.130.5*255 c2 fb 28 83 40 2a d8 7c 9e a2 7a cc fd 04 28 20 6f bd 06 56 6b a7 95 7c 5e b0 de 50 54 a4 40 ab d5 5a 6d 94 d6 77 17 6f dd f8 05 c2 3f 8d ef 1e

    libsml-testing/EMH_eHZ-IW8E2A5L0EK2P_crash.bin
    129-129:199.130.3*255 45 4d 48
    1-0:0.0.9*255         06 45 4d 48 01 07 19 7c 24 56
    1-0:1.8.0*255            2795692.7 Wh
    1-0:1.8.1*255            2795692.7 Wh
    1-0:1.8.2*255                  0.0 Wh
    1-0:16.7.0*255               136.7 W
    129-129:199.130.5*255 8b 6a 0e 6e 12 f5 d9 80 f7 30 b6 bd 5e 19 41 83 4e b0 e4 3e 4a 63 23 d9 99 25 95 56 f5 e5 6e 04 04 98 c8 97 38 f0 f6 df f8 78 5b 04 5d 84 e0 d6
    1-0:96.50.2*4               6370.0
    1-0:96.50.2*6

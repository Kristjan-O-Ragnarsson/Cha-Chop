# Cha-Chop

A simple project to test the performance of ChaCha20-Poly1305 encryption using WolfSSL for Linux systems.

It has been tested on:
- Raspberry Pi running Raspberry Pi OS (6.6.31+rpt-rpi-v8 #1 SMP PREEMPT Debian)
- Intel N95 Mini PC running Debian (6.1.0-37-rt-amd64 #1 SMP PREEMPT_RT Debian)

## Installation

This project requires:
- `linux-perf`
- `WolfSSL`
- ``GNU Time``

### Building WolfSSL

Requirements
- `git`
- ``make``
- ``autoconf``
- ``libtool``
- ``build-essentials``

```shell
git submodulte init

git submodule update

cd wolfssl

git checkout 1c56a2674a5ef51c

cd ..

./install.sh
```

## Usage
```shell
./test.sh

./mem-test.sh
```

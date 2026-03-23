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

### Building TestCrypt-cli

Requirements
- `go`
```bash
cd testcrypt-cli

go build
```


## Usage

### Running preliminary tests (Bash scripts):
```shell
./test.sh

./mem-test.sh
```

### Running TestCrypt-cli
```shell
cd testcrypt-cli

./TestCryptLinux -r 100000 -path "$(cd .. && pwd)"
```

## Citation

If you find our work useful in your research, please consider citing:

```bibtex
@inprocedin{Ragnarsson2026IoTBDS,
    title     = {Performance Testing of ChaCha20-Poly1305 for Internet of Things and Industrial Control System Devices},
    booktitle = {11th International Conference on Internet of Things, Big Data and Security (IoTBDS 2026)},
    author    = {Ragnarsson, Kristjan Orri and Mallett, Jacky},
    year      = 2026,
}
```

Pre-print avaliable (Submitted Manuscript): https://arxiv.org/abs/2603.19150

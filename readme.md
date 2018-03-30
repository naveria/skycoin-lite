# Skycoin Liteclient

This repository contains a liteclient for Skycoin written in Go. At the moment it is only used to compile an iOS Framework

## Compiling

For the compilation process to iOS Framework, we use [Go Mobile](https://github.com/golang/mobile).

```bash
$ gomobile bind -target=ios github.com/naveria/skycoin-lite/mobile/mobile
```

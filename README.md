# Cryptography Week 3 — Hashing & Password Cracking Labs

**Student:** Chan Arch Vathanak

**Programming language:** Go


## Overview

This university cybersecurity lab explores deterministic hashing, the effect of a
one-character input change, dictionary-based password recovery, and regular-expression
flag verification. Each lab is an independent Go module. The original implementations
and recorded results are retained, with targeted corrections to module configuration,
LAB1's copied SHA1 settings, formatting, and error handling.

## Repository structure

```text
W3-Hash/
├── LAB0/
│   ├── LAB0.go
│   ├── go.mod
│   ├── go.sum
│   └── output.txt
├── LAB1/
│   ├── main.go
│   ├── go.mod
│   ├── utils/crack/md5.go
│   ├── nord_vpn.txt             
│   └── verbose.txt
├── LAB2/
│   ├── main.go
│   ├── go.mod
│   ├── utils/crack/sha1.go
│   ├── nord_vpn.txt         
│   └── verbose.txt
├── LAB3/
│   ├── main.go
│   ├── go.mod
│   ├── utils/crack/sha512.go
│   ├── nord_vpn.txt             
│   └── verbose.txt
├── LAB4/
│   ├── main.go
│   ├── main_test.go
│   ├── go.mod
│   └── output.txt
├── docs/VERIFICATION.md
├── screenshots/README.md
├── README.md
└── .gitignore
```


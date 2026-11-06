# Verification record

Verified on 2026-11-06 with `go version go1.22.2 linux/amd64`.

## Module checks

| Module | go fmt ./... | go test ./... | go vet ./... | go mod verify | go run . |
| --- | --- | --- | --- | --- | --- |
| LAB0 | Passed | Passed | Passed | Passed | Passed |
| LAB1 | Passed | Passed | Passed | Passed | Passed |
| LAB2 | Passed | Passed | Passed | Passed | Passed |
| LAB3 | Passed | Passed | Passed | Passed | Passed |
| LAB4 | Passed | Passed | Passed | Passed | Passed |

LAB0–LAB3 report `[no test files]`; their full default runs were verified separately.
LAB4 has eight table-driven regex cases. Dependency tidying also succeeded in every module.
Go commands ran with `GOTOOLCHAIN=local`, `GOWORK=off`, and `GOPROXY=off`, using
the existing cached dependencies and Go 1.22.2. This verifies compatibility without
silently downloading a newer toolchain.

## Runtime evidence

- LAB0: all ten digest values matched independent Python `hashlib` calculations.
- LAB1: 337 candidate hashes checked independently; original output reproduced byte for byte.
- LAB2: 593 candidate hashes checked independently; original output reproduced byte for byte.
- LAB3: 343 candidate hashes checked independently; original output reproduced byte for byte.
- LAB4: default output reproduced byte for byte; text clue and valid/invalid CLI candidates checked.
- LAB1–LAB3: separately built binaries exited with status 1 when run from a directory without the dictionary.
- Original wordlists and saved LAB1–LAB4 output files retained their SHA-256 checksums.

An initial explicit binary build encountered a VCS-status error before repository
initialization. After the requested root Git repository was initialized, the build
and all checks passed without disabling VCS stamping.

## Inspection and repairs

- LAB1 and LAB2 originally had byte-identical `main.go` files. LAB1 imported
  `lab2/utils/crack`, used the SHA1 target, and labeled MD5 output as SHA1. These
  copied settings were corrected using the existing MD5 transcript as evidence.
- LAB0 module name changed from `lab1` to `lab0`; SHA3 is a direct dependency.
- Each module declares Go 1.22. LAB4 received its own `go.mod`.
- Go sources were formatted. Dictionary I/O errors now return unsuccessful exit status.
- LAB4 retains its original flag/pattern and successful default output, with optional
  `-leak` and `-flag` arguments plus regex boundary tests.
- Removed only `LAB2/utils/crack/verbose.txt`, which was confirmed to be zero bytes.
- All three identical local dictionary copies were retained for independent lab execution.
- No compiled binaries, nested repositories, or screenshots existed in the original project.
- Added actual LAB0 output, repository documentation, screenshot instructions, and `.gitignore`.

## Original artifact checksums

These are the SHA-256 values of the original, preserved artifacts.

| Artifact | SHA-256 |
| --- | --- |
| `LAB1/nord_vpn.txt` | `1b57ef6acc09c2393328ccdcdd379b5792ec4c430a4729ac7bb5a83cd6bb825e` |
| `LAB1/verbose.txt` | `006d9b30a60acb6ba68eb701d204116573e22dde01ca2de2bb304f0d18e36476` |
| `LAB2/nord_vpn.txt` | `1b57ef6acc09c2393328ccdcdd379b5792ec4c430a4729ac7bb5a83cd6bb825e` |
| `LAB2/verbose.txt` | `512dfdd21ad557a88fba8c41363e5fb68a0e27ff65e1a437761a942e1a97d8a7` |
| `LAB3/nord_vpn.txt` | `1b57ef6acc09c2393328ccdcdd379b5792ec4c430a4729ac7bb5a83cd6bb825e` |
| `LAB3/verbose.txt` | `3d72f959b1d2d15190c0352213a2448c67b8d9f8de17ef0cc4684e5a8a026c15` |
| `LAB4/output.txt` | `bcd603cf14d7a80446289704319d6e6131633fed0f6262eb02f07b6924390077` |

The wordlist filename provides no verified upstream source or redistribution
permission; the dictionaries are excluded from Git. Existing output is retained
as explicitly requested lab evidence. No experimental timings or screenshots are claimed.

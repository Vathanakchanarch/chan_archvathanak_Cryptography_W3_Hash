# Cryptography Week 3 — Hashing & Password Cracking Labs

**Student:** Chan Arch Vathanak

**Programming language:** Go

**Requirements:** Go 1.22+ (verified with Go 1.22.2 on Linux)

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
│   ├── nord_vpn.txt             # local only; ignored by Git
│   └── verbose.txt
├── LAB2/
│   ├── main.go
│   ├── go.mod
│   ├── utils/crack/sha1.go
│   ├── nord_vpn.txt             # local only; ignored by Git
│   └── verbose.txt
├── LAB3/
│   ├── main.go
│   ├── go.mod
│   ├── utils/crack/sha512.go
│   ├── nord_vpn.txt             # local only; ignored by Git
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

## Labs

| Lab | Topic | Behavior |
| --- | --- | --- |
| LAB0 | Hash comparison | Hashes `hello` and `Hello` using MD5, SHA1, SHA256, SHA512, and SHA3-256, then compares each pair. |
| LAB1 | MD5 dictionary comparison | Hashes dictionary candidates until the configured MD5 target matches. |
| LAB2 | SHA1 dictionary comparison | Hashes dictionary candidates until the configured SHA1 target matches. |
| LAB3 | SHA512 dictionary comparison | Hashes dictionary candidates until the configured SHA512 target matches. |
| LAB4 | Mini-CTF cat clue and flag verification | Offers a text cat clue and checks an exact flag pattern with an anchored regular expression. |

## Setup

Install Go 1.22 or later, then check the installation:

```bash
go version
```

From a fresh clone, enter the repository directory:

```bash
git clone https://github.com/Vathanakchanarch/chan_archvathanak_Cryptography_W3_Hash.git
cd chan_archvathanak_Cryptography_W3_Hash
```

The existing local project directory may instead be named `W3-Hash`.
Run each lab from its own directory. There is no root module or workspace file.

LAB0 imports `golang.org/x/crypto/sha3`: Go 1.22 does not provide `crypto/sha3`
in its standard library. Its existing dependency versions are retained:
`golang.org/x/crypto v0.31.0` and indirect `golang.org/x/sys v0.28.0`.
Their module requirements support Go 1.22. Download and verify them with:

```bash
(cd LAB0 && go mod download && go mod verify)
```

LAB1–LAB4 use the standard library only. LAB1–LAB3 additionally need an authorized
copy of the original dictionary named `nord_vpn.txt` in each lab directory.
The existing local copies are preserved but are excluded from Git. A fresh clone
will compile without them; running a dictionary lab requires supplying the file.

## Run each lab

Run these commands from the repository root. Parentheses keep each directory change
local to its command.

### LAB0 — Hash comparison

```bash
(cd LAB0 && go run .)
```

The two inputs are defined in `LAB0/LAB0.go`. The verified run is saved in
[LAB0/output.txt](LAB0/output.txt).

### LAB1 — MD5

```bash
(cd LAB1 && go run .)
(cd LAB1 && set -o pipefail && go run . | tee verbose.txt)
```

### LAB2 — SHA1

```bash
(cd LAB2 && go run .)
(cd LAB2 && set -o pipefail && go run . | tee verbose.txt)
```

### LAB3 — SHA512

```bash
(cd LAB3 && go run .)
(cd LAB3 && set -o pipefail && go run . | tee verbose.txt)
```

In LAB1–LAB3, `go run . | tee verbose.txt` displays every attempted candidate and
saves the same output. `tee` replaces an existing file; copy any prior result you
want to retain before rerunning this command. `pipefail` is supported by Bash and
Zsh and makes execution errors visible through the pipeline's exit status.
The committed transcripts were preserved byte for byte and reproduced during
verification. File-open and dictionary-read errors print to standard error and
exit unsuccessfully. An exhausted dictionary prints `Password not found!`.

### LAB4 — Mini-CTF

```bash
(cd LAB4 && go run .)
(cd LAB4 && go run . -leak -flag 'cryptoCTF{meowmeow}')
(cd LAB4 && go run . -flag 'cryptoCTF{meow}')
```

The default run preserves the original successful flag demonstration. `-leak`
prints `Cat leak: meowmeow`, a text clue derived from the existing verifier.
`-flag` supplies a candidate; the final example deliberately fails with a nonzero
exit status. The pattern `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$` accepts exactly
`cryptoCTF{meowmeow}`. The original project supplied the regex and successful output;
no separate cat image or leak asset was supplied.

## Hashing algorithms

| Algorithm | Digest size | Go package |
| --- | --- | --- |
| MD5 | 128 bits | `crypto/md5` |
| SHA1 | 160 bits | `crypto/sha1` |
| SHA256 | 256 bits | `crypto/sha256` |
| SHA512 | 512 bits | `crypto/sha512` |
| SHA3-256 | 256 bits | `golang.org/x/crypto/sha3` |

The dictionary labs compare lowercase hexadecimal digests of unsalted candidate
strings. MD5 and SHA1 are legacy algorithms with known collision weaknesses.
These fast, unsalted hashes are unsuitable for storing real passwords; the labs
use them to demonstrate why dictionary attacks work.

## Verified results

These observations come from actual execution with the existing local dictionary.
Candidate counts include the matching candidate and exclude the final status line.

| Lab | Observed result | Saved evidence |
| --- | --- | --- |
| LAB0 | All five algorithms produced different digests for `hello` and `Hello`; all ten digests were independently checked. | [output.txt](LAB0/output.txt) |
| LAB1 | MD5 target matched `Google101` after 337 candidates. | [verbose.txt](LAB1/verbose.txt) |
| LAB2 | SHA1 target matched `Princess1` after 593 candidates. | [verbose.txt](LAB2/verbose.txt) |
| LAB3 | SHA512 target matched `Gutter11` after 343 candidates. | [verbose.txt](LAB3/verbose.txt) |
| LAB4 | Printed `Correct flag!` for `cryptoCTF{meowmeow}`; incorrect candidates were rejected. | [output.txt](LAB4/output.txt) |

LAB1's source originally contained LAB2's SHA1 target and import. Its MD5 target
was restored from the existing MD5 transcript, then independently checked by
execution. No cracking results or timings were invented.

## Verification

Run checks separately inside every module:

```bash
for lab in LAB0 LAB1 LAB2 LAB3 LAB4; do
    (cd "$lab" && go fmt ./... && go test ./... && go vet ./... && go mod verify) || exit 1
done
```

All five modules passed on Go 1.22.2. LAB4 includes tests for exact matching,
case sensitivity, repetition count, and rejection of prefixes, suffixes, and
trailing newlines. LAB0–LAB3 have no Go unit-test files; their execution and output
were checked separately. See [verification details](docs/VERIFICATION.md).

## Wordlist source and attribution

The original project supplied three identical local files named `nord_vpn.txt`.
No upstream URL, author, license, or redistribution permission was included, so
the original publisher cannot be verified from the available files. The filename
alone does not establish NordVPN as its author or publisher.

The dictionaries remain on the student's machine and are excluded from this
repository. Obtain the course-provided or otherwise authorized original list
through its legitimate source. Record that source and its attribution/license
before redistributing it. Different dictionaries or ordering can change whether
a match is found and how many candidates are attempted.

## Screenshots

The [screenshots directory](screenshots/README.md) is ready for genuine captures
of each lab's terminal output or VS Code session. No screenshots were supplied
with the original project, and none have been fabricated. The text output files
above provide the verified execution evidence currently available.

## Educational security disclaimer

This repository is for university coursework, authorized security testing, and
CTF practice. Run password-recovery experiments only on data you own or have
explicit permission to test. Do not use the examples against other people's
accounts, systems, or credentials. The recovered strings and CTF flag shown here
are the supplied lab examples.

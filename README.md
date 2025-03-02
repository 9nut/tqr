# tqr: QR code in term
Generates QR code using fancy Unicode chars.

## Usage
`tqr -l [LMQH] -b <black dot character> -w <white dot character> [<string to encode>]`

Encodes a single string argument or the contents of stdin, using the most
excellent `rsc.io/qr` package, then outputs it to the terminal using
Unicode characters for black and white dots. All
flags are optional. The `-l` controls the error correction redundancy
level. The `-c` and the `-w` flags control the Unicode string to use for
black and white dots respectively. For help run `tqr` without arguments.
Party on ...

## Examples
Encode URLs:

```shell
$ echo -n 'https://www.9netics.com' | tqr
# OR
$ tqr 'https://www.9netics.com'
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️
⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️
⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️
⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬜️
⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬜️
⬜️⬜️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️
⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬜️
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬜️⬛️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️
```
or a colorful and topical, although less useful version:
```shell
$ ./tqr -w 💀 -b 🎃 'https://www.9netics.com'  
💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀
💀🎃🎃🎃🎃🎃🎃🎃💀💀🎃💀💀🎃💀🎃🎃💀💀🎃🎃🎃🎃🎃🎃🎃💀
💀🎃💀💀💀💀💀🎃💀💀💀🎃🎃🎃💀💀🎃🎃💀🎃💀💀💀💀💀🎃💀
💀🎃💀🎃🎃🎃💀🎃💀🎃🎃🎃💀🎃💀💀💀🎃💀🎃💀🎃🎃🎃💀🎃💀
💀🎃💀🎃🎃🎃💀🎃💀💀🎃🎃🎃💀💀🎃💀💀💀🎃💀🎃🎃🎃💀🎃💀
💀🎃💀🎃🎃🎃💀🎃💀💀💀🎃💀💀💀💀🎃💀💀🎃💀🎃🎃🎃💀🎃💀
💀🎃💀💀💀💀💀🎃💀💀🎃💀💀💀💀💀🎃💀💀🎃💀💀💀💀💀🎃💀
💀🎃🎃🎃🎃🎃🎃🎃💀🎃💀🎃💀🎃💀🎃💀🎃💀🎃🎃🎃🎃🎃🎃🎃💀
💀💀💀💀💀💀💀💀💀🎃🎃🎃💀🎃🎃🎃🎃🎃💀💀💀💀💀💀💀💀💀
💀🎃🎃🎃💀🎃🎃🎃🎃🎃💀🎃🎃💀🎃🎃💀💀🎃🎃💀💀💀🎃💀💀💀
💀🎃💀💀🎃🎃💀💀💀💀💀🎃🎃💀💀💀💀🎃💀🎃💀💀💀💀💀🎃💀
💀💀💀💀🎃💀🎃🎃💀💀🎃💀💀💀🎃💀💀💀🎃🎃🎃💀💀🎃🎃🎃💀
💀🎃🎃🎃🎃🎃🎃💀💀💀💀💀🎃💀🎃💀🎃🎃💀🎃💀💀💀💀🎃💀💀
💀🎃💀🎃🎃🎃💀🎃💀🎃💀💀💀🎃🎃💀💀🎃🎃🎃🎃💀🎃💀🎃🎃💀
💀💀🎃💀🎃🎃🎃💀🎃🎃💀💀🎃🎃🎃💀💀🎃💀🎃💀💀🎃💀💀🎃💀
💀🎃💀🎃🎃🎃💀🎃🎃💀💀🎃🎃🎃🎃🎃💀🎃🎃🎃🎃💀💀🎃🎃🎃💀
💀💀🎃🎃💀💀🎃💀💀💀🎃💀💀🎃🎃🎃💀🎃💀🎃💀🎃💀💀🎃💀💀
💀🎃💀💀🎃💀💀🎃🎃💀💀💀🎃💀🎃💀🎃🎃🎃🎃🎃🎃🎃💀💀💀💀
💀💀💀💀💀💀💀💀💀🎃💀🎃🎃💀💀💀🎃🎃💀💀💀🎃🎃💀🎃🎃💀
💀🎃🎃🎃🎃🎃🎃🎃💀🎃🎃💀💀💀🎃💀🎃🎃💀🎃💀🎃🎃💀🎃🎃💀
💀🎃💀💀💀💀💀🎃💀🎃🎃🎃🎃💀🎃🎃💀🎃💀💀💀🎃🎃💀💀🎃💀
💀🎃💀🎃🎃🎃💀🎃💀🎃🎃💀💀🎃🎃💀💀🎃🎃🎃🎃🎃🎃💀🎃🎃💀
💀🎃💀🎃🎃🎃💀🎃💀💀💀💀🎃🎃💀💀💀💀🎃💀🎃🎃🎃🎃💀💀💀
💀🎃💀🎃🎃🎃💀🎃💀🎃🎃🎃🎃🎃💀💀💀💀💀💀💀🎃💀💀💀🎃💀
💀🎃💀💀💀💀💀🎃💀🎃💀💀💀🎃🎃🎃💀🎃💀🎃💀🎃🎃💀🎃💀💀
💀🎃🎃🎃🎃🎃🎃🎃💀🎃💀💀🎃🎃💀🎃💀🎃💀💀🎃💀💀💀🎃🎃💀
💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀
```
To generate a MeCARD
```shell
$ tqr 'MECARD:N:Doe,John;TEL:13035551212;EMAIL:john.doe@example.com;;'
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️
⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️
⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬜️
⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️
⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬜️
⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬜️
⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬜️
⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬛️⬛️⬛️⬜️
⬜️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬜️
⬜️⬛️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬜️
⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️
⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️
⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬜️
⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬜️
⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️
⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬜️
⬜️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬜️
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬜️⬜️⬛️⬜️
⬜️⬛️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬛️⬛️⬜️⬛️⬜️⬛️⬜️
⬜️⬛️⬜️⬜️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬜️⬛️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬜️⬛️⬜️⬜️⬛️⬜️⬜️
⬜️⬛️⬛️⬛️⬛️⬛️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬜️⬜️⬜️⬛️⬛️⬜️⬛️⬛️⬜️⬜️⬛️⬛️⬛️⬛️⬛️⬜️⬜️⬜️⬛️⬛️⬜️
⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️
```

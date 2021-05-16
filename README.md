# pls

A utility cli tool with a grab bag of commands.

## Commands

| Command | Description |
|------|-------------|
| downcase | Downcases whatever follows the command |
| sha256 | Gets the SHA256 digest of the string following the command |
| upcase | Upcases whatever follows the command |
| uuid | Returns a random UUID v4 |
| version | Returns the version of pls |

## Examples

```console
$ pls uuid
f1781c30-3259-4d2f-823f-57d6487f6ca2
```

```console
$ pls downcase THIS should be downCASED
this should be downcased
```

Like most Unix-like commands the output of pls can be piped to other places.

```console
$ pls uuid >> output.txt
$ pls uuid >> output.txt
$ pls uuid >> output.txt
```

```console
$ cat output.txt
2d207514-84d8-424b-8e83-1114a0a75514
b4215934-eb84-44b2-9aff-3176b98e18d1
7498efc0-61e4-4ffc-b85e-0f77d7c589fc
```


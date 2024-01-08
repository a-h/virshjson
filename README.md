# virsh-json

Converts `virsh` output into a JSON table.

## Examples

```bash
virsh list | virsh-json
```

```json
[
 {
  "Id": "7",
  "Name": "nix-visor",
  "State": "running"
 }
]
```

## Installation

### Go

```bash
go install github.com/a-h/virsh-json/cmd/virsh-json@latest
```

### Nix

```bash
nix shell github:a-h/virsh-json
```

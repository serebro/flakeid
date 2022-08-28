# FlakeID

FlakeID is a CLI tool to generate and parse Snowflake IDs in base36 encoding.

## Installation

### From source code

```
go get -u github.com/serebro/flakeid
```

### Build

```
make build
```

## Examples

```
$ flakeid
MLCKDZGNGN4
```

```
$ flakeid -n 3 -m 1023                           
MLBX46XNNK0
MLBX46XNNK1
MLBX46XNNK2
```

```
$ flakeid -p MLBX46XNNK2

ID 82601872124276738
MACHINE_ID 1023
SEQUENCE 2
TIMESTAMP 1660689020982
DATE_TIME 2022-08-16T22:30:20.982Z
```

## Usage

```
$ flakeid -h

FlakeID - Generate Snowflake IDs
Version: 1.0.0
Usage: flakeid [OPTIONS]
  -l    To lower-case. Ex: mlbx46xnnk2
  -m int
        Machine ID. From 0 to 1023 (default 0)
  -n int
        Number of IDs to generate (default 1)
  -p string
        Parse Flake ID. Ex: flakeid -p MLBX46XNNK2
  -s string
        Epoch start time (default "2022-01-01T00:00:00Z")
```

## License

MIT

# GO-Port-scanner

## CLI port scanner made in GOlang

### Usage and options:

```bash
go run main.go --help-long

usage: PortScanner.exe [<flags>] <command> [<args> ...]

Flags:
      --help  Show context-sensitive help (also try --help-long and --help-man).
  -a, --all   Display all results. By default it will only display the open
              ports

Commands:
  help [<command>...]
    Show help.


  scan [<flags>] <target> [<ports>...]
    Scans ports on a specific target

    -s, --specific      Scans only a few specific ports that you specifed
    -p, --protocol=tcp  What protocol you want to use. Default is set to tcp
        --timeout=10s   Set the connection timeout. Amount of seconds
```

## Real use examples

```bash
go run PortScanner.go -a scan 192.168.0.10 -s 22, 80, 8080

** Target: 192.168.0.10 **
Scanning specific ports: [22 80 8080]
0 - 192.168.0.10:22 - Open
1 - 192.168.0.10:80 - Closed
2 - 192.168.0.10:8080 - Closed
```

```bash
go run PortScanner.go scan 192.168.0.10 --timeout=10s --protocol=tcp

** Target: 192.168.0.10 **
192.168.0.10:22 - Open
```

## GO mogule dependencies

```bash
go get gopkg.in/alecthomas/kingpin.v2
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

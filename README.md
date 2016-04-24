# Proper host name lookup by IP on Windows

This package provides `LookupAddr` function based on [GetNameInfoW](https://msdn.microsoft.com/en-us/library/windows/desktop/ms738531(v=vs.85).aspx) instead of [DnsQuery](https://msdn.microsoft.com/en-us/library/windows/desktop/ms682016(v=vs.85).aspx) used in Go standard library.
Use this one if `net.LookupAddr` cannot resolve your IP.

Windows-only

Usage:
```Go
host, err := lookup.LookupAddr("8.8.8.8")
```

## Installing

    go get github.com/postromantic/lookup


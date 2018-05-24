# ipinfo
Simple program to get some information regarding your ip address. \
ipinfo calls an available [IP-API]("http://ip-api.com/json/") to get its information.

## Dependencies
To use ipinfo you will need a working go installation.

## Build
```bash
cd ipinfo
go install .
ipinfo [API-field]
```

## Available commands
```
ipinfo [ip | city | country | countryCode | isp]
```
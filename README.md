# libre-computer-ping

### Build
```
go build -o ping main.go
```

### Usage
```
./ping 5s 1.1.1.1:443 1.0.0.1:443 8.8.8.8:443 8.8.4.4:443 208.67.222.222:443 208.67.220.220:443
```

Access Prometheus metrics at `127.0.0.1:2112/metrics`

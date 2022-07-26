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

### Running

`sudo vim /etc/cron.d/startup`:

```
@reboot root /home/libre/libre-computer-ping/ping 5s 1.1.1.1:443 1.0.0.1:443 8.8.8.8:443 8.8.4.4:443 208.67.222.222:443 208.67.220.220:443
@reboot root /home/libre/prometheus-2.37.0.linux-arm64/prometheus --config.file=/home/libre/prometheus-2.37.0.linux-arm64/prometheus.yml --web.listen-address="0.0.0.0:9090"
```

### prometheus.yml
```yml
global:
  scrape_interval: 5s
scrape_configs:
  - job_name: "ping"
    static_configs:
      - targets: ["localhost:2112"]
```

# Lightweight Monitoring Agent

## To Build
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o giysmon-agent
```

## To Run

Put giysmon-agent under your bin folder that is specified in PATH.
```bash
mv giysmon-agent /usr/local/bin
```

Edit configuration file giysmon-agent.yaml as following:
```bash
rabbit_ip: "192.168.1.100" # IP Address of your RabbitMQ Server
rabbit_port: "5672" # Port of your RabbitMQ Server
rabbit_user: "giys" # Username for your RabbitMQ Server
rabbit_pass: "giys" # Password for your RabbitMQ Server
interval: 30 # Interval for sending info to server (in Seconds)
hostname: "pardus23-deneme" # Name of the device, does not have to be the hostname, can be an alias
```

Put giysmon-agent.yaml under /etc folder.
```bash
mv giysmon-agent.yaml /etc
```



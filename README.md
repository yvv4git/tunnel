# Tunnel


## Netcat - messaging
On server:
```
nc -v -l -p 12346
nc -v -s 10.0.0.1 -l -p 12346
```

On client:
```
nc -v 10.0.0.1 12346
```


## Netcat - send file
On server:
```
nc -l -p 12346 > file.txt
```

On client:
```
nc 10.0.0.1 12346 < file.txt
```


## Netcat - send file v2
On server:
````
nc -v -s 10.0.0.1 -l -p 12346 | pv > testfile
````
I got speed 1000MiB 0:01:05 [15.3MiB/s] directly from client container to server container in docker network.

On client:
````
nc 10.0.0.1 12346 < testfile
````

## Netcat - send file v3
On server:
````
nc -v -s 192.168.2.2 -l -p 12346 | pv > testfile
````
I got speed 1000MiB 0:01:41 [9.88MiB/s] via tun & tcp+tls.

On client:
````
ip route add 192.168.2.2 dev tun0
nc 192.168.2.2 12346 < testfile
````

## TCP-dump check
Port 1234 is the port between the server and the client, i.e. the tunnel.
```
tcpdump -i eth0 'tcp and port 1234 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)' -X -vv
```

Look on the traffic:
````
tcpdump -i wlp2s0 -X -vvv
tcpdump -i wlp2s0 -X -vvv 'src host 192.168.0.13 and dst host 34.117.59.81'
tcpdump -i eth0@if78 -X -vvv 'src host 10.0.0.1 and dst host 10.0.0.2'
tcpdump -i tun0 -X -vvv 'src host 10.0.0.1 and dst host 10.0.0.2'
tcpdump -i tun0 -X -vvv 'host 10.0.0.1 or host 10.0.0.2'
tcpdump -i tun0 -X -vvv 'host 10.0.0.1 or host 10.0.0.2 or 34.160.111.145'
tcpdump -i eth0 -X -vvv 'host 10.0.0.1 or host 10.0.0.2 or 34.160.111.145'
tcpdump -i eth0 -X -vvv 'host 10.0.0.1 or host 10.0.0.2 or 142.251.31.198'
````


## Static routing on macos
Add host to routing table via tunnel interface:
```
route -n add -host 34.160.111.145 -interface utun4
route -n get 34.160.111.145
route -n delete -host 34.160.111.145
```

Check routing to host:
````
traceroute 34.160.111.145
curl ifconfig.me
````


## NAT Forwarding
```
sysctl net.ipv4.ip_forward
sysctl -w net.ipv4.ip_forward=1

iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
iptables -A FORWARD -i tun0 -o eth0 -j ACCEPT

iptables -t nat -L -v -n
iptables -L -v -n --line-numbers
iptables -t nat -L -v -n --line-numbers
```


## Run app
```
./tunnel_linux server --config ./configs/config.server.toml
sudo go run main.go client --config ./configs/config.server.toml
```

## Schema
![Schema](public/tunnel_schema.png)
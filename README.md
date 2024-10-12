# Tunnel


## Netcat - messaging
On server:
```
nc -v -l -p 12346
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


## TCPDump
Port 1234 is the port between the server and the client, i.e. the tunnel.
```
tcpdump -i eth0 'tcp and port 1234 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)' -X -vv
```


## Schema
![Schema](public/tunnel_schema.png)
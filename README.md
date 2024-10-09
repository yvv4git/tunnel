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
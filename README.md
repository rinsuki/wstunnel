# wstunnel

TCP over WebSocket

```
[something tcp server]
 |
 |  <= TCP
 |
[wstunnel server]
 ||
 || <= WebSocket
 ||
[you can include some reverse-proxy or other]
 ||
 || <= WebSocket
 ||
[wstunnel client]
 |
 | <= TCP
 |
[something tcp client]
```

## How to Use

1. Launch `wstunnel server`. e.g. `wstunnel server 0.0.0.0:8888 127.0.0.1:25565`
1. Launch `wstunnel client`. e.g. `wstunnel client 127.0.0.1:25565 ws://server:8888/ws`
1. Connect to your local port.
1. :tada:
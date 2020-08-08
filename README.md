## Fastcli

fastcli is a command line interface (CLI) for doing network testing and speedtesting using Go coroutines and http://fast.com/. It is memory efficient, performant, and descriptive in the information that it provides, including test servers/locations used, Round Trip Time (RTT), Client IP Address & Location, [fast.com](http://fast.com/) endpoint used, test duration, data downloaded, and speed in Megabits per second. 

### Installation

Assuming you already have Golang installed on your machiine, simply run
```
go get github.com/claytonblythe/fastcli
```


### Usage

```
fastcli
```


### Visual Output

```
~ $ fastcli

Connecting to fast.com...
Connecting to test servers...
Server locations:
Seattle, US, https://ipv4-c149-sea001-ix.1.oca.nflxvideo.net/speedtest?c=us&n=46562&v=5&e=1596865919&t=BoUeqjkOueI6QIoXYIE_wVhv40I, 92ms Avg RTT
Seattle, US, https://ipv4-c110-sea001-ix.1.oca.nflxvideo.net/speedtest?c=us&n=46562&v=5&e=1596865919&t=scmOHag68yXZtAYGYDOXrwnoohg, 95ms Avg RTT
San Jose, US, https://ipv4-c650-sjc002-dev-ix.1.oca.nflxvideo.net/speedtest?c=us&n=46562&v=5&e=1596865919&t=YxBcK84F3FlR4ereZHdeR87G5Ic, 116ms Avg RTT
San Jose, US, https://ipv4-c179-sjc002-ix.1.oca.nflxvideo.net/speedtest?c=us&n=46562&v=5&e=1596865919&t=iygQP9ZwD4-Q465vnOw9Ig9G8lw, 122ms Avg RTT
Los Angeles, US, https://ipv4-c012-lax009-ix.1.oca.nflxvideo.net/speedtest?c=us&n=46562&v=5&e=1596865919&t=z0n8tsHHGk-32rZkO3AtAz3zuFo, 129ms Avg RTT

Testing Download Speed...
Client: Seattle, US, 104.200.138.201
Fast.com endpoint: https://api.fast.com/netflix/speedtest/v2?https=true&token=YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm&urlCount=5

Duration: 43.58 seconds
Data downloaded: 125.00 MB
Speed: 24.06 Mbps
```

![Alt fastcli](https://github.com/claytonblythe/fastcli/raw/master/demo.png)


### Next Steps

- Add Upload Speed Support
- Add concurrent pinging of test servers
- IPv6 support
- Determine optimal # of go coroutines/concurrent urls
- Modularize code
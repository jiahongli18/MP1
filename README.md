# MP1

## Notes

* Our first initial approach was to try to establish a connection between each pair of processes as described by both Professor Tseng and Darius by implementing two for loops(one for dialing and the other for listening on requests).
* However, we came into a lot of troubles and in the end the code was not functional(after trouble shooting for many hours). The dialing portion of the code works(we including the code for our attempt in test.go)
* Thus, we decided to abandon our initial attempt and implement the other parts of the problem(concurrent TCP server, delay, messaging sending from multiple clients to one server, allowing users to choose which process to send to).
* We wanted to connect what we had in test.go with our second attempt, but there wasn't enough time in the end and we faced troubles.


## Usage

Start up the server you want to communicate with. For example, port 6001 corresponds with process 2 in our config:

```bash
cd processes
go run server.go 6001
```

Start up the rest of the processes in different terminals:
```bash
cd processes
go run client.go
```
At this point, you should get prompted with `Please enter a command:`. Please enter in the format of `send <destination> <message>`

Then the process should print:

```Sent 'hello' to process 2 , system time is Sep 14 20:48:53.521```

The server should print something like:

```Received "hello" from process 2, system time is Sep 14 20:48:59.734```

## Structure and Design


## Resources
* Darius Russell Kish
* [TCP Server](https://opensource.com/article/18/5/building-concurrent-tcp-server-go)
## Authors
* Jiahong Li
* Zheng Zhou

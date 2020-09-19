# MP1

## Usage

Configure the config.txt to your desire and then start the all the processes by running:

```bash
go run *.go [process_num](Ex. 1)
```

Until all the processes are started up, each process should print out: 

`Awaiting connections...Retrying in 2 secs.` 

At this point, you should get prompted with `Please enter a command:`. Please enter in the format of `send <destination> <message>`

Then the process should print(for example):

```Sent 'hello' to process 2 , system time is Sep 14 20:48:53.521```

And after a random delay bounded by your choices of max and min delay in the config file, the other process should print something like:

```Received "hello" from process 1, system time is Sep 14 20:48:59.734```

## Structure and Design
### File Structure and Abstraction
* Our project is broken down in the following structure:

```bash
- processes
  - server.go
  - client.go
- utils
   - utils.go
- main.go
```
We decided to abstract certain helper functions such as `FetchPorts and FetchDelay` and put them in utils because there was a lot of reused code between the files.


#### There are two layers in our design: application layer and the network layer:


### 1) Network Layer 
* Once each process is started, a goroutine is fired up immediately and runs until the program is exited. This goroutine calls a function in 
our `server.go` in order to make each process start its server and listens for requests.
* The next step in the network layer is creating the dial connections between each pair of processes in the config file. We put this process in a function called `initialize()` in main.go. The idea is to try alive ports and dial between the process itself and every other process in the config file. This connection is attempted every two seconds until a successful connection is made(we coded a delay here because we didn't want the program to spam the connections too fast).
* After this is done, the networking layer is fully complete.
* At this point the server is listening on requests and when a request is incoming, `unicast_receive` is called in order to print the time, message, and sender.

### 2) Application Layer 
* Our application code is located in `client.go` under processes.
* The process runs in a loop to continuously listen and ask the user for input as describe in #Usage at the top and then calls `unicast_send` once an input is received. 
* It calls the helper function `FetchHostPort` to convert the user's choosen destination into ip/host in order to be able to send the message using a TCP Dial. 
* However, the message is not sent instantly because we want to mimic a network delay. In order to do this, we call a goroutine called `Delay()` inside of `utils.go` that utilizes waitgroups.
* After this delay, the message is sent and in the other process, `unicast_receive` is kicked off. 

## Resources
* Darius Russell Kish
* Professor Tseng
* [TCP Server](https://opensource.com/article/18/5/building-concurrent-tcp-server-go)
* [WaitGroups](https://www.golangprograms.com/how-to-use-waitgroup-to-delay-execution-of-the-main-function-until-after-all-goroutines-are-complete.html)
## Authors
* Jiahong Li
* Zheng Zhou

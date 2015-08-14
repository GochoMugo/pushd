
# pusher

> A [pushbullet](https://pushbullet.com/) daemon; for pushing notifications from your system

> This is still a **Proof of Concept**. It is **not** production ready.


## use cases:

With the daemon running, you can push notifications when:

* your battery gets low
* system shutdowns
* anything, worth noting, happens

*(you do this stuff on your own!)*


## installation:

Using Go tools:

```bash
$ go get github.com/GochoMugo/pusher
```


## usage:


### configuring the daemon:

The daemon requires an access token set as the environment variable `${pusher_TOKEN}`. For example,

```bash
# set environment variable
export PUSHER_TOKEN='AbCdEfgHijKLmNopQrstUvWXyZ'
```


### starting the daemon:

```bash
$ pusher start
```

This starts the daemon **but** remains attached to the terminal. **This is by design**. You need to use external programs such as [forever](https://github.com/foreverjs/forever), [upstart](http://upstart.ubuntu.com/), [supervisord](http://supervisord.org/index.html), etc to keep the daemon running forever and in background.

The daemon listens on port `9701` by default. However, you can specify a port using the flag `--port=<port-number>` e.g. `--port=8080`.


### pushing a notification:

```bash
$ pusher notify --message="<message-goes-here>"
# for example
$ pusher notify --message="battery low"
```

This will broadcast the message(note/notification) to all of your devices using `username + @ + hostname` as the header.


### checking status of the daemon:

```bash
$ pusher status
```

This pings the daemon.


### stopping the daemon:

```bash
$ pusher stop
```

This sends a message to the daemon asking it to stop.


## license:

__The MIT License (MIT)__

Copyright (c) 2015 GochoMugo <mugo@forfuture.co.ke>

Examine whether half-open connections are left hanging around even if the server periodically tries to write to the connection.

To reproduce:

* Run this server
* Connect from a client
* Unplug the networking from the client (or if on a mobile, turn off wifi)


Expected:

* As soon as the server tries to write the next heartbeat, it should close the connection (and the entry in lsof should disappear)

Actual:

* The connection hangs around for 90 mins


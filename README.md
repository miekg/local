# local

## Name

*local* - respond to local names.

## Description

*local* will respond with a basic reply to a "local request". Local request are defined to be
names in the following zones: localhost, 0.in-addr.arpa, 127.in-addr.arpa and 255.in-addr.arpa.

With *local* enabled any query falling under these zones will get a reply. The prevents the query
from "escaping" to the internet and putting strain on external infrastructure.

Most zones are empty, only `localhost.` address (A and AAAA) records a defined and a
`1.0.0.127.in-addr.arpa.` reverse record.

## Syntax

~~~ txt
local
~~~

## Examples

~~~ corefile
. {
    local
}
~~~

## Also See

BIND9's configuration in Debian comes with these zones preconfigured.

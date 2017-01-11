# URL Shortener in Golang
## How to use
Requirements:

* Golang compatible with v. 1.6
* make
* bash
```
git clone https://github.com/maxbeutel/urlshortener.git
make
make demo
```
The server will bind to localhost, at port 8080.
### Considerations
* The URLs are only stored _in memory_ (in a Hashmap), for the sake of demoing this functionality.
* There are some obvious _security holes_, e. g. as the hashing algorithm is base 62, it's trivial to enumerate all existing URLs.
* The input validation _doesn't protect against DoS attacks_, e. g. by supplying arbitrary long hashes as input
* It would probably make sense to not expose the Go webserver directly to the internet, but _use nginx as a reverse proxy_ in front of it.
* It _lacks tests_ for the handlers, which I elided but was the source of some bugs.
* There should be some _integration tests_, which test how the server behaves under load.

### Design principle
The idea was was to implement a minimal URL shortener service in Golang.

At first I decided between file-storage or in-memory storage for the URLs, in order to avoid a heavyweight external database as storage. I decided for in-memory because I wanted to see how this works in conjunction with Go's goroutines. I found that Go's map is not threadsafe by default but needs locking.

The general implementation is straightforward, two http handlers are registered. The shorten handler creates a new surrogate key, by atomically incrementing an integer. This integer is converted from base 10 to base 62 in order to have a usable hash for the URL. The original URL is stored using the id as the key in the map. The shorturl is returned then to the client in the response body.

The expand handler converts the base 62 encoded hash back to base 10 and does a lookup in the map if the decoded id exists. If this is the case, a redirect response is returned to the client pointing to the original URL.

Both handlers implement some basic validation to protect against invalid input.

### Things I learned
As I never used Golang before, this was an interesting learning experience.

- Concurrency works quite well in Go and is made explicit via goroutines/channels.
- Go comes with a rich standard library, it wasn't needed to include any external libraries.
- Working with the $GOPATH is a bit weird, it seems it's not needed anymore in v. 1.8, but I found it quite cumbersome and therefore wrote a Makefile to abstract away the build process/setting of the env variable.
- The language-level support for unit tests is very nice.
- Writing Makefiles is still not so funny and will maybe never be.

## Implemented methods

### Create new shorturl
```
curl -s -v -sX POST 'localhost:8080/shorten' -d http://www.google.com
...
< HTTP/1.1 200 OK
< Date: Wed, 11 Jan 2017 06:49:06 GMT
< Content-Length: 23
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
http://localhost:8080/<hash>
```
The response body will contain the shorturl as string.

### Request shorturl

```
curl -s -v <the short url>
```
Will contain a 302 redirect to the saved shortened URL.

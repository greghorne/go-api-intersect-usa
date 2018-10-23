# go-api-intersects-usa

_**First try at GOLANG.**_

Write and deploy a simple API where given an x,y coordinate the API will determine if the coordinates intersect U.S. boundaries.  Return value is 1 (true) or 0 (false).

I use this functionality in a few of my other personal projects so  I may point them to this API once deployed.

_**API:**_ https://zotac1.ddns.net/v1/getinfo/{lng}/{lat}

    - ex. http://zotac1.ddns.net/v1/intersects-usa/-95.992775/36.153980 ==> returns 1 (Oklahoma)
    - ex. http://zotac1.ddns.net/v1/intersects-usa/48.516388/15.552727  ==> returns 0 (Yemen)

_**Depolyment:**_ 

	- Docker Container on Ubuntu Server 18.04 - Phenom II x6 1045T - 16GB RAM - 160GB SSD

_**Tech Stack:**_

	- go1.10.1 linux/386
	- Docker version 18.06.1-ce, build e68fc7a
	- Ubuntu Server 18.04
	- Postman (testing)
# go-api-intersect-usa

_**First try at GOLANG.**_

Write and deploy a simple API where given an x,y coordinate the API will determine if the coordinates intersect U.S. State boundaries.  Returns true or false.

The API interacts with a Postgres/PostGIS server loaded with U.S. boundaries.  The DB resides on a Raspberry Pi II.

I use this functionality in a few of my other personal projects so  I may point them to this API once deployed.

_**API:**_ 

    - http://zotac1.ddns.net:8000/v1/intersects-usa/{lng}/{lat}
    - 'lng' and 'lat' are decimal degrees
    - returns ==> {"intersects":true or false}

    - ex. http://zotac1.ddns.net:8000/v1/intersects-usa/-95.992775/36.153980 ==> {"intersects":true} (Oklahoma)
    - ex. http://zotac1.ddns.net:8000/v1/intersects-usa/48.516388/15.552727  ==> {"intersects":false} (Yemen)

_**Depolyment:**_ 

	- Docker Container on Ubuntu Server 18.04 - Phenom II x6 1045T - 16GB RAM - 160GB SSD

_**Tech Stack:**_

    - Vagrant ubuntu/xenial32
	- go1.10.1 linux/386
    - Postman (testing)
	- Docker version 18.06.1-ce, build e68fc7a
	- Ubuntu Server 18.04

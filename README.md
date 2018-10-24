# go-api-intersect-usa

_**First try at GOLANG.**_

Write and deploy a basic API where given an x,y coordinate the API will determine if the coordinates intersect U.S. State boundaries.  API returns true or false.

The API interacts with a Postgres/PostGIS server loaded with U.S. boundaries.  The DB resides on a Raspberry Pi II.

_**API:**_ 

    - http://zotac1.ddns.net:8000/v1/intersects-usa/{lng}/{lat}
    - 'lng' and 'lat' are decimal degrees
    - returns ==> {"intersects":true or false}

    - ex. http://zotac1.ddns.net:8000/v1/intersects-usa/-95.992775/36.153980 ==> {"intersects":true} (Oklahoma)
    - ex. http://zotac1.ddns.net:8000/v1/intersects-usa/48.516388/15.552727  ==> {"intersects":false} (Yemen)

_**Depolyment:**_ 

	- Docker Container on Raspberry Pi 3 Model B running HypriotOS (Debian v9.4)

_**Tech Stack:**_

	Developement
    - Vagrant ubuntu/xenial32
	- go1.10.1 linux/386
    - Postman (testing)
    
    Depolyment
	- Docker version 18.04.0-ce, build 3d479c0
	- go1.11.1 linux/arm64


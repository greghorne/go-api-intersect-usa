# go-api-intersects-usa

_**First try at GOLANG.**_

Write and deploy a simple API where given an x,y coordinate the API will determine if the coordinates intersect U.S. boundaries.  Return value is 1 (true) or 0 (false).

I use this functionality in a few of my other personal projects so  I may point them to this API once deployed.

_**API:**_ https://someserver.com/v1/getinfo/{lng}/{lat}

    - ex. https://someserver.com/v1/getinfo/-95.992775/36.153980 ==> returns 1 (Oklahoma)
    - ex. https://someserver.com/v1/getinfo/48.516388/15.552727  ==> returns 0 (Yemen)

_**Depolyment:**_ ?

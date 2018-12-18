### Weather GO

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-89%25-brightgreen.svg?longCache=true&style=flat)</a>
[![Build Status](https://travis-ci.org/flaviojmendes/weathergo.svg?branch=master)](https://travis-ci.org/flaviojmendes/weathergo)

Simple App that returns the __Weather__ based on the _Latitude_ and _Longitude_.

##### Before you start:

It is required to have an _ENV Variable_ named **CONFIG_PATH** 
 pointing to a YAML file with the following scructure.

```
Port: :8000       # the port you want
CacheExp: 5       # time in minutes to the cache expiration
CachePurge: 10    # time in minutes to the purge of expired cached weathers
AuthKey: secretKey
AuthSecret: secretSecret
OpenWeatherKeys:  # any number of keys you want
  - key1
  - key2
  - ...
``` 
_You can find an example of this file [here](config_sample.yml)._



**And now?**
 
To Get the weather you should do a GET Request to:

```
HTTP GET http://localhost:8000/weather/{latitude}/{longitude}/{provider}

HEADERS:
X-Auth-Key: {your_key}
X-Auth-Secret: {your_key}
```


#### Example

```
HTTP GET http://localhost:8000/weather/52.1044634/-9.7957984/OPENWEATHER

HEADERS:
X-Auth-Key: k3y1
X-Auth-Secret: s3cr3t

```

Will result in:

```
{  
   "lat":"52.1044634",
   "lon":"-9.7957984",
   "temp":9.78,
   "location":"Killorglin",
   "humidity":94,
   "created_at":"2018-12-05T21:39:04.769538Z"
}
```

.
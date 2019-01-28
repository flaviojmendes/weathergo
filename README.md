### Weather GO
![Logo](logo.png)

[![Build Status](https://travis-ci.org/flaviojmendes/weathergo.svg?branch=master)](https://travis-ci.org/flaviojmendes/weathergo)
[![codecov](https://codecov.io/gh/flaviojmendes/weathergo/branch/master/graph/badge.svg)](https://codecov.io/gh/flaviojmendes/weathergo)

Simple App that returns the __Weather__ based on the _Latitude_ and _Longitude_.

##### Before you start:

It is required to have an _ENV Variable_ named **CONFIG_PATH** 
 pointing to a YAML file with the following scructure.

```
Port: :8000       # the port you want
CacheExp: 5       # time in minutes to the cache expiration
CachePurge: 10    # time in minutes to the purge of expired cached weathers
DbFile: path/to/your/weather.db
WhiteListHosts:
  - host1
  - host2
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

```


#### Example

```
HTTP GET http://localhost:8000/weather/52.1044634/-9.7957984/OPENWEATHER

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

### Getting from DockerHub

```
docker run -i -p 443:443 -e CONFIG_FILE='/usr/local/config_sample.yml' -v {PATH_TO_YOUR_LOCAL_YML_CONFIG}:/usr/local/config_sample.yml flaviojmendes/weather
```
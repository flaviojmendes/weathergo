### Weather GO

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-72%25-brightgreen.svg?longCache=true&style=flat)</a>

Simple App that returns the __Weather__ based on the _Latitude_ and _Longitude_.

##### Before you start:

It is required to have an _ENV Variable_ named **CONFIG_PATH** 
 pointing to a YAML file with the following scructure.

```
Port: :8000       # the port you want
CacheExp: 5       # time in minutes to the cache expiration
CachePurge: 10    # time in minutes to the purge of expired cached weathers
OpenWeatherKeys:  # any number of keys you want
  - key1
  - key2
  - ...
``` 
_You can find an example of this file [here](config_sample.yml)._



**And now?**
 
To Get the weather you should do a GET Request to:

`HTTP GET http://localhost:8000/weather/{latitude}/{longitude}/{provider}`

#### Example

`HTTP GET http://localhost:8000/weather/52.1044634/-9.7957984/OPENWEATHER`

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
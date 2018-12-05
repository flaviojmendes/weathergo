### Weather GO

Simple App that returns the __Weather__ based on the _Latitude_ and _Longitude_.

Example:

`HTTP GET http://localhost:8000/weather/52.1044634/-9.7957984`

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
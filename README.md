# JSON-to-map [![Build Status](https://travis-ci.org/arthurkiller/JSON-to-map.svg?branch=master)](https://travis-ci.org/arthurkiller/JSON-to-map)[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkiller/JSON-to-map)](https://goreportcard.com/report/github.com/arthurkiller/JSON-to-map)
Convert the json into a string <---> string map

the key can be set into multi-prefix style to make sure if there get some keys in the same.
```
    "Subtree":{"Redis":{"Masters":
```

will be transfer into this key:
```
    Subtree-Redis-Master  <----->  [value]...
```

## TODO 
add more type support just like the struct & file ...

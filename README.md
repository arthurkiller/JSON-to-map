# JSON-to-map
Convert the json into a string <---> string map

the key can be set into multi-prefix style to make sure if there get some keys in the same.
```
    "Subtree":{"Redis":{"Masters":
```

will be transfer into this key:
```
    Subtree-Redis-Master  <----->  [value]...
```

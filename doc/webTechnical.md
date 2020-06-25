## Web technical

### Get information in request
#### context: func(res http.ResponseWriter, req *http.Request)
```text
// Get query .../locations?search=vietnam
query := req.URL.Query()
search := query.Get("search")
```
```text
// Get path .../restaurant/noodle
path := req.URL.Path
parts := strings.Split(path,"/") // [restaurant, noodle]
```


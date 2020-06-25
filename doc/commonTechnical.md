## Common technical

### Read os.Stdin with channel
```text
func readStdin(out chan<- []byte) {
    for {
        data := make([]byte, 1024)
        l, _ := os.Stdin.Read(data)
        if l > 0 {
            out <- data
        }
    }
}
```


### Mapping file json to object
##### Library "encoding/json"
##### configuration.json
```json
{
"enabled": true,
"path": "/usr/local"
}
```
##### coding
```text
type configuration struct {
    Enabled bool    `json:"enabled"`
    Path string     `json:"path"`
}

func convert() {
    file, _ := os.Open("configuration.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    conf := configuration{}
    err := decoder.Decode(&conf)
}
```

## Mapping file yaml to object
#### library"github.com/kylelemons/go-gypsy/yaml"

#### configuration.yaml
```yaml
enabled: true
path: /usr/local
```
#### code
```text
func convert() {
    config, err := yaml.ReadFile("configuration.yaml")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println(config.Get("path"))
    fmt.Println(config.GetBool("enabled"))
}
```
```text
type configuration struct {
    Enabled bool    `yaml:"enabled"`
    Path string     `yaml:"path"`
}
func convert() {
        data, err := ioutil.ReadFile("configuration.yaml")
    	if err != nil {
    		panic(err)
    	}
    
    	errUnmarshal := yaml.Unmarshal([]byte(data), &configuration)
    
    	if errUnmarshal != nil {
    		panic(err)
    	}
}
```

### Get environment variable
```text
os.Getenv("PORT")
```


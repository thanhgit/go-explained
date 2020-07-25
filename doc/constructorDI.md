#### Constructor DI
- require dependencies as parameters in constructor method
```text
type Address struct {
    street      string
    ward        string
    district    string
    province    string
}
type Peron struct {
    name    String
    address Address
}
func NewPerson(name string, address Address) *Person {}
```
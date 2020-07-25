## Common coding 

#### Start with simple
#### Apply just enough abstraction 
#### Follow industry, team and language conventions
#### Export only what you must
```text
NewPet("Meo", true) // true -> difficult to understand
---
NewPet("Meo")
```
#### Unit test for code, they make sure you have no bugs 
* ##### recommend testing from high-level or black-box level
Type of testing | Description |
--- | --- |
Happy path | compare actual result and expected result. Tend to document how to use the code 
Input errors | make sure that our code handles in a predictable way 
Dependencies issues | when dependencies fail 
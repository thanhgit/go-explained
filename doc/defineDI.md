#### Common DI
DI is coding such as functions, struct, ... depend on the abstractions
Because change to them, but not change its code => decoupling

- ###### DI reduce knowledge required by expressing dependencies in a abstraction 
- ###### DI enable test in isolation
- ###### DI enable us to quickly and reliable test situations
- ###### DI reduces impact of extensions or changes

### Smell code fall into 4 categories:
* #### Code bloat
```text
- code is added to functions or struct 
+ Long methods
+ Long struct
+ Long parameters
+ Long conditional blocks: such as switch case
```
* #### Resistance to change
```text
- difficult/slow to add new features
+ harder to write test
+ Site-effect 
```
* #### Wasted effort
```text
- cost to maintain the code is higher it needs to be
+ duplicated code
+ excessive comment
+ complicated code to understand
+ format and convention code
```
* #### Tight coupling
```text
- add complexity and maintenance cost
+ dependence on God object 
+ circular dependencies: (A->B and B->A)
+ object orgy: page.httpClient.Get(url) (inf)
+ yoyo-problem: inheritance is so long and complicated 
+ feature envy: when a function is responsible for other object functionality
```
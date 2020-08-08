# Aggregate pattern

## Difference between composition and aggregation
- Think of aggregation as referring to the "has-a" relationship
- Think of composition as referring to the "has-to-has-a"

- Aggreation in the UML class diagram is an open diamond ("use" relationship)
- Composition in the URL class diagram is a closed diamond ("depends on" relationship) and it's a stronger relationship where composed object must be a part of composer

## Usecase: using aggregate design pattern to trasmission data to the application layer 
- API composition patterns resolve:
    - Diff clients may have diff access privileges for the same data
    - Diff clients may need diff data formats of the same data
    - Is implemented as gateway API 
- Event sourcing patterns resolve:
    - atomically updating the database and also publishing an event
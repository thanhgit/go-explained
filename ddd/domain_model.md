# Domain model in detail
- The intent of the modeling process is to capture a set of <b>high level and low level DDD artifacts</b>
    - high level artifacts: low degree of implementation, more design concepts with minimal physical artifacts required 
    - low level artifacts: have a high degree of implementation
- -> domain modeling processing identify whether monolithic or microservices architecture

## Business language
- Business entities
- Business rules
- Business flows 
- Business operations 
- Business events 

## Technical in DDD world
- Aggregates - Entities - Value objects 
- Domain rules
- Sagas
- Commands / Queries 
- Events

## Aggregates
- Aggregate = principal identifier of your bounded context 
- Entity object = secondary identifiers of your bounded context are responsible for capturing all state 
- Business rules associated with the bounded context 

## Aggregate identifiers
- each aggregate needs to be uniquely identified using a Aggregate Identifier, that is implemented using a business key

## Domain rules 
- are pure business rule definitions 
- assist the aggregate for any kind of business logic execution within a bounded context 
- can use to validate aggregate 

## Commands / Queries 
- represent any kind of operations within the bounded context 
    - Command : affect the state of the aggregate/entity
    - Query : query the state of the aggregate/entity 

## Events 
- captures any kind of state change of aggregate or entity within the bounded context 

## Sagas
- flush out any kind of business processes / work-flows within your business domain
- can span across multiple bounded context, react to multiple business events across bounded context and "orchestrate the business process" by coordinating interactions
- maintain data consistency for use case, that may span across multiple microservices, with 2 ways:
    - Event Choreography: in a particular saga will raise and subscribe to events directly
    - Event Orchestration: the lifecycle coordinate happens through a central component, that is responsible for saga creation, coordination of the flow across bounded context 


## Domain model 
- Core domain model: aggregate, aggregate identifiers, entities and value objects
- Domain model operations: commands, queries and events 

## Domain model services:
- is used by 2 reasons;
    - to enable domain model available to external parties through well-defined interface
    - interacting with external parties to persist bounded context's state to datastore, pushlish bounded context's state change events to external message broker or communicate with other bounded context 
- 3 type of domain model services:
    - inbound services: implement well-defined interface, which enable external parites to interact with the domain model
    - outbound services: implement all interactions with external responsitoryes or other boundedd context  
    - application services: act as a facade layer between domain model and (inbound and outbound services)
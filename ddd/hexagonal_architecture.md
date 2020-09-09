# Hexagonal architecture in detail
- perfect fit to help us model/ design and implement the domain model supporting services 
- uses the concepts of ports and adapters to implement the domain model services 
    - inbound port: provide a interface to the business operations of our domain model. This is implemented via the application services 
    - outbound port: provide a interface to the technical operations required by our domain model. Domain model use to store or publish state from the sub-domain 
    - inbound adapter: use the inbound port to provide the capabilities for external clients to consume the domain model 
    - outbound adapter: is an implementation of the outbound port for the specific repository 
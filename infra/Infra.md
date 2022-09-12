## Device Repository Service Infrastructure

The Device Repository Service is hosted in a Kubernetes cluster.  
The service depends on:

1. A discovery service - for now the service subscribes to a NATS service to subscribe to discovery events.

The Repository Service uses

1. Neo4j
2. Istio
3. Redis (undecided - instead using a service like memcached) - (justification)[https://aws.amazon.com/elasticache/redis-vs-memcached/]

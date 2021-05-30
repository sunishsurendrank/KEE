# KEE (Kube Event Expoter)

## Requirement

Design and implement an application that reads all the events in kubernetes and stores them in backend.  **-Considered in Design**

The data should be viewable on a timeseries view using tools like splunk or wavefront etc. **-Need to check**

The system should store event data for 5 years. **-Considered in Design**

The data stored should be highly available and consistent. **-Considered in Design**

While fetching the data if there are any failure events identified those events should trigger an alert to the users such that they can be handled. **-Considered in Design**

The system should have an application that has the following end points that can handle ~1000tps and latency of these api's should be minimal (~200ms)**-Need to check in Ingress**

/events - should display all the events in the kubernetes cluster /events/namespace - should display events in a namespace /events/namespace/resource - should display events of a specific resource in namespace. resources could be ingress, deployment, pod etc. **-Considered in Design**

Have a ratelimit on the endpoints such that the same user agent can only trigger an api call/sec. **-Need to check in Ingress**

The application should be dockerized, should run in a kubernetes cluster. **-Considered in Design**

The application should expose metrics to a monitoring tool and any failure on the application should also trigger an alert to the users. **-Considered in Design**

### Extended:

 Extend the above application to handle the events across multiple clusters to store event data. **-Considered in Design**

 The endpoints in this case would be:/events - 5 events from all clusters. /events/clustername - should display all the events in the kubernetes cluster /events/clustername/namespace - should display events in a namespace of a kubernetes cluster /events/clustername/namespace/resource - should display events of a specific resource in namespace. resources could be ingress, deployment, pod etc., in a kubernetes cluster. **-Considered in Design**


## Architecture

![image](/docs/pictures/architecture.png)

## Deployment Architecture

![image](/docs/pictures/deployment__architecture.png)

## Code Folder Structure
 
 Polyglot Microservice

 Event Manager Microservice (Golang) : /event_manager/

 Event Query Microservice(Python) : /event_query/


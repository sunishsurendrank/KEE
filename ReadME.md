# KEE (Kube Event Expoter)

## Requirement

Design and implement an application that reads all the events in kubernetes and stores them in backend. The data should be viewable on a timeseries view using tools like splunk or wavefront etc.,. The system should store event data for 5 years. The data stored should be highly available and consistent. While fetching the data if there are any failure events identified those events should trigger an alert to the users such that they can be handled.

The system should have an application that has the following end points that can handle ~1000tps and latency of these api's should be minimal (~200ms):

/events - should display all the events in the kubernetes cluster
/events/namespace - should display events in a namespace
/events/namespace/resource - should display events of a specific resource in namespace. resources could be ingress, deployment, pod etc.,
Have a ratelimit on the endpoints such that the same useragent can only trigger an api call/sec.
The application should be dockerized, should run in a kubernetes cluster. The application should expose metrics to a monitoring tool and any failure on the application should also trigger an alert to the users.

 
Extended: Extend the above application to handle the events across multiple clusters to store event data. The endpoints in this case would be:/events - 5 events from all clusters.
/events/clustername - should display all the events in the kubernetes cluster
/events/clustername/namespace - should display events in a namespace of a kubernetes cluster
/events/clustername/namespace/resource - should display events of a specific resource in namespace. resources could be ingress, deployment, pod etc., in a kubernetes cluster.

## Architecture

![image](/docs/pictures/architecture.png)

## Deployment Architecture

![image](/docs/pictures/deployment_architecture.png)

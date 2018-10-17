## Cloud-native example application

### Run Consul

`make start-consul`

### Run the application (with Prometheus)

`make start`
  
In the browser open the url http://localhost or check the load balancer stats http://localhost/stats. 

#### Scale the application

`docker-compose scale app=2`

### Build the application

`make build`

### Destroy all

`make stop`
## Cloud-native example application

### Run Consul

`make start-consul`

### Run the application (with Prometheus)

`make start`

#### Get the exposed port

```
docker-compose ps app
         Name            Command   State            Ports         
------------------------------------------------------------------
cloud-native-app_app_1   ./app     Up      0.0.0.0:32797->8080/tcp
```   
So, the url to open in the browser is http://localhost:32797

#### Scale the application

`docker-compose scale app=2`

### Build the application

`make build`

### Destroy all

`make stop`
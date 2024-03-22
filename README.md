# Employee Registration API

## Build and Run

### Docker

Network
```bash
docker network create employee-registration-network
```

MySQL Container
```bash
docker run \
  --name employee-registration-mysql \
  --network employee-registration-network \
  -p 3306:3306 \
  -d \
  -v $(pwd)/migrations:/docker-entrypoint-initdb.d/ \
  -e MYSQL_ROOT_PASSWORD=123 \
  mysql:8.3.0
```

RabbitMQ Container
```bash
docker run \
  --name employee-registration-rabbitmq \
  --network employee-registration-network \
  -p 5672:5672 \
  -p 15672:15672 \
  -d \
  rabbitmq:3.12-management
```

API Image
```bash
docker build -t employee-registration-api:latest .
```

API Container
```bash
docker run \
  --name=employee-registration-api \
  --network=employee-registration-network \
  -p 8080:8080 \
  -d \
  -e MYSQL_HOST=employee-registration-mysql \
  -e MYSQL_PORT=3306 \
  -e MYSQL_USERNAME=root \
  -e MYSQL_PASSWORD=123 \
  -e MYSQL_SCHEMA=employee_registration \
  -e RABBITMQ_HOST=employee-registration-rabbitmq \
  -e RABBITMQ_PORT=5672 \
  -e RABBITMQ_USERNAME=guest \
  -e RABBITMQ_PASSWORD=guest \
  employee-registration-api
```

### Docker Compose
```bash
docker-compose up -d
```

### Kubernetes

#### Secrets DB
```bash
kubectl create secret generic production-mongo \
  --from-literal=username=CHANGE_HERE \
  --from-literal=password=CHANGE_HERE
```

#### Secrets RabbitMQ
```bash
kubectl create secret generic production-rabbitmq \
  --from-literal=username=CHANGE_HERE \
  --from-literal=password=CHANGE_HERE
```

#### Mongo Pods and Services
```bash
kubectl apply -f kubernetes/db/pod.yaml
kubectl apply -f kubernetes/db/service.yaml
kubectl apply -f kubernetes/db/nodeport.yaml # Optional to local access
```

#### RabbitMQ Pods and Services
```bash
kubectl apply -f kubernetes/rabbitmq/pod.yaml
kubectl apply -f kubernetes/rabbitmq/service.yaml
kubectl apply -f kubernetes/rabbitmq/nodeport.yaml # Optional to local access
kubectl apply -f kubernetes/rabbitmq/load-balancer-service.yaml # Optional to local access
kubectl apply -f kubernetes/rabbitmq/load-balancer-service-2.yaml # Optional to local access
```

#### API Pods and Services
```bash
kubectl apply -f kubernetes/api/deployment.yaml
kubectl apply -f kubernetes/api/load-balancer-service.yaml
kubectl apply -f kubernetes/api/service.yaml
```


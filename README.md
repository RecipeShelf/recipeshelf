# recipeshelf - A recipe bookmark service

## Kubernetes cluster
Note: This is to get everything running in minikube locally



### Cockroach Database
- Install [CochroackDB](https://www.cockroachlabs.com/docs/stable/orchestrate-a-local-cluster-with-kubernetes.html) in Kubernetes
- Install the custom resource definition used by the cockroach operator  
  ```kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.11.0/install/crds.yaml```
- Install the operator  
  `kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.11.0/install/operator.yaml`
- Change your context to the cockroach operator system  
  `kubectl config set-context --current --namespace=cockroach-operator-system`
- Initialize the cluster    
  `kubectl apply -f crdbcluster.yaml`
- Setup a pod running the cockroach client  
  `kubectl create -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.11.0/examples/client-secure-operator.yaml`
- Connect to the cockroach client  
  `kubectl exec -it cockroachdb-client-secure -- ./cockroach sql --certs-dir=/cockroach/cockroach-certs --host=cockroachdb-public`  
  Run the following commands in the console to setup a user
  - `CREATE USER [username] WITH PASSWORD '[password]'`  
  - `GRANT admin TO [username]`
  - `\q`  
- [DB Console](https://localhost:8080) port forwarding  
  `kubectl port-forward service/cockroachdb-public 8080`
- Port forwarding for service acces from outside cluster  
  `kubectl port-forward service/cockroachdb 26257:26257`

### Redis Cache
- Create the redis namespace and set your context to it  
  - `kubectl create namespace redis`
  - `kubectl config set-context --current --namespace=redis`
- Use helm to install [redis](https://bitnami.com/stack/redis/helm)
  `helm install redis oci://registry-1.docker.io/bitnamicharts/redis`
- To get your password run  
  `kubectl get secret --namespace redis redis -o jsonpath="{.data.redis-password}" | base64 -d`
- To connect to your Redis&reg; server:
  1. Echo redis password to an environment variable  
    `export REDIS_PASSWORD=$(kubectl get secret --namespace redis redis -o jsonpath="{.data.redis-password}" | base64 -d)`
  2. Run a Redis&reg; pod that you can use as a client  
    `kubectl run --namespace redis redis-client --restart='Never'  --env  REDIS_PASSWORD=$REDIS_PASSWORD  --image docker.io/bitnami/redis:7.0.11-debian-11-r12 --command -- sleep infinity`
  3. Use the following command to attach to the pod  
    `kubectl exec --tty -i redis-client --namespace redis -- bash`
  4. Connect using the Redis&reg; CLI  
    `REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-master`  
    `REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-replicas`
- Setup port forwarding to connect to the cache from outside the cluster   
  `kubectl port-forward --namespace redis svc/redis-master 6379:6379`

### Applications
- Create namespace for recipeshelf  
  `kubectl create namespace recipeshelf`
- Replicate redis secret in recipeshelf namespace  
  `kubectl get secret --namespace redis redis -oyaml | grep -v '^\s*namespace:\s' | kubectl apply --namespace=recipeshelf -f -`

## Scraper

Scraper reads urls from the specified redis import queue and writes the scraped data to the export queue

- Provide a .env file for debugging containing the following  
  redis-host=localhost:6379
  redis-password=[RedisPassword]  
  import-queue=urls-to-scrape  
  import-queue-timeout=1s  
  export-queue=recipes-scraped  
  dead-letter-queue=urls-to-scrape-dead  
- Running it in kubernetes  
  1. Build an image locally with the name "scraper".  
     `docker build -t recipeshelf/scraper src/scraper/.`
  2. Load the image  
     `minikube image load recipeshelf/scraper`
  3. Create the cronjob  
     `kubectl --namespace recipeshelf create -f kube/scraper.yaml`
     
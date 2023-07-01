# recipeshelf - A recipe bookmark service

## Kubernetes cluster

### Cockroach Database
- Install [CochroackDB](https://www.cockroachlabs.com/docs/stable/orchestrate-a-local-cluster-with-kubernetes.html) in Kubernetes
- Install the custom resource definition used by the cockroach operator  
  `kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.10.0/install/crds.yaml`
- Install the operator  
  `kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.10.0/install/operator.yaml`
- Initialize the cluster  
  `kubectl apply -f crdbcluster.yaml`
- Connect to the DB console  
  `kubectl --namespace cockroach-operator-system exec -it cockroachdb-client-secure -- ./cockroach sql --certs-dir=/cockroach/cockroach-certs --host=cockroachdb-public`
- Setup port forwarding for access from a sql GUI Tool  
  `kubectl port-forward --namespace cockroach-operator-system service/cockroachdb 26257:26257`

### Redis Cache
- Use helm to install [redis](https://bitnami.com/stack/redis/helm)
  `helm install redis oci://registry-1.docker.io/bitnamicharts/redis`
- To get your password run  
  `export REDIS_PASSWORD=$(kubectl get secret --namespace redis redis -o jsonpath="{.data.redis-password}" | base64 -d)`
- To connect to your Redis&reg; server:
  1. Run a Redis&reg; pod that you can use as a client  
    `kubectl run --namespace redis redis-client --restart='Never'  --env  REDIS_PASSWORD=$REDIS_PASSWORD  --image docker.io/bitnami/redis:7.0.11-debian-11-r12 --command -- sleep infinity`
  2. Use the following command to attach to the pod  
    `kubectl exec --tty -i redis-client --namespace redis -- bash`
  3. Connect using the Redis&reg; CLI  
    `REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-master`  
    `REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-replicas`
- Setup port forwarding to connect to the cache from outside the cluster   
  `kubectl port-forward --namespace redis svc/redis-master 6379:6379`

## Scraper

Scraper reads urls from the specified redis import queue and writes the scraped data to the export queue

### Provide a .env file for debugging containing the dollowing  
REDIS_PASSWORD=[RedisPassword]  
IMPORT_QUEUE=urls-to-scrape  
EXPORT_QUEUE=recipes-scraped  
DEAD_LETTER_QUEUE=urls-to-scrape-dead

### Building an image
`cd src/scraper`  

Build an image with the name "scraper" and tag "1".  
`docker build -t scraper:1 .`

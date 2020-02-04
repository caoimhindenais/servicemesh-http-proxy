#!/usr/bin/env bash
eval $(minikube docker-env)
kubectl delete ns nash 2>/dev/null
mvn clean install -DskipTests -f demo/pom.xml
docker build -t springguides/demo demo
docker build -t nash/init init
docker build -t nash/proxy proxy
kubectl apply -f deployment.yaml -n nash
export INGRESS_IP=`minikube ip`
sleep 10
echo $INGRESS_IP
curl -v  $INGRESS_IP/villager -H 'Content-type:application/json' -d '{"name": "Bilbo"}' | jq
curl -v POST $INGRESS_IP/villager -H 'Content-type:application/json' -d '{"name": "Dragon"}' | jq






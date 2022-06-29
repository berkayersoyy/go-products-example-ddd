echo "Stopping"

echo "Deleting deployments"
kubectl delete deployments go-app
kubectl delete deployments go-app-dynamodb
kubectl delete deployments go-app-mysql
kubectl delete deployments go-app-redis
kubectl delete deployments go-app-grafana
kubectl delete deployments go-app-prometheus
kubectl delete deployments go-app-jaeger

echo "Deleting services"
kubectl delete svc go-app
kubectl delete svc go-app-dynamodb
kubectl delete svc go-app-mysql
kubectl delete svc go-app-redis
kubectl delete svc go-app-grafana
kubectl delete svc go-app-prometheus
kubectl delete svc go-app-jaeger

#echo "Deleting pv/pvc"
#kubectl delete pvc --all
#kubectl delete pv --all

echo "Deleting config map"
kubectl delete cm prometheus-config

echo "Deleting containers"
docker container prune

echo "Done!"
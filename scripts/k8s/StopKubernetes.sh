echo "Stopping"

echo "Deleting deployments"
kubectl delete deployments go-app
kubectl delete deployments go-app-dynamodb
kubectl delete deployments go-app-mysql
kubectl delete deployments go-app-redis

echo "Deleting services"
kubectl delete svc go-app
kubectl delete svc go-app-dynamodb
kubectl delete svc go-app-mysql
kubectl delete svc go-app-redis

echo "Done!"
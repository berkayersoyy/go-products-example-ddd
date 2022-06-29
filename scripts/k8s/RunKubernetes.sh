echo "Running deploy"
echo "Building docker image"
#Docker build -t berkayersoyy/go-products-example-ddd .

#echo "Creating config map for prometheus"
#kubectl create configmap prometheus-config --from-file prometheus/prometheus-k8s.yml

echo "Deploying via customization"
kubectl apply -k k8s/


echo "Done!"
echo "Check the port go-products-example-ddd-svc"
kubectl get svc
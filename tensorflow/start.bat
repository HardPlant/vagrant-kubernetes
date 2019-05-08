kubectl apply -f tensorflow_deployment.yml
kubectl expose deployment tf-deployment --type=NodePort --name=tf-service
minikube service tf-service --url
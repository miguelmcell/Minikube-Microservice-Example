sudo kubectl delete deployment backend
sudo docker build -t backend:latest .
sudo docker tag backend localhost:5000/backend
sudo docker push localhost:5000/backend
sudo kubectl create -f backend-deployment.yaml

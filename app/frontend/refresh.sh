sudo kubectl delete deployment.apps/frontend
./build.sh
sudo kubectl create -f frontend-deployment.yaml

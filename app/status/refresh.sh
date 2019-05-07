sudo kubectl delete deployment.apps/status
./build.sh
sudo kubectl create -f status-deployment.yaml

sudo kubectl port-forward --namespace kube-system $(sudo kubectl get po -n kube-system | grep kube-registry-v0 | \awk '{print $1;}') 5000:5000

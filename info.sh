#!/usr/bin/env bash
kubectl get pods -n kube-system | grep nginx-ingress-controller
kubectl get svc -n nash
kubectl get pods -n nash -o wide
kubectl get ingress -n nash -o wide

#Note you need to change the POD ID's
#kubectl logs -n kube-system nginx-ingress-controller-6fc5bcc8c9-thddg -f
#kubectl exec -it -n kube-system nginx-ingress-controller-6fc5bcc8c9-thddg cat /etc/nginx/nginx.conf
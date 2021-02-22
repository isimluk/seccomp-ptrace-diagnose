# A diagnose seccomp-ptrace within kubernetes

```
curl https://raw.githubusercontent.com/isimluk/seccomp-ptrace-diagnose/main/deployment.yaml \
    | kubectl apply -f -

kubectl logs seccomp-ptrace-diagnose > report.txt

curl https://raw.githubusercontent.com/isimluk/seccomp-ptrace-diagnose/main/deployment.yaml \
    | kubectl delete -f -
```

Apply the deployment:

```
kubectl apply -f deployment.yaml
```

Check it:

```bash
kubectl logs deploy/httpget
```

Forward the deployment:

```bash
kubectl port-forward deploy/httpget  8080:8080 &
```

Trigger a HTTP get

```bash
curl -4 "http://127.0.0.1:8080/?q=https://inlets.dev/"
```


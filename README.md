Apply the deployment:

```
kubectl apply -f deployment.yaml
```

Forward the deployment:

```bash
kubectl port-forward 8080:8080 deploy/http-get
```

Trigger a HTTP get

```bash
curl http://127.0.0.1:8080/?q=https://inlets.dev/
```


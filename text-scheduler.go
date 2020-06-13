apiVersion: apps/v1
kind: Deployment
metadata:
  name: pause
spec:
  replicas: 10
  selector:
    matchLabels:
      app: pause
  template:
    metadata:
      labels:
        app: pause
    spec:
      containers:
      - name: pause
        image: gcr.azk8s.cn/google_containers/pause:3.1

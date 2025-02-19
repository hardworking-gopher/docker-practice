docker build -t pandaemoniumplaza/docker-practice:latest pandaemoniumplaza/docker-practice:$SHA .

docker push pandaemoniumplaza/docker-practice:latest pandaemoniumplaza/docker-practice:$SHA

kubectl apply -f k8s

kubectl set image deployments/app-deployment visits-app=pandaemoniumplaza/docker-practice:$SHA

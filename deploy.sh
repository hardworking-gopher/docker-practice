docker build -t pandaemoniumplaza/docker-practice:latest -t pandaemoniumplaza/docker-practice:$SHA .

docker push pandaemoniumplaza/docker-practice:latest
docker push pandaemoniumplaza/docker-practice:$SHA

kubectl apply -f k8s

kubectl set image deployments/app-deployment visits-app=pandaemoniumplaza/docker-practice:$SHA

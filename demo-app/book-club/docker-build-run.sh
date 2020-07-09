docker build -t fasttrackistio/bookclub:1.0.0 .
docker run -p 8082:8080 -it --detach --name bookclub fasttrackistio/bookclub:1.0.0
docker push fasttrackistio/bookclub:1.0.0
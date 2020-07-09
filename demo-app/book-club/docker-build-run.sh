docker build -t rahulrai/bookclub:1.0.0 .
docker run -p 8082:8080 -it --detach --name bookclub rahulrai/bookclub:1.0.0
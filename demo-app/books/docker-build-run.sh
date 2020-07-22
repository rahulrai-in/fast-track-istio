docker build -t fasttrackistio/books:1.0.0 .
docker run --rm -it -p 8081:80 fasttrackistio/books:1.0.0
docker push fasttrackistio/books:1.0.0
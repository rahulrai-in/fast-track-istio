docker build -t fasttrackistio/legacy-books:1.0.0 .
docker run --rm -it -p 8081:80 fasttrackistio/legacy-books:1.0.0
docker push fasttrackistio/legacy-books:1.0.0
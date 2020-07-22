docker build -t fasttrackistio/movies:1.0.0 .
docker run --rm -it -p 8080:8080 fasttrackistio/movies:1.0.0
docker push fasttrackistio/movies:1.0.0
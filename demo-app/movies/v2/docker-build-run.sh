docker build -t fasttrackistio/movies:2.0.0 .
docker run --rm -it -p 8080:8080 fasttrackistio/movies:2.0.0
docker push fasttrackistio/movies:2.0.0
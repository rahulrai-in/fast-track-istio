docker build -t fasttrackistio/independent:1.0.0 .
docker run --rm -it -p 7000:7000 fasttrackistio/independent:1.0.0
docker push fasttrackistio/independent:1.0.0
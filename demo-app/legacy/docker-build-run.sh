docker build -t rahulrai/legacy-books:1.0.0 .
docker run --rm -it -p 8081:80 rahulrai/legacy-books:1.0.0
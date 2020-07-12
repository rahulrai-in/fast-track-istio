#!/bin/bash

openssl req -x509 -sha256 -newkey rsa:4096 -keyout fast-track-istio.ca.key -out fast-track-istio.ca.crt -days 1825 -nodes -subj '/CN=Fast Track Istio'
openssl req -new -newkey rsa:4096 -keyout books.fast-track-istio.key -out books.fast-track-istio.csr -nodes -subj '/CN=books.fast-track-istio.io'
openssl x509 -req -sha256 -days 1825 -in books.fast-track-istio.csr -CA fast-track-istio.ca.crt -CAkey fast-track-istio.ca.key -set_serial 01 -out books.fast-track-istio.crt
openssl req -new -newkey rsa:4096 -keyout client.books.fast-track-istio.key -out client.books.fast-track-istio.csr -nodes -subj '/CN=Book Club Client'
openssl x509 -req -sha256 -days 1825 -in client.books.fast-track-istio.csr -CA fast-track-istio.ca.crt -CAkey fast-track-istio.ca.key -set_serial 02 -out client.books.fast-track-istio.crt

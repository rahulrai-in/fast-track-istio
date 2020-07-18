"use strict";
var restify = require("restify");
const fs = require("fs");

let moviesDataFile = fs.readFileSync("movies.json");

function getMovies(req, res, next) {
  var moviesDataResponse = JSON.parse(moviesDataFile);
  moviesDataResponse.forEach((md) => {
    md["thumbnail"] = fs.readFileSync(md["thumbnail"], { encoding: "base64" });
    md["summary"] = "";
  });

  res.json(moviesDataResponse);
}

function getMovie(req, res, next) {
  var movie = JSON.parse(moviesDataFile).find((m) => m.id == req.params.id);
  movie["thumbnail"] = fs.readFileSync(movie["thumbnail"], {
    encoding: "base64",
  });
  res.json(movie);
}

function hello(req, res, next) {
  console.log(
    "Hello request received by v1 from %s",
    req.connection.remoteAddress
  );
  res.send("Hello from v1 at " + new Date(Date.now()).toString());
}

var server = restify.createServer();
server.pre(restify.pre.sanitizePath());

server.get("/movies", getMovies);
server.get("/movies/:id", getMovie);
server.get("/", hello);

server.listen(8080, function () {
  console.log("%s listening at %s", server.name, server.url);
});

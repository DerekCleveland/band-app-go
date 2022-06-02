# Mongo Dockerfile Setup/Configuration

This directory contains the Dockerfile for the MongoDB container as well as the scripts to generate a fresh MongoDB database.

## Change Directory

For the following commands to work, you must be in the directory containing this file.

```bash
cd deploy/mongo
```

## Building the MongoDB Container

```bash
docker build -t band-app-mongo:0.0.0 .
```

## Running the MongoDB Container

```bash
docker run -p 27017:27017 --name band-app-mongo -d band-app-mongo:0.0.0
```

## Stopping the MongoDB Container

```bash
docker stop band-app-mongo
```

## Removing the MongoDB Container

```bash
docker rm band-app-mongo
```

## Starting the MongoDB Container

```bash
docker start band-app-mongo
```

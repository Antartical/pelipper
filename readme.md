[![Build Status](https://travis-ci.com/Antartical/pelipper.svg?branch=master)](https://travis-ci.com/Antartical/pelipper)
[![Coverage Status](https://coveralls.io/repos/github/Antartical/pelipper/badge.svg?branch=alvarogf97/add_ci)](https://coveralls.io/github/Antartical/pelipper?branch=alvarogf97/add_ci)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

<p align="center">
  <img width="250" height="250" src="https://i.imgur.com/JTRP3z8.jpg">
</p>

## Pelipper in a nutshell

In the first pokemon mistery dugueon, pelipper was the one who deliver notices to the
rescue team, therefore, this service accomplish this task. Pelipper is a service trought the
one you can deliver email and phones notifications letting your other services to forget
about this tedious task.


## Development guide

1. Create your template in the `templates` folder.
2. Write your validators for both, the endpoint and the template data in `validators`.
3. Make the controller that will handler the POST request in `controllers`.
4. Register the created controller in the router in `routes`
5. Build tests for your created code

## Local development

Pelipper is easy to develop in a local environment by using docker. just type in your terminal `make`
and everything you need will make up by itselt. Please copy the content of `build/env/.env.sample` to
your own *.env* in `build/env/.env`. You can do this by executting:
```cmd
cp ./build/env/.env.sample ./build/env/.env
```

Moreover you can perform the following operations:
 - **make sh**: attach a console inside pelipper.
 - **make logs**: show pelipper logs
 - **make local.build**: recompiles pelipper image

## Configure pre-commit (Python3 required)
pre-commit is a useful tool which checks your files before any commit push preventings fails in early steps.

Install pre-commit is easy:
```
pip install pre-commit
python3 -m pre_commit install
```

## How to setup Pelipper as a dependency in your own docker-compose

Just include the following code in your `docker-compose.yml`

```docker
mailhog:
  image: mailhog/mailhog
  container_name: mailhog
  ports: 
    - 1025:1025
    - 8025:8025

pelipper:
  image: ghcr.io/antartical/pelipper
  container_name: pelipper
  ports:
    - "9000:9000"
  environment:
    - SMTP_HOST=mailhog
    - SMTP_PORT=1025
    - SMTP_USER=admin
    - SMTP_PASSWORD=admin
```

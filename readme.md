<p align="center">
  <img width="250" height="250" src="https://ih1.redbubble.net/image.511501561.2977/st,small,507x507-pad,600x600,f8f8f8.u7.jpg">
</p>

## Pelipper in a nutshell

In the first pokemonm mistery dugueon, pelipper was the one who deliver notices to the
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

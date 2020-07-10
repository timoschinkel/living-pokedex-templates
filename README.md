# Living Pokédex templates

I'm not ashamed to admit I still like to play [Pokémon](https://www.pokemon.com/us/). And as I'm a bit of a [completionist](https://www.dictionary.com/browse/completionist) I like to achieve a [living Pokédex](https://bulbapedia.bulbagarden.net/wiki/Living_Pok%C3%A9dex) in every game I play.

This repository is a helping hand in organizing my living Pokédex. For now the only supported Pokédex is for the [Galar region](https://bulbapedia.bulbagarden.net/wiki/Galar). I hope to expand this to all regions later.

The templates are available via https://timoschinkel.github.io/living-pokedex-templates/

The data is based on the [veekun/pokedex](https://github.com/veekun/pokedex), except for the Galar data as this is not yet available in Veekun.

## Building the templates
As an experiment the generation process is built using [Go](https://golang.org).

The process of building the templates consists of two steps; generating JSON and generating HTML. The JSON files are committed to Git in the folder `/data`. This step should only be necessary when new content is made available.

Before running any of the generators the application needs to be compiled. In order to prevent the need to install Go locally a Docker container is available:

```bash
docker-compose run go-cli go build -i -o generate
```

**NB** Due to the usage of a [CGO](https://golang.org/cmd/cgo/) module for connecting to Sqlite compilation can take about 30 seconds.

### Generating JSON sources
First build the veekun/pokedex database following the [instructions](https://github.com/veekun/pokedex/wiki/Getting-Data). The database should be built to Sqlite. Make sure the resulting Sqlite data is copied to the working directory of this repository. The Sqlite database is dumped to `pokedex/data/pokedex.sqlite` by default by Veekun.

To (re)generate the json files run:

```bash
docker-compose run go-cli ./generate json
```

After completion a batch of JSON files is written to `/data`.

### Generating HTML
[WIP]
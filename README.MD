# NextGO: NextJS + GoLang

Simple GoLang Application with NextJS (Static HTML Exports) as Frontend.

## Usage

This is a Simple Golang Application, you just compile it and run it.

```bash
go build -o nextgo cmd/main.go
./nextgo 
```

- Reminder that you have to Build your front-end application before you build the GoLang Application.

```bash
cd nextjs
npm install # install dependencies (haha node_modules are going to eat your disk space)
cp next.config.build.js next.config.js # to use the build config, since we're using NextJS 13 (experimental) and exporting requires special next.config
npm run export
```

## Development

- you're a developer like me, so go figure it out yourself.
- just a thing

```bash 
cd nextjs
cp next.config.dev.js next.config.js # to replace the development configuration for development (we're using NextJS 13 (experimental) and TurboPack cuz we love it faaaast.)
```

## TODO

- [ ] Add a Database ORM (Most likely [beeorm](https://beeorm.io))
- [ ] add a json/yaml/env/... parser for configuration
- [ ] add a custom logger
- [ ] you want more? create an issue, and I will add it here, and someone (or me) will do it.
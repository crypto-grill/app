## Engineering standards 🤝

Here we describe certain regularities which help to streamline collaboration.

### Linter

Use `golangci-lint` to lint your `Go` code:

```bash
golangci-lint run
```

Use `prettier` to check formatting of all other files in the repository:

```bash
npm i
npm run prettier:check
npm run prettier:fix
```

### Documentation

We use `OpenAPI` specification and `Redocly` to generate static sites.
Compiled docs for `main` branch are available at GitBook after deployment succeeded.

In case you need to compile docs from source, run:

```bash
cd docs
npm i
npm run preview
```

See more in [docs/README.md](./docs/README.md).

### Releasing

We use `Semantic Release`. All you need is to make PR titles to conform with `conventional commits`
and everything else will be done automatically.
No need to push tags, create releases yourself or merge release branches.

# defi-loan
**loan** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Interacting with the blockchain

### Request a Loan

```
loand tx loan request-loan 1000token 100token 1000foocoin 500 --from alice --chain-id loan
```

### Approve a Loan

```
loand tx loan approve-loan 0 --from bob --chain-id loan
```

### Repay a Loan

```
loand tx loan repay-loan 0 --from alice --chain-id loan
```

### Liquidate a Loan

```
loand tx loan request-loan 1000token 100token 1000foocoin 20 --from alice --chain-id loan -y
loand tx loan approve-loan 1 --from bob --chain-id loan -y
loand tx loan liquidate-loan 1 --from bob --chain-id loan -y
```

### List all loans

```
loand q loan list-loan
```

### Query the blockchain to see changes in balances

```
loand q bank balances $(loand keys show alice -a)
```

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/loan@latest! | sudo bash
```
`username/loan` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

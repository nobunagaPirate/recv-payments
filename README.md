# Simple receive payments

Simple app to receive payments in bitcoin. For now, the API is a simple
cli program:

```bash
rcv-btc-payment --wallet <wallet_name> --amount <amount in btc>
```

The program halts until a payment is made. 

### How to use

To see all available commands

```bash
rcv-btc-payment help
```

There are a couple configuration options. For now, if you'd like to change them,
you need to recompile this 

# WasmBFT

![banner](docs/static/img/banner.jpg)

[![Version][version-badge]][version-url]
[![Go version][go-badge]][go-url]
[![License][license-badge]][license-url]
[![Sourcegraph][sg-badge]][sg-url]

## Introduction

WasmBFT is a clone of CometBFT v0.38.2, customized to add functionalities for batching transactions into sets of 25 and displaying this information via port 26657. This enhancement aims to improve the efficiency of transaction processing and zk proof generation.

## Features

- **Transaction Batching**: Groups transactions into batches of 25.
- **API Endpoints**: Provides endpoints to access batch details.

## Using WasmBFT as a Go Module

To use WasmBFT in your Go project, you can add it to your `go.mod` file or replace the default CometBFT module.

### Adding WasmBFT to Your Project

Add the following line to your `go.mod` file:

```plaintext
require github.com/airchains-network/wasmbft v0.0.1
```

Then run:

```bash
go mod tidy
```

### Replacing the Default CometBFT Module

If your project already uses CometBFT and you want to replace it with WasmBFT, add the following replace directive to your `go.mod` file:

```plaintext
replace github.com/cometbft/cometbft => github.com/airchains-network/wasmbft v0.0.1
```

Then run:

```bash
go mod tidy
```

## API Endpoints

WasmBFT provides the following API endpoints to access transaction batch details:

- **Get Pod Count**:
  ```http
  GET http://localhost:26657/tracks_pod_count
  ```
  **Sample Response**:
  ```json
  {
    "jsonrpc": "2.0",
    "id": -1,
    "result": "1"
  }
  ```

- **Get Pod Details**:
  ```http
  GET http://localhost:26657/tracks_get_pod?podNumber=1
  ```
  **Sample Response**:
  ```json
  {
    "jsonrpc": "2.0",
    "id": -1,
    "result": [
      {
        "From": "wasmstation1hdxuy3g2jln7lsq5vu3qvp34877jyy2mkvjgha",
        "To": "wasmstation1p20ppaqaeks77h53xrgw83h6a4twzpxqwnafp3",
        "Amount": "1amf",
        "Gas": "119665",
        "TxHash": "64255eb67ea1a93f9943bcccc0ac3d06265b89c5ebd066e1515aaa312639e4ea",
        "ToBalance": "0amf",
        "FromBalance": "999999998amf",
        "Nonce": "0"
      },
      {
        "From": "wasmstation1hdxuy3g2jln7lsq5vu3qvp34877jyy2mkvjgha",
        "To": "wasmstation1p20ppaqaeks77h53xrgw83h6a4twzpxqwnafp3",
        "Amount": "1stake",
        "Gas": "119665",
        "TxHash": "64255eb67ea1a93f9943bcccc0ac3d06265b89c5ebd066e1515aaa312639e4ea",
        "ToBalance": "0stake",
        "FromBalance": "29999998stake",
        "Nonce": "0"
      },
      {
        "From": "wasmstation1hdxuy3g2jln7lsq5vu3qvp34877jyy2mkvjgha",
        "To": "wasmstation148p9k67uuf6smg30e32ndyf0kq9jxx6k6hat95",
        "Amount": "1amf",
        "Gas": "119665",
        "TxHash": "64255eb67ea1a93f9943bcccc0ac3d06265b89c5ebd066e1515aaa312639e4ea",
        "ToBalance": "0amf",
        "FromBalance": "999999998amf",
        "Nonce": "1"
      },
      {
        "From": "wasmstation1hdxuy3g2jln7lsq5vu3qvp34877jyy2mkvjgha",
        "To": "wasmstation148p9k67uuf6smg30e32ndyf0kq9jxx6k6hat95",
        "Amount": "1stake",
        "Gas": "119665",
        "TxHash": "64255eb67ea1a93f9943bcccc0ac3d06265b89c5ebd066e1515aaa312639e4ea",
        "ToBalance": "0stake",
        "FromBalance": "29999998stake",
        "Nonce": "1"
      }
    ]
  }
  ```

## Contributing

Please follow the [Code of Conduct](./CODE_OF_CONDUCT.md) in all interactions. For contribution guidelines, see [CONTRIBUTING.md](./CONTRIBUTING.md).

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](https://github.com/airchains-network/wasmbft/blob/main/LICENSE) file for details.

## Contact

For any inquiries or issues, please contact us via [Airchains Discord](https://discord.gg/airchains).

[version-badge]: https://img.shields.io/github/v/release/airchains-network/wasmbft.svg
[version-url]: https://github.com/airchains-network/wasmbft/releases/latest
[go-badge]: https://img.shields.io/badge/go-1.20-blue.svg
[go-url]: https://github.com/moovweb/gvm
[license-badge]: https://img.shields.io/github/license/airchains-network/wasmbft.svg
[license-url]: https://github.com/airchains-network/wasmbft/blob/main/LICENSE
[sg-badge]: https://sourcegraph.com/github.com/airchains-network/wasmbft/-/badge.svg
[sg-url]: https://sourcegraph.com/github.com/airchains-network/wasmbft?badge

# KubeNake [![Go Report Card](https://goreportcard.com/badge/github.com/smileisak/kubenake)](https://goreportcard.com/report/github.com/smileisak/kubenake)

KubeNake is a simple filter written in go to decode [Kubernetes](https://kubernetes.io) secrets.

## Install

```bash
go get github.com/smileisak/kubenake
```

## Usage

By passing `kubenake` after a pipelined secret, all base64 encoded strings will be decoded.

```bash
$ ➜ k get secrets example -o yaml | kubenake

apiVersion: v1
kind: Secret
data:
  password: pass
type: Opaque
metadata:
  creationTimestamp: "2018-07-12T15:36:15Z"
  name: example
  namespace: default
```

## Contributing

Feel free to send PRs, i'll be so grateful.

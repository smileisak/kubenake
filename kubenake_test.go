package main

import (
	"testing"

	yaml "gopkg.in/yaml.v2"
)

var data = `
apiVersion: v1
data:
  password: cGFzcw==
kind: Secret
metadata:
  creationTimestamp: 2018-03-06T11:44:45Z
  labels:
    app: postgres-prod-postgresql
    chart: postgresql-0.8.12
    heritage: Tiller
    release: postgres-prod
  name: postgres-prod-postgresql
  namespace: digitalcity-prod
  resourceVersion: "3212388"
type: Opaque`

func TestDecode(t *testing.T) {
	secret := Secret{}
	_ = yaml.Unmarshal([]byte(data), &secret)

	decodedSecret := decode(secret)
	for _, v := range decodedSecret.Data {
		if v.(string) != "pass" {
			t.Error("decode does not work! ")
		}
		break
	}
}

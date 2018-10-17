# goscale
Proxy over multiples TCP connections

# Usage

```
  go get github.com/schweigert/goscale
  go install github.com/schweigert/goscale

  PULLSIZE=3 EP_0=localhost:3030 EP_1=localhost:3031 EP_2=localhost:3032 BIND=0.0.0.0:3000 goscale
```

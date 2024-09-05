module github.com/Haeven/finch

go 1.22

toolchain go1.22.7

require (
	flux v0.0.0-00010101000000-000000000000
	github.com/Haeven/codec v0.0.0-20240905163527-7ea0073cedeb
	github.com/lib/pq v1.10.9
	github.com/twmb/franz-go v1.17.1
	github.com/uptrace/bun v1.2.3
	github.com/uptrace/bun/dialect/pgdialect v1.2.3
	github.com/uptrace/bun/driver/pgdriver v1.2.3
)

require (
	github.com/gomodule/redigo v1.8.4 // indirect
	github.com/googollee/go-socket.io v1.7.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/puzpuzpuz/xsync/v3 v3.4.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/twmb/franz-go/pkg/kmsg v1.8.0 // indirect; indirec
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	mellium.im/sasl v0.3.1 // indirect
)

replace flux => ./

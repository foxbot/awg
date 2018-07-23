// Package awg provides "a worker in go"
//
// awg is designed to be compatible with github.com/foxbot/adidis,
// a distributed discord sharder. It should work with any sharder
// that sends JSON-encoded Discord events over a Redis gateway
// keyed with `exchange:events` - though this can pretty easily
// be adapted to work in your environment.
//
// sub-packages are provided to help make writing a fully-fledged
// bot with this worker easier.
package awg

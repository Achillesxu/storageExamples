// Package rs
// Time    : 2022/7/12 21:00
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package rs

const (
	DATA_SHARDS     = 4
	PARITY_SHARDS   = 2
	ALL_SHARDS      = DATA_SHARDS + PARITY_SHARDS
	BLOCK_PER_SHARD = 8000
	BLOCK_SIZE      = BLOCK_PER_SHARD * DATA_SHARDS
)

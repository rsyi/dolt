// Copyright 2020 Liquidata, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nbs

import (
	"context"
	"io"

	"github.com/liquidata-inc/dolt/go/store/chunks"

	"github.com/liquidata-inc/dolt/go/store/hash"
)

// NBSMetricWrapper is a ChunkStore implementation that wraps a ChunkStore, and collects metrics on the calls.
type NBSMetricWrapper struct {
	*chunks.CSMetricWrapper
	nbs *NomsBlockStore
}

// NewCSMetricWrapper returns a new NBSMetricWrapper
func NewNBSMetricWrapper(nbs *NomsBlockStore) *NBSMetricWrapper {
	csMW := chunks.NewCSMetricWrapper(nbs)
	return &NBSMetricWrapper{
		csMW,
		nbs,
	}
}

// Sources retrieves the current root hash, and a list of all the table files
func (nbsMW *NBSMetricWrapper) Sources(ctx context.Context) (hash.Hash, []TableFile, error) {
	return nbsMW.nbs.Sources(ctx)
}

// WriteTableFile will read a table file from the provided reader and write it to the TableFileStore
func (nbsMW *NBSMetricWrapper) WriteTableFile(ctx context.Context, fileId string, numChunks int, rd io.Reader, contentLength uint64, contentHash []byte) error {
	return nbsMW.nbs.WriteTableFile(ctx, fileId, numChunks, rd, contentLength, contentHash)
}

// SetRootChunk changes the root chunk hash from the previous value to the new root.
func (nbsMW *NBSMetricWrapper) SetRootChunk(ctx context.Context, root, previous hash.Hash) error {
	return nbsMW.nbs.SetRootChunk(ctx, root, previous)
}

// GetManyCompressed gets the compressed Chunks with |hashes| from the store. On return,
// |foundChunks| will have been fully sent all chunks which have been
// found. Any non-present chunks will silently be ignored.
func (nbsMW *NBSMetricWrapper) GetManyCompressed(ctx context.Context, hashes hash.HashSet, cmpChChan chan<- CompressedChunk) error {
	nbsMW.TotalChunkGets += len(hashes)
	for h := range hashes {
		nbsMW.UniqueGets.Insert(h)
	}

	return nbsMW.nbs.GetManyCompressed(ctx, hashes, cmpChChan)
}

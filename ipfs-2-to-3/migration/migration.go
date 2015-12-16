package mg1

import (
	"encoding/json"
	"fmt"
	"path"

	migrate "github.com/ipfs/fs-repo-migrations/go-migrate"
	lock "github.com/ipfs/fs-repo-migrations/ipfs-1-to-2/repolock"
	bst "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/blocks/blockstore"
	bs "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/blockservice"
	offline "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/exchange/offline"
	dag "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/merkledag"
	newpin "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/pin"
	u "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/util"
	dstore "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore"
	flatfs "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore/flatfs"
	leveldb "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore/leveldb"
	mount "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore/mount"
	dsq "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore/query"
	sync "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore/sync"
	mfsr "github.com/ipfs/fs-repo-migrations/mfsr"
	log "github.com/ipfs/fs-repo-migrations/stump"
)

var recursePinDatastoreKey = dstore.NewKey("/local/pins/recursive/keys")
var directPinDatastoreKey = dstore.NewKey("/local/pins/direct/keys")
var indirectPinDatastoreKey = dstore.NewKey("/local/pins/indirect/keys")

type Migration struct{}

func (m Migration) Versions() string {
	return "2-to-3"
}

func (m Migration) Reversible() bool {
	return true
}

func (m Migration) Apply(opts migrate.Options) error {
	log.Verbose = opts.Verbose
	log.Log("applying %s repo migration", m.Versions())

	log.VLog("locking repo at %q", opts.Path)
	lk, err := lock.Lock2(opts.Path)
	if err != nil {
		return err
	}
	defer lk.Close()

	repo := mfsr.RepoPath(opts.Path)

	log.VLog("  - verifying version is '2'")
	if err := repo.CheckVersion("2"); err != nil {
		return err
	}

	if err := transferPins(opts.Path, opts.Verbose); err != nil {
		return err
	}

	log.Log("pin transfer completed successfuly")

	err = repo.WriteVersion("3")
	if err != nil {
		return err
	}
	log.Log("updated version file")

	return nil
}

func (m Migration) Revert(opts migrate.Options) error {
	lk, err := lock.Lock2(opts.Path)
	if err != nil {
		return err
	}
	defer lk.Close()

	repo := mfsr.RepoPath(opts.Path)
	if err := repo.CheckVersion("3"); err != nil {
		return err
	}

	if err := revertPins(opts.Path, opts.Verbose); err != nil {
		return err
	}

	// 3) change version number back down
	err = repo.WriteVersion("2")
	if err != nil {
		return err
	}
	if opts.Verbose {
		fmt.Println("lowered version number to 2")
	}

	return nil
}

func openDatastore(repopath string) (dstore.ThreadSafeDatastore, error) {
	log.VLog("  - opening datastore at %q", repopath)
	ldbpath := path.Join(repopath, "datastore")
	ldb, err := leveldb.NewDatastore(ldbpath, nil)
	if err != nil {
		return nil, err
	}

	blockspath := path.Join(repopath, "blocks")
	fds, err := flatfs.New(blockspath, 4)
	if err != nil {
		return nil, err
	}

	return sync.MutexWrap(mount.New([]mount.Mount{
		{
			Prefix:    dstore.NewKey("/blocks"),
			Datastore: fds,
		},
		{
			Prefix:    dstore.NewKey("/"),
			Datastore: ldb,
		},
	})), nil
}

func constructDagServ(ds dstore.ThreadSafeDatastore) (dag.DAGService, error) {
	bstore := bst.NewBlockstore(ds)
	bserve, err := bs.New(bstore, offline.Exchange(bstore))
	if err != nil {
		return nil, err
	}
	return dag.NewDAGService(bserve), nil
}

func transferPins(repopath string, verbose bool) error {
	log.Log("beginning pin transfer")
	ds, err := openDatastore(repopath)
	if err != nil {
		return err
	}

	dserv, err := constructDagServ(ds)
	if err != nil {
		return err
	}

	pinner := newpin.NewPinner(ds, dserv)
	log.VLog("  - created version 3 pinner")

	log.VLog("  - loading recursive pins")
	recKeys, err := loadOldKeys(ds, recursePinDatastoreKey)
	if err != nil {
		return err
	}

	for _, k := range recKeys {
		pinner.PinWithMode(k, newpin.Recursive)
	}
	log.VLog("  - transfered recursive pins")

	log.VLog("  - loading direct pins")
	dirKeys, err := loadOldKeys(ds, directPinDatastoreKey)
	if err != nil {
		return err
	}

	for _, k := range dirKeys {
		pinner.PinWithMode(k, newpin.Direct)
	}
	log.VLog("  - transfered direct pins")

	err = pinner.Flush()
	if err != nil {
		return err
	}
	log.Log("pinner synced to disk")

	return cleanupOldPins(ds, verbose)
}

func cleanupOldPins(ds dstore.Datastore, verbose bool) error {
	log.Log("cleaning old pins")
	err := cleanupKeyspace(ds, recursePinDatastoreKey)
	if err != nil {
		return err
	}
	log.VLog("  - cleaned up oldstyle recursive pins")

	err = cleanupKeyspace(ds, directPinDatastoreKey)
	if err != nil {
		return err
	}
	log.VLog("  - cleaned up oldstyle direct pins")

	err = cleanupKeyspace(ds, indirectPinDatastoreKey)
	if err != nil {
		return err
	}
	log.VLog("  - cleaned up oldstyle indirect pins")

	return nil
}

func cleanupKeyspace(ds dstore.Datastore, k dstore.Key) error {
	log.VLog("deleting recursePin root key:", recursePinDatastoreKey)
	err := ds.Delete(recursePinDatastoreKey)
	if err != nil {
		return err
	}

	res, err := ds.Query(dsq.Query{
		Prefix:   recursePinDatastoreKey.String(),
		KeysOnly: true,
	})
	if err != nil {
		return err
	}
	for k := range res.Next() {
		log.VLog("deleting pin key:", k.Key)
		err := ds.Delete(dstore.NewKey(k.Key))
		if err != nil {
			res.Close()
			return err
		}
	}

	return nil
}

func revertPins(repopath string, verbose bool) error {
	ds, err := openDatastore(repopath)
	if err != nil {
		return err
	}

	dserv, err := constructDagServ(ds)
	if err != nil {
		return err
	}

	pinner, err := newpin.LoadPinner(ds, dserv)
	if err != nil {
		return err
	}

	if err := writeOldKeys(ds, recursePinDatastoreKey, pinner.RecursiveKeys()); err != nil {
		return err
	}

	if err := writeOldKeys(ds, directPinDatastoreKey, pinner.DirectKeys()); err != nil {
		return err
	}

	if err := writeOldIndirPins(ds, indirectPinDatastoreKey, pinner.IndirectKeys()); err != nil {
		return err
	}

	return nil
}

func writeOldIndirPins(to dstore.Datastore, k dstore.Key, pins map[u.Key]uint64) error {
	refs := make(map[string]int)
	for k, v := range pins {
		refs[k.String()] = int(v)
	}

	b, err := json.Marshal(refs)
	if err != nil {
		return err
	}

	return to.Put(k, b)
}

func loadOldKeys(from dstore.Datastore, k dstore.Key) ([]u.Key, error) {
	v, err := from.Get(k)
	if err != nil {
		return nil, err
	}

	vb, ok := v.([]byte)
	if !ok {
		return nil, fmt.Errorf("pinset %s was not stored as []byte")
	}

	var keys []u.Key
	err = json.Unmarshal(vb, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func writeOldKeys(to dstore.Datastore, k dstore.Key, pins []u.Key) error {
	log.Log("writing keys: ", k, pins)
	b, err := json.Marshal(pins)
	if err != nil {
		return err
	}

	return to.Put(k, b)
}

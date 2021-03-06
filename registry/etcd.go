package registry

import (
	"context"
	"encoding/json"
	"errors"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/ofavor/micro-lite/internal/log"
)

var (

	// ErrNotFound error when GetService is called
	ErrNotFound = errors.New("service not found")

	prefix        = "/micro/registry/"
	clientTimeout = 3 * time.Second
)

type etcdRegistry struct {
	sync.RWMutex
	opts Options

	client *clientv3.Client
}

func newETCDRegistry(opts ...Option) Registry {
	options := defaultOptions()
	for _, o := range opts {
		o(&options)
	}

	return &etcdRegistry{
		opts: options,
	}
}

func (r *etcdRegistry) getClient() (*clientv3.Client, error) {
	r.RLock()
	if r.client != nil {
		r.RUnlock()
		return r.client, nil
	}
	r.RUnlock()

	r.Lock()
	defer r.Unlock()
	cfg := clientv3.Config{
		Endpoints: r.opts.Addrs,
	}
	// log.Debug("Etcd registry config:", cfg)

	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Error("Create ectd client error: ", err)
		return nil, err
	}
	r.client = cli
	return cli, nil
}

func encode(s *Service) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func decode(ds []byte) *Service {
	var s *Service
	json.Unmarshal(ds, &s)
	return s
}

func nodePath(s, id string) string {
	service := strings.Replace(s, "/", "-", -1)
	node := strings.Replace(id, "/", "-", -1)
	return path.Join(prefix, service, node)
}

func servicePath(s string) string {
	return path.Join(prefix, strings.Replace(s, "/", "-", -1))
}

func (r *etcdRegistry) registerNode(svc *Service, node *Node) error {
	cli, err := r.getClient()
	if err != nil {
		return err
	}

	service := &Service{
		Name:      svc.Name,
		Version:   svc.Version,
		Metadata:  svc.Metadata,
		Endpoints: svc.Endpoints,
		Nodes:     []*Node{node},
	}
	key := nodePath(service.Name, node.ID)
	val := encode(service)
	log.Debugf("Register service node: %s", key)
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
	defer cancel()
	lgr, err := cli.Grant(ctx, int64(r.opts.TTL.Seconds()))
	_, err = cli.Put(ctx, key, val, clientv3.WithLease(lgr.ID))
	return err
}

func (r *etcdRegistry) Init(o Option) {
	o(&r.opts)
}

func (r *etcdRegistry) Register(svc *Service, opts ...Option) error {
	for _, o := range opts {
		o(&r.opts)
	}
	for _, n := range svc.Nodes {
		if err := r.registerNode(svc, n); err != nil {
			return err
		}
	}
	return nil
}

func (r *etcdRegistry) Deregister(svc *Service) error {
	cli, err := r.getClient()
	if err != nil {
		return err
	}
	for _, n := range svc.Nodes {
		ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
		defer cancel()

		path := nodePath(svc.Name, n.ID)
		log.Debugf("Deregistering service node: %s", path)
		_, err := cli.Delete(ctx, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *etcdRegistry) GetService(name string) ([]*Service, error) {
	cli, err := r.getClient()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
	defer cancel()

	rsp, err := cli.Get(ctx, servicePath(name)+"/", clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	if len(rsp.Kvs) == 0 {
		return nil, ErrNotFound
	}

	serviceMap := map[string]*Service{}

	for _, n := range rsp.Kvs {
		if sn := decode(n.Value); sn != nil {
			s, ok := serviceMap[sn.Version]
			if !ok {
				s = &Service{
					Name:      sn.Name,
					Version:   sn.Version,
					Metadata:  sn.Metadata,
					Endpoints: sn.Endpoints,
				}
				serviceMap[s.Version] = s
			}

			s.Nodes = append(s.Nodes, sn.Nodes...)
		}
	}

	services := make([]*Service, 0, len(serviceMap))
	for _, service := range serviceMap {
		services = append(services, service)
	}

	return services, nil
}

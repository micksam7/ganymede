// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/chapter"
	"github.com/zibbp/ganymede/ent/mutedsegment"
	"github.com/zibbp/ganymede/ent/playlist"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/ent/vod"
)

// VodQuery is the builder for querying Vod entities.
type VodQuery struct {
	config
	ctx               *QueryContext
	order             []vod.OrderOption
	inters            []Interceptor
	predicates        []predicate.Vod
	withChannel       *ChannelQuery
	withQueue         *QueueQuery
	withPlaylists     *PlaylistQuery
	withChapters      *ChapterQuery
	withMutedSegments *MutedSegmentQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VodQuery builder.
func (vq *VodQuery) Where(ps ...predicate.Vod) *VodQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VodQuery) Limit(limit int) *VodQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VodQuery) Offset(offset int) *VodQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VodQuery) Unique(unique bool) *VodQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VodQuery) Order(o ...vod.OrderOption) *VodQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryChannel chains the current query on the "channel" edge.
func (vq *VodQuery) QueryChannel() *ChannelQuery {
	query := (&ChannelClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vod.ChannelTable, vod.ChannelColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQueue chains the current query on the "queue" edge.
func (vq *VodQuery) QueryQueue() *QueueQuery {
	query := (&QueueClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, selector),
			sqlgraph.To(queue.Table, queue.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, vod.QueueTable, vod.QueueColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlaylists chains the current query on the "playlists" edge.
func (vq *VodQuery) QueryPlaylists() *PlaylistQuery {
	query := (&PlaylistClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, selector),
			sqlgraph.To(playlist.Table, playlist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, vod.PlaylistsTable, vod.PlaylistsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChapters chains the current query on the "chapters" edge.
func (vq *VodQuery) QueryChapters() *ChapterQuery {
	query := (&ChapterClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, selector),
			sqlgraph.To(chapter.Table, chapter.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vod.ChaptersTable, vod.ChaptersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMutedSegments chains the current query on the "muted_segments" edge.
func (vq *VodQuery) QueryMutedSegments() *MutedSegmentQuery {
	query := (&MutedSegmentClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vod.Table, vod.FieldID, selector),
			sqlgraph.To(mutedsegment.Table, mutedsegment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vod.MutedSegmentsTable, vod.MutedSegmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Vod entity from the query.
// Returns a *NotFoundError when no Vod was found.
func (vq *VodQuery) First(ctx context.Context) (*Vod, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vod.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VodQuery) FirstX(ctx context.Context) *Vod {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vod ID from the query.
// Returns a *NotFoundError when no Vod ID was found.
func (vq *VodQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vod.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VodQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vod entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vod entity is found.
// Returns a *NotFoundError when no Vod entities are found.
func (vq *VodQuery) Only(ctx context.Context) (*Vod, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vod.Label}
	default:
		return nil, &NotSingularError{vod.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VodQuery) OnlyX(ctx context.Context) *Vod {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vod ID in the query.
// Returns a *NotSingularError when more than one Vod ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VodQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vod.Label}
	default:
		err = &NotSingularError{vod.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VodQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vods.
func (vq *VodQuery) All(ctx context.Context) ([]*Vod, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Vod, *VodQuery]()
	return withInterceptors[[]*Vod](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VodQuery) AllX(ctx context.Context) []*Vod {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vod IDs.
func (vq *VodQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(vod.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VodQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VodQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VodQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VodQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VodQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VodQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VodQuery) Clone() *VodQuery {
	if vq == nil {
		return nil
	}
	return &VodQuery{
		config:            vq.config,
		ctx:               vq.ctx.Clone(),
		order:             append([]vod.OrderOption{}, vq.order...),
		inters:            append([]Interceptor{}, vq.inters...),
		predicates:        append([]predicate.Vod{}, vq.predicates...),
		withChannel:       vq.withChannel.Clone(),
		withQueue:         vq.withQueue.Clone(),
		withPlaylists:     vq.withPlaylists.Clone(),
		withChapters:      vq.withChapters.Clone(),
		withMutedSegments: vq.withMutedSegments.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithChannel tells the query-builder to eager-load the nodes that are connected to
// the "channel" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VodQuery) WithChannel(opts ...func(*ChannelQuery)) *VodQuery {
	query := (&ChannelClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withChannel = query
	return vq
}

// WithQueue tells the query-builder to eager-load the nodes that are connected to
// the "queue" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VodQuery) WithQueue(opts ...func(*QueueQuery)) *VodQuery {
	query := (&QueueClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withQueue = query
	return vq
}

// WithPlaylists tells the query-builder to eager-load the nodes that are connected to
// the "playlists" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VodQuery) WithPlaylists(opts ...func(*PlaylistQuery)) *VodQuery {
	query := (&PlaylistClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withPlaylists = query
	return vq
}

// WithChapters tells the query-builder to eager-load the nodes that are connected to
// the "chapters" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VodQuery) WithChapters(opts ...func(*ChapterQuery)) *VodQuery {
	query := (&ChapterClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withChapters = query
	return vq
}

// WithMutedSegments tells the query-builder to eager-load the nodes that are connected to
// the "muted_segments" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VodQuery) WithMutedSegments(opts ...func(*MutedSegmentQuery)) *VodQuery {
	query := (&MutedSegmentClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withMutedSegments = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExtID string `json:"ext_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vod.Query().
//		GroupBy(vod.FieldExtID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VodQuery) GroupBy(field string, fields ...string) *VodGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VodGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = vod.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExtID string `json:"ext_id,omitempty"`
//	}
//
//	client.Vod.Query().
//		Select(vod.FieldExtID).
//		Scan(ctx, &v)
func (vq *VodQuery) Select(fields ...string) *VodSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VodSelect{VodQuery: vq}
	sbuild.label = vod.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VodSelect configured with the given aggregations.
func (vq *VodQuery) Aggregate(fns ...AggregateFunc) *VodSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VodQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !vod.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vod, error) {
	var (
		nodes       = []*Vod{}
		withFKs     = vq.withFKs
		_spec       = vq.querySpec()
		loadedTypes = [5]bool{
			vq.withChannel != nil,
			vq.withQueue != nil,
			vq.withPlaylists != nil,
			vq.withChapters != nil,
			vq.withMutedSegments != nil,
		}
	)
	if vq.withChannel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, vod.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Vod).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Vod{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withChannel; query != nil {
		if err := vq.loadChannel(ctx, query, nodes, nil,
			func(n *Vod, e *Channel) { n.Edges.Channel = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withQueue; query != nil {
		if err := vq.loadQueue(ctx, query, nodes, nil,
			func(n *Vod, e *Queue) { n.Edges.Queue = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withPlaylists; query != nil {
		if err := vq.loadPlaylists(ctx, query, nodes,
			func(n *Vod) { n.Edges.Playlists = []*Playlist{} },
			func(n *Vod, e *Playlist) { n.Edges.Playlists = append(n.Edges.Playlists, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withChapters; query != nil {
		if err := vq.loadChapters(ctx, query, nodes,
			func(n *Vod) { n.Edges.Chapters = []*Chapter{} },
			func(n *Vod, e *Chapter) { n.Edges.Chapters = append(n.Edges.Chapters, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withMutedSegments; query != nil {
		if err := vq.loadMutedSegments(ctx, query, nodes,
			func(n *Vod) { n.Edges.MutedSegments = []*MutedSegment{} },
			func(n *Vod, e *MutedSegment) { n.Edges.MutedSegments = append(n.Edges.MutedSegments, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VodQuery) loadChannel(ctx context.Context, query *ChannelQuery, nodes []*Vod, init func(*Vod), assign func(*Vod, *Channel)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Vod)
	for i := range nodes {
		if nodes[i].channel_vods == nil {
			continue
		}
		fk := *nodes[i].channel_vods
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(channel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "channel_vods" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VodQuery) loadQueue(ctx context.Context, query *QueueQuery, nodes []*Vod, init func(*Vod), assign func(*Vod, *Queue)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Vod)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Queue(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(vod.QueueColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.vod_queue
		if fk == nil {
			return fmt.Errorf(`foreign-key "vod_queue" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "vod_queue" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VodQuery) loadPlaylists(ctx context.Context, query *PlaylistQuery, nodes []*Vod, init func(*Vod), assign func(*Vod, *Playlist)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Vod)
	nids := make(map[uuid.UUID]map[*Vod]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(vod.PlaylistsTable)
		s.Join(joinT).On(s.C(playlist.FieldID), joinT.C(vod.PlaylistsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(vod.PlaylistsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(vod.PlaylistsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Vod]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Playlist](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "playlists" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (vq *VodQuery) loadChapters(ctx context.Context, query *ChapterQuery, nodes []*Vod, init func(*Vod), assign func(*Vod, *Chapter)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Vod)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Chapter(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(vod.ChaptersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.vod_chapters
		if fk == nil {
			return fmt.Errorf(`foreign-key "vod_chapters" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "vod_chapters" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VodQuery) loadMutedSegments(ctx context.Context, query *MutedSegmentQuery, nodes []*Vod, init func(*Vod), assign func(*Vod, *MutedSegment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Vod)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.MutedSegment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(vod.MutedSegmentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.vod_muted_segments
		if fk == nil {
			return fmt.Errorf(`foreign-key "vod_muted_segments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "vod_muted_segments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vq *VodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(vod.Table, vod.Columns, sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vod.FieldID)
		for i := range fields {
			if fields[i] != vod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vod.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = vod.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VodGroupBy is the group-by builder for Vod entities.
type VodGroupBy struct {
	selector
	build *VodQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VodGroupBy) Aggregate(fns ...AggregateFunc) *VodGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VodGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VodQuery, *VodGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VodGroupBy) sqlScan(ctx context.Context, root *VodQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VodSelect is the builder for selecting fields of Vod entities.
type VodSelect struct {
	*VodQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VodSelect) Aggregate(fns ...AggregateFunc) *VodSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VodSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VodQuery, *VodSelect](ctx, vs.VodQuery, vs, vs.inters, v)
}

func (vs *VodSelect) sqlScan(ctx context.Context, root *VodQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

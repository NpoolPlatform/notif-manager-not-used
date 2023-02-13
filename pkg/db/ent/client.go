// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/notif"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/notifchannel"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/readannouncement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/sendannouncement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/txnotifstate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/userannouncement"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Announcement is the client for interacting with the Announcement builders.
	Announcement *AnnouncementClient
	// Notif is the client for interacting with the Notif builders.
	Notif *NotifClient
	// NotifChannel is the client for interacting with the NotifChannel builders.
	NotifChannel *NotifChannelClient
	// ReadAnnouncement is the client for interacting with the ReadAnnouncement builders.
	ReadAnnouncement *ReadAnnouncementClient
	// SendAnnouncement is the client for interacting with the SendAnnouncement builders.
	SendAnnouncement *SendAnnouncementClient
	// TxNotifState is the client for interacting with the TxNotifState builders.
	TxNotifState *TxNotifStateClient
	// UserAnnouncement is the client for interacting with the UserAnnouncement builders.
	UserAnnouncement *UserAnnouncementClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Announcement = NewAnnouncementClient(c.config)
	c.Notif = NewNotifClient(c.config)
	c.NotifChannel = NewNotifChannelClient(c.config)
	c.ReadAnnouncement = NewReadAnnouncementClient(c.config)
	c.SendAnnouncement = NewSendAnnouncementClient(c.config)
	c.TxNotifState = NewTxNotifStateClient(c.config)
	c.UserAnnouncement = NewUserAnnouncementClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Announcement:     NewAnnouncementClient(cfg),
		Notif:            NewNotifClient(cfg),
		NotifChannel:     NewNotifChannelClient(cfg),
		ReadAnnouncement: NewReadAnnouncementClient(cfg),
		SendAnnouncement: NewSendAnnouncementClient(cfg),
		TxNotifState:     NewTxNotifStateClient(cfg),
		UserAnnouncement: NewUserAnnouncementClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Announcement:     NewAnnouncementClient(cfg),
		Notif:            NewNotifClient(cfg),
		NotifChannel:     NewNotifChannelClient(cfg),
		ReadAnnouncement: NewReadAnnouncementClient(cfg),
		SendAnnouncement: NewSendAnnouncementClient(cfg),
		TxNotifState:     NewTxNotifStateClient(cfg),
		UserAnnouncement: NewUserAnnouncementClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Announcement.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Announcement.Use(hooks...)
	c.Notif.Use(hooks...)
	c.NotifChannel.Use(hooks...)
	c.ReadAnnouncement.Use(hooks...)
	c.SendAnnouncement.Use(hooks...)
	c.TxNotifState.Use(hooks...)
	c.UserAnnouncement.Use(hooks...)
}

// AnnouncementClient is a client for the Announcement schema.
type AnnouncementClient struct {
	config
}

// NewAnnouncementClient returns a client for the Announcement from the given config.
func NewAnnouncementClient(c config) *AnnouncementClient {
	return &AnnouncementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `announcement.Hooks(f(g(h())))`.
func (c *AnnouncementClient) Use(hooks ...Hook) {
	c.hooks.Announcement = append(c.hooks.Announcement, hooks...)
}

// Create returns a builder for creating a Announcement entity.
func (c *AnnouncementClient) Create() *AnnouncementCreate {
	mutation := newAnnouncementMutation(c.config, OpCreate)
	return &AnnouncementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Announcement entities.
func (c *AnnouncementClient) CreateBulk(builders ...*AnnouncementCreate) *AnnouncementCreateBulk {
	return &AnnouncementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Announcement.
func (c *AnnouncementClient) Update() *AnnouncementUpdate {
	mutation := newAnnouncementMutation(c.config, OpUpdate)
	return &AnnouncementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnnouncementClient) UpdateOne(a *Announcement) *AnnouncementUpdateOne {
	mutation := newAnnouncementMutation(c.config, OpUpdateOne, withAnnouncement(a))
	return &AnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnnouncementClient) UpdateOneID(id uuid.UUID) *AnnouncementUpdateOne {
	mutation := newAnnouncementMutation(c.config, OpUpdateOne, withAnnouncementID(id))
	return &AnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Announcement.
func (c *AnnouncementClient) Delete() *AnnouncementDelete {
	mutation := newAnnouncementMutation(c.config, OpDelete)
	return &AnnouncementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AnnouncementClient) DeleteOne(a *Announcement) *AnnouncementDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AnnouncementClient) DeleteOneID(id uuid.UUID) *AnnouncementDeleteOne {
	builder := c.Delete().Where(announcement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnnouncementDeleteOne{builder}
}

// Query returns a query builder for Announcement.
func (c *AnnouncementClient) Query() *AnnouncementQuery {
	return &AnnouncementQuery{
		config: c.config,
	}
}

// Get returns a Announcement entity by its id.
func (c *AnnouncementClient) Get(ctx context.Context, id uuid.UUID) (*Announcement, error) {
	return c.Query().Where(announcement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnnouncementClient) GetX(ctx context.Context, id uuid.UUID) *Announcement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AnnouncementClient) Hooks() []Hook {
	hooks := c.hooks.Announcement
	return append(hooks[:len(hooks):len(hooks)], announcement.Hooks[:]...)
}

// NotifClient is a client for the Notif schema.
type NotifClient struct {
	config
}

// NewNotifClient returns a client for the Notif from the given config.
func NewNotifClient(c config) *NotifClient {
	return &NotifClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notif.Hooks(f(g(h())))`.
func (c *NotifClient) Use(hooks ...Hook) {
	c.hooks.Notif = append(c.hooks.Notif, hooks...)
}

// Create returns a builder for creating a Notif entity.
func (c *NotifClient) Create() *NotifCreate {
	mutation := newNotifMutation(c.config, OpCreate)
	return &NotifCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notif entities.
func (c *NotifClient) CreateBulk(builders ...*NotifCreate) *NotifCreateBulk {
	return &NotifCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notif.
func (c *NotifClient) Update() *NotifUpdate {
	mutation := newNotifMutation(c.config, OpUpdate)
	return &NotifUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotifClient) UpdateOne(n *Notif) *NotifUpdateOne {
	mutation := newNotifMutation(c.config, OpUpdateOne, withNotif(n))
	return &NotifUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotifClient) UpdateOneID(id uuid.UUID) *NotifUpdateOne {
	mutation := newNotifMutation(c.config, OpUpdateOne, withNotifID(id))
	return &NotifUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notif.
func (c *NotifClient) Delete() *NotifDelete {
	mutation := newNotifMutation(c.config, OpDelete)
	return &NotifDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NotifClient) DeleteOne(n *Notif) *NotifDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *NotifClient) DeleteOneID(id uuid.UUID) *NotifDeleteOne {
	builder := c.Delete().Where(notif.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotifDeleteOne{builder}
}

// Query returns a query builder for Notif.
func (c *NotifClient) Query() *NotifQuery {
	return &NotifQuery{
		config: c.config,
	}
}

// Get returns a Notif entity by its id.
func (c *NotifClient) Get(ctx context.Context, id uuid.UUID) (*Notif, error) {
	return c.Query().Where(notif.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotifClient) GetX(ctx context.Context, id uuid.UUID) *Notif {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NotifClient) Hooks() []Hook {
	hooks := c.hooks.Notif
	return append(hooks[:len(hooks):len(hooks)], notif.Hooks[:]...)
}

// NotifChannelClient is a client for the NotifChannel schema.
type NotifChannelClient struct {
	config
}

// NewNotifChannelClient returns a client for the NotifChannel from the given config.
func NewNotifChannelClient(c config) *NotifChannelClient {
	return &NotifChannelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notifchannel.Hooks(f(g(h())))`.
func (c *NotifChannelClient) Use(hooks ...Hook) {
	c.hooks.NotifChannel = append(c.hooks.NotifChannel, hooks...)
}

// Create returns a builder for creating a NotifChannel entity.
func (c *NotifChannelClient) Create() *NotifChannelCreate {
	mutation := newNotifChannelMutation(c.config, OpCreate)
	return &NotifChannelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NotifChannel entities.
func (c *NotifChannelClient) CreateBulk(builders ...*NotifChannelCreate) *NotifChannelCreateBulk {
	return &NotifChannelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NotifChannel.
func (c *NotifChannelClient) Update() *NotifChannelUpdate {
	mutation := newNotifChannelMutation(c.config, OpUpdate)
	return &NotifChannelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotifChannelClient) UpdateOne(nc *NotifChannel) *NotifChannelUpdateOne {
	mutation := newNotifChannelMutation(c.config, OpUpdateOne, withNotifChannel(nc))
	return &NotifChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotifChannelClient) UpdateOneID(id uuid.UUID) *NotifChannelUpdateOne {
	mutation := newNotifChannelMutation(c.config, OpUpdateOne, withNotifChannelID(id))
	return &NotifChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NotifChannel.
func (c *NotifChannelClient) Delete() *NotifChannelDelete {
	mutation := newNotifChannelMutation(c.config, OpDelete)
	return &NotifChannelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NotifChannelClient) DeleteOne(nc *NotifChannel) *NotifChannelDeleteOne {
	return c.DeleteOneID(nc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *NotifChannelClient) DeleteOneID(id uuid.UUID) *NotifChannelDeleteOne {
	builder := c.Delete().Where(notifchannel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotifChannelDeleteOne{builder}
}

// Query returns a query builder for NotifChannel.
func (c *NotifChannelClient) Query() *NotifChannelQuery {
	return &NotifChannelQuery{
		config: c.config,
	}
}

// Get returns a NotifChannel entity by its id.
func (c *NotifChannelClient) Get(ctx context.Context, id uuid.UUID) (*NotifChannel, error) {
	return c.Query().Where(notifchannel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotifChannelClient) GetX(ctx context.Context, id uuid.UUID) *NotifChannel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NotifChannelClient) Hooks() []Hook {
	hooks := c.hooks.NotifChannel
	return append(hooks[:len(hooks):len(hooks)], notifchannel.Hooks[:]...)
}

// ReadAnnouncementClient is a client for the ReadAnnouncement schema.
type ReadAnnouncementClient struct {
	config
}

// NewReadAnnouncementClient returns a client for the ReadAnnouncement from the given config.
func NewReadAnnouncementClient(c config) *ReadAnnouncementClient {
	return &ReadAnnouncementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `readannouncement.Hooks(f(g(h())))`.
func (c *ReadAnnouncementClient) Use(hooks ...Hook) {
	c.hooks.ReadAnnouncement = append(c.hooks.ReadAnnouncement, hooks...)
}

// Create returns a builder for creating a ReadAnnouncement entity.
func (c *ReadAnnouncementClient) Create() *ReadAnnouncementCreate {
	mutation := newReadAnnouncementMutation(c.config, OpCreate)
	return &ReadAnnouncementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ReadAnnouncement entities.
func (c *ReadAnnouncementClient) CreateBulk(builders ...*ReadAnnouncementCreate) *ReadAnnouncementCreateBulk {
	return &ReadAnnouncementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ReadAnnouncement.
func (c *ReadAnnouncementClient) Update() *ReadAnnouncementUpdate {
	mutation := newReadAnnouncementMutation(c.config, OpUpdate)
	return &ReadAnnouncementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReadAnnouncementClient) UpdateOne(ra *ReadAnnouncement) *ReadAnnouncementUpdateOne {
	mutation := newReadAnnouncementMutation(c.config, OpUpdateOne, withReadAnnouncement(ra))
	return &ReadAnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReadAnnouncementClient) UpdateOneID(id uuid.UUID) *ReadAnnouncementUpdateOne {
	mutation := newReadAnnouncementMutation(c.config, OpUpdateOne, withReadAnnouncementID(id))
	return &ReadAnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ReadAnnouncement.
func (c *ReadAnnouncementClient) Delete() *ReadAnnouncementDelete {
	mutation := newReadAnnouncementMutation(c.config, OpDelete)
	return &ReadAnnouncementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReadAnnouncementClient) DeleteOne(ra *ReadAnnouncement) *ReadAnnouncementDeleteOne {
	return c.DeleteOneID(ra.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ReadAnnouncementClient) DeleteOneID(id uuid.UUID) *ReadAnnouncementDeleteOne {
	builder := c.Delete().Where(readannouncement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReadAnnouncementDeleteOne{builder}
}

// Query returns a query builder for ReadAnnouncement.
func (c *ReadAnnouncementClient) Query() *ReadAnnouncementQuery {
	return &ReadAnnouncementQuery{
		config: c.config,
	}
}

// Get returns a ReadAnnouncement entity by its id.
func (c *ReadAnnouncementClient) Get(ctx context.Context, id uuid.UUID) (*ReadAnnouncement, error) {
	return c.Query().Where(readannouncement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReadAnnouncementClient) GetX(ctx context.Context, id uuid.UUID) *ReadAnnouncement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ReadAnnouncementClient) Hooks() []Hook {
	hooks := c.hooks.ReadAnnouncement
	return append(hooks[:len(hooks):len(hooks)], readannouncement.Hooks[:]...)
}

// SendAnnouncementClient is a client for the SendAnnouncement schema.
type SendAnnouncementClient struct {
	config
}

// NewSendAnnouncementClient returns a client for the SendAnnouncement from the given config.
func NewSendAnnouncementClient(c config) *SendAnnouncementClient {
	return &SendAnnouncementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sendannouncement.Hooks(f(g(h())))`.
func (c *SendAnnouncementClient) Use(hooks ...Hook) {
	c.hooks.SendAnnouncement = append(c.hooks.SendAnnouncement, hooks...)
}

// Create returns a builder for creating a SendAnnouncement entity.
func (c *SendAnnouncementClient) Create() *SendAnnouncementCreate {
	mutation := newSendAnnouncementMutation(c.config, OpCreate)
	return &SendAnnouncementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SendAnnouncement entities.
func (c *SendAnnouncementClient) CreateBulk(builders ...*SendAnnouncementCreate) *SendAnnouncementCreateBulk {
	return &SendAnnouncementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SendAnnouncement.
func (c *SendAnnouncementClient) Update() *SendAnnouncementUpdate {
	mutation := newSendAnnouncementMutation(c.config, OpUpdate)
	return &SendAnnouncementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SendAnnouncementClient) UpdateOne(sa *SendAnnouncement) *SendAnnouncementUpdateOne {
	mutation := newSendAnnouncementMutation(c.config, OpUpdateOne, withSendAnnouncement(sa))
	return &SendAnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SendAnnouncementClient) UpdateOneID(id uuid.UUID) *SendAnnouncementUpdateOne {
	mutation := newSendAnnouncementMutation(c.config, OpUpdateOne, withSendAnnouncementID(id))
	return &SendAnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SendAnnouncement.
func (c *SendAnnouncementClient) Delete() *SendAnnouncementDelete {
	mutation := newSendAnnouncementMutation(c.config, OpDelete)
	return &SendAnnouncementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SendAnnouncementClient) DeleteOne(sa *SendAnnouncement) *SendAnnouncementDeleteOne {
	return c.DeleteOneID(sa.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SendAnnouncementClient) DeleteOneID(id uuid.UUID) *SendAnnouncementDeleteOne {
	builder := c.Delete().Where(sendannouncement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SendAnnouncementDeleteOne{builder}
}

// Query returns a query builder for SendAnnouncement.
func (c *SendAnnouncementClient) Query() *SendAnnouncementQuery {
	return &SendAnnouncementQuery{
		config: c.config,
	}
}

// Get returns a SendAnnouncement entity by its id.
func (c *SendAnnouncementClient) Get(ctx context.Context, id uuid.UUID) (*SendAnnouncement, error) {
	return c.Query().Where(sendannouncement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SendAnnouncementClient) GetX(ctx context.Context, id uuid.UUID) *SendAnnouncement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SendAnnouncementClient) Hooks() []Hook {
	hooks := c.hooks.SendAnnouncement
	return append(hooks[:len(hooks):len(hooks)], sendannouncement.Hooks[:]...)
}

// TxNotifStateClient is a client for the TxNotifState schema.
type TxNotifStateClient struct {
	config
}

// NewTxNotifStateClient returns a client for the TxNotifState from the given config.
func NewTxNotifStateClient(c config) *TxNotifStateClient {
	return &TxNotifStateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `txnotifstate.Hooks(f(g(h())))`.
func (c *TxNotifStateClient) Use(hooks ...Hook) {
	c.hooks.TxNotifState = append(c.hooks.TxNotifState, hooks...)
}

// Create returns a builder for creating a TxNotifState entity.
func (c *TxNotifStateClient) Create() *TxNotifStateCreate {
	mutation := newTxNotifStateMutation(c.config, OpCreate)
	return &TxNotifStateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TxNotifState entities.
func (c *TxNotifStateClient) CreateBulk(builders ...*TxNotifStateCreate) *TxNotifStateCreateBulk {
	return &TxNotifStateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TxNotifState.
func (c *TxNotifStateClient) Update() *TxNotifStateUpdate {
	mutation := newTxNotifStateMutation(c.config, OpUpdate)
	return &TxNotifStateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TxNotifStateClient) UpdateOne(tns *TxNotifState) *TxNotifStateUpdateOne {
	mutation := newTxNotifStateMutation(c.config, OpUpdateOne, withTxNotifState(tns))
	return &TxNotifStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TxNotifStateClient) UpdateOneID(id uuid.UUID) *TxNotifStateUpdateOne {
	mutation := newTxNotifStateMutation(c.config, OpUpdateOne, withTxNotifStateID(id))
	return &TxNotifStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TxNotifState.
func (c *TxNotifStateClient) Delete() *TxNotifStateDelete {
	mutation := newTxNotifStateMutation(c.config, OpDelete)
	return &TxNotifStateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TxNotifStateClient) DeleteOne(tns *TxNotifState) *TxNotifStateDeleteOne {
	return c.DeleteOneID(tns.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TxNotifStateClient) DeleteOneID(id uuid.UUID) *TxNotifStateDeleteOne {
	builder := c.Delete().Where(txnotifstate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TxNotifStateDeleteOne{builder}
}

// Query returns a query builder for TxNotifState.
func (c *TxNotifStateClient) Query() *TxNotifStateQuery {
	return &TxNotifStateQuery{
		config: c.config,
	}
}

// Get returns a TxNotifState entity by its id.
func (c *TxNotifStateClient) Get(ctx context.Context, id uuid.UUID) (*TxNotifState, error) {
	return c.Query().Where(txnotifstate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TxNotifStateClient) GetX(ctx context.Context, id uuid.UUID) *TxNotifState {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TxNotifStateClient) Hooks() []Hook {
	hooks := c.hooks.TxNotifState
	return append(hooks[:len(hooks):len(hooks)], txnotifstate.Hooks[:]...)
}

// UserAnnouncementClient is a client for the UserAnnouncement schema.
type UserAnnouncementClient struct {
	config
}

// NewUserAnnouncementClient returns a client for the UserAnnouncement from the given config.
func NewUserAnnouncementClient(c config) *UserAnnouncementClient {
	return &UserAnnouncementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userannouncement.Hooks(f(g(h())))`.
func (c *UserAnnouncementClient) Use(hooks ...Hook) {
	c.hooks.UserAnnouncement = append(c.hooks.UserAnnouncement, hooks...)
}

// Create returns a builder for creating a UserAnnouncement entity.
func (c *UserAnnouncementClient) Create() *UserAnnouncementCreate {
	mutation := newUserAnnouncementMutation(c.config, OpCreate)
	return &UserAnnouncementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserAnnouncement entities.
func (c *UserAnnouncementClient) CreateBulk(builders ...*UserAnnouncementCreate) *UserAnnouncementCreateBulk {
	return &UserAnnouncementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserAnnouncement.
func (c *UserAnnouncementClient) Update() *UserAnnouncementUpdate {
	mutation := newUserAnnouncementMutation(c.config, OpUpdate)
	return &UserAnnouncementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserAnnouncementClient) UpdateOne(ua *UserAnnouncement) *UserAnnouncementUpdateOne {
	mutation := newUserAnnouncementMutation(c.config, OpUpdateOne, withUserAnnouncement(ua))
	return &UserAnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserAnnouncementClient) UpdateOneID(id uuid.UUID) *UserAnnouncementUpdateOne {
	mutation := newUserAnnouncementMutation(c.config, OpUpdateOne, withUserAnnouncementID(id))
	return &UserAnnouncementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserAnnouncement.
func (c *UserAnnouncementClient) Delete() *UserAnnouncementDelete {
	mutation := newUserAnnouncementMutation(c.config, OpDelete)
	return &UserAnnouncementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserAnnouncementClient) DeleteOne(ua *UserAnnouncement) *UserAnnouncementDeleteOne {
	return c.DeleteOneID(ua.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserAnnouncementClient) DeleteOneID(id uuid.UUID) *UserAnnouncementDeleteOne {
	builder := c.Delete().Where(userannouncement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserAnnouncementDeleteOne{builder}
}

// Query returns a query builder for UserAnnouncement.
func (c *UserAnnouncementClient) Query() *UserAnnouncementQuery {
	return &UserAnnouncementQuery{
		config: c.config,
	}
}

// Get returns a UserAnnouncement entity by its id.
func (c *UserAnnouncementClient) Get(ctx context.Context, id uuid.UUID) (*UserAnnouncement, error) {
	return c.Query().Where(userannouncement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserAnnouncementClient) GetX(ctx context.Context, id uuid.UUID) *UserAnnouncement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserAnnouncementClient) Hooks() []Hook {
	hooks := c.hooks.UserAnnouncement
	return append(hooks[:len(hooks):len(hooks)], userannouncement.Hooks[:]...)
}

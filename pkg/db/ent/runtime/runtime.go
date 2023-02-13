// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/announcement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/notif"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/notifchannel"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/readannouncement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/schema"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/sendannouncement"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/txnotifstate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/userannouncement"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	announcementMixin := schema.Announcement{}.Mixin()
	announcement.Policy = privacy.NewPolicies(announcementMixin[0], schema.Announcement{})
	announcement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := announcement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	announcementMixinFields0 := announcementMixin[0].Fields()
	_ = announcementMixinFields0
	announcementFields := schema.Announcement{}.Fields()
	_ = announcementFields
	// announcementDescCreatedAt is the schema descriptor for created_at field.
	announcementDescCreatedAt := announcementMixinFields0[0].Descriptor()
	// announcement.DefaultCreatedAt holds the default value on creation for the created_at field.
	announcement.DefaultCreatedAt = announcementDescCreatedAt.Default.(func() uint32)
	// announcementDescUpdatedAt is the schema descriptor for updated_at field.
	announcementDescUpdatedAt := announcementMixinFields0[1].Descriptor()
	// announcement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	announcement.DefaultUpdatedAt = announcementDescUpdatedAt.Default.(func() uint32)
	// announcement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	announcement.UpdateDefaultUpdatedAt = announcementDescUpdatedAt.UpdateDefault.(func() uint32)
	// announcementDescDeletedAt is the schema descriptor for deleted_at field.
	announcementDescDeletedAt := announcementMixinFields0[2].Descriptor()
	// announcement.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	announcement.DefaultDeletedAt = announcementDescDeletedAt.Default.(func() uint32)
	// announcementDescAppID is the schema descriptor for app_id field.
	announcementDescAppID := announcementFields[1].Descriptor()
	// announcement.DefaultAppID holds the default value on creation for the app_id field.
	announcement.DefaultAppID = announcementDescAppID.Default.(func() uuid.UUID)
	// announcementDescLangID is the schema descriptor for lang_id field.
	announcementDescLangID := announcementFields[2].Descriptor()
	// announcement.DefaultLangID holds the default value on creation for the lang_id field.
	announcement.DefaultLangID = announcementDescLangID.Default.(func() uuid.UUID)
	// announcementDescTitle is the schema descriptor for title field.
	announcementDescTitle := announcementFields[3].Descriptor()
	// announcement.DefaultTitle holds the default value on creation for the title field.
	announcement.DefaultTitle = announcementDescTitle.Default.(string)
	// announcementDescContent is the schema descriptor for content field.
	announcementDescContent := announcementFields[4].Descriptor()
	// announcement.DefaultContent holds the default value on creation for the content field.
	announcement.DefaultContent = announcementDescContent.Default.(string)
	// announcementDescChannels is the schema descriptor for channels field.
	announcementDescChannels := announcementFields[5].Descriptor()
	// announcement.DefaultChannels holds the default value on creation for the channels field.
	announcement.DefaultChannels = announcementDescChannels.Default.([]string)
	// announcementDescEndAt is the schema descriptor for end_at field.
	announcementDescEndAt := announcementFields[6].Descriptor()
	// announcement.DefaultEndAt holds the default value on creation for the end_at field.
	announcement.DefaultEndAt = announcementDescEndAt.Default.(uint32)
	// announcementDescType is the schema descriptor for type field.
	announcementDescType := announcementFields[7].Descriptor()
	// announcement.DefaultType holds the default value on creation for the type field.
	announcement.DefaultType = announcementDescType.Default.(string)
	// announcementDescID is the schema descriptor for id field.
	announcementDescID := announcementFields[0].Descriptor()
	// announcement.DefaultID holds the default value on creation for the id field.
	announcement.DefaultID = announcementDescID.Default.(func() uuid.UUID)
	notifMixin := schema.Notif{}.Mixin()
	notif.Policy = privacy.NewPolicies(notifMixin[0], schema.Notif{})
	notif.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := notif.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	notifMixinFields0 := notifMixin[0].Fields()
	_ = notifMixinFields0
	notifFields := schema.Notif{}.Fields()
	_ = notifFields
	// notifDescCreatedAt is the schema descriptor for created_at field.
	notifDescCreatedAt := notifMixinFields0[0].Descriptor()
	// notif.DefaultCreatedAt holds the default value on creation for the created_at field.
	notif.DefaultCreatedAt = notifDescCreatedAt.Default.(func() uint32)
	// notifDescUpdatedAt is the schema descriptor for updated_at field.
	notifDescUpdatedAt := notifMixinFields0[1].Descriptor()
	// notif.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notif.DefaultUpdatedAt = notifDescUpdatedAt.Default.(func() uint32)
	// notif.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notif.UpdateDefaultUpdatedAt = notifDescUpdatedAt.UpdateDefault.(func() uint32)
	// notifDescDeletedAt is the schema descriptor for deleted_at field.
	notifDescDeletedAt := notifMixinFields0[2].Descriptor()
	// notif.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	notif.DefaultDeletedAt = notifDescDeletedAt.Default.(func() uint32)
	// notifDescAppID is the schema descriptor for app_id field.
	notifDescAppID := notifFields[1].Descriptor()
	// notif.DefaultAppID holds the default value on creation for the app_id field.
	notif.DefaultAppID = notifDescAppID.Default.(func() uuid.UUID)
	// notifDescUserID is the schema descriptor for user_id field.
	notifDescUserID := notifFields[2].Descriptor()
	// notif.DefaultUserID holds the default value on creation for the user_id field.
	notif.DefaultUserID = notifDescUserID.Default.(func() uuid.UUID)
	// notifDescAlreadyRead is the schema descriptor for already_read field.
	notifDescAlreadyRead := notifFields[3].Descriptor()
	// notif.DefaultAlreadyRead holds the default value on creation for the already_read field.
	notif.DefaultAlreadyRead = notifDescAlreadyRead.Default.(bool)
	// notifDescLangID is the schema descriptor for lang_id field.
	notifDescLangID := notifFields[4].Descriptor()
	// notif.DefaultLangID holds the default value on creation for the lang_id field.
	notif.DefaultLangID = notifDescLangID.Default.(func() uuid.UUID)
	// notifDescEventType is the schema descriptor for event_type field.
	notifDescEventType := notifFields[5].Descriptor()
	// notif.DefaultEventType holds the default value on creation for the event_type field.
	notif.DefaultEventType = notifDescEventType.Default.(string)
	// notifDescUseTemplate is the schema descriptor for use_template field.
	notifDescUseTemplate := notifFields[6].Descriptor()
	// notif.DefaultUseTemplate holds the default value on creation for the use_template field.
	notif.DefaultUseTemplate = notifDescUseTemplate.Default.(bool)
	// notifDescTitle is the schema descriptor for title field.
	notifDescTitle := notifFields[7].Descriptor()
	// notif.DefaultTitle holds the default value on creation for the title field.
	notif.DefaultTitle = notifDescTitle.Default.(string)
	// notifDescContent is the schema descriptor for content field.
	notifDescContent := notifFields[8].Descriptor()
	// notif.DefaultContent holds the default value on creation for the content field.
	notif.DefaultContent = notifDescContent.Default.(string)
	// notifDescChannels is the schema descriptor for channels field.
	notifDescChannels := notifFields[9].Descriptor()
	// notif.DefaultChannels holds the default value on creation for the channels field.
	notif.DefaultChannels = notifDescChannels.Default.([]string)
	// notifDescEmailSend is the schema descriptor for email_send field.
	notifDescEmailSend := notifFields[10].Descriptor()
	// notif.DefaultEmailSend holds the default value on creation for the email_send field.
	notif.DefaultEmailSend = notifDescEmailSend.Default.(bool)
	// notifDescExtra is the schema descriptor for extra field.
	notifDescExtra := notifFields[11].Descriptor()
	// notif.DefaultExtra holds the default value on creation for the extra field.
	notif.DefaultExtra = notifDescExtra.Default.(string)
	// notifDescID is the schema descriptor for id field.
	notifDescID := notifFields[0].Descriptor()
	// notif.DefaultID holds the default value on creation for the id field.
	notif.DefaultID = notifDescID.Default.(func() uuid.UUID)
	notifchannelMixin := schema.NotifChannel{}.Mixin()
	notifchannel.Policy = privacy.NewPolicies(notifchannelMixin[0], schema.NotifChannel{})
	notifchannel.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := notifchannel.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	notifchannelMixinFields0 := notifchannelMixin[0].Fields()
	_ = notifchannelMixinFields0
	notifchannelFields := schema.NotifChannel{}.Fields()
	_ = notifchannelFields
	// notifchannelDescCreatedAt is the schema descriptor for created_at field.
	notifchannelDescCreatedAt := notifchannelMixinFields0[0].Descriptor()
	// notifchannel.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifchannel.DefaultCreatedAt = notifchannelDescCreatedAt.Default.(func() uint32)
	// notifchannelDescUpdatedAt is the schema descriptor for updated_at field.
	notifchannelDescUpdatedAt := notifchannelMixinFields0[1].Descriptor()
	// notifchannel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifchannel.DefaultUpdatedAt = notifchannelDescUpdatedAt.Default.(func() uint32)
	// notifchannel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifchannel.UpdateDefaultUpdatedAt = notifchannelDescUpdatedAt.UpdateDefault.(func() uint32)
	// notifchannelDescDeletedAt is the schema descriptor for deleted_at field.
	notifchannelDescDeletedAt := notifchannelMixinFields0[2].Descriptor()
	// notifchannel.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	notifchannel.DefaultDeletedAt = notifchannelDescDeletedAt.Default.(func() uint32)
	// notifchannelDescAppID is the schema descriptor for app_id field.
	notifchannelDescAppID := notifchannelFields[1].Descriptor()
	// notifchannel.DefaultAppID holds the default value on creation for the app_id field.
	notifchannel.DefaultAppID = notifchannelDescAppID.Default.(func() uuid.UUID)
	// notifchannelDescEventType is the schema descriptor for event_type field.
	notifchannelDescEventType := notifchannelFields[2].Descriptor()
	// notifchannel.DefaultEventType holds the default value on creation for the event_type field.
	notifchannel.DefaultEventType = notifchannelDescEventType.Default.(string)
	// notifchannelDescChannel is the schema descriptor for channel field.
	notifchannelDescChannel := notifchannelFields[3].Descriptor()
	// notifchannel.DefaultChannel holds the default value on creation for the channel field.
	notifchannel.DefaultChannel = notifchannelDescChannel.Default.(string)
	// notifchannelDescID is the schema descriptor for id field.
	notifchannelDescID := notifchannelFields[0].Descriptor()
	// notifchannel.DefaultID holds the default value on creation for the id field.
	notifchannel.DefaultID = notifchannelDescID.Default.(func() uuid.UUID)
	readannouncementMixin := schema.ReadAnnouncement{}.Mixin()
	readannouncement.Policy = privacy.NewPolicies(readannouncementMixin[0], schema.ReadAnnouncement{})
	readannouncement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := readannouncement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	readannouncementMixinFields0 := readannouncementMixin[0].Fields()
	_ = readannouncementMixinFields0
	readannouncementFields := schema.ReadAnnouncement{}.Fields()
	_ = readannouncementFields
	// readannouncementDescCreatedAt is the schema descriptor for created_at field.
	readannouncementDescCreatedAt := readannouncementMixinFields0[0].Descriptor()
	// readannouncement.DefaultCreatedAt holds the default value on creation for the created_at field.
	readannouncement.DefaultCreatedAt = readannouncementDescCreatedAt.Default.(func() uint32)
	// readannouncementDescUpdatedAt is the schema descriptor for updated_at field.
	readannouncementDescUpdatedAt := readannouncementMixinFields0[1].Descriptor()
	// readannouncement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	readannouncement.DefaultUpdatedAt = readannouncementDescUpdatedAt.Default.(func() uint32)
	// readannouncement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	readannouncement.UpdateDefaultUpdatedAt = readannouncementDescUpdatedAt.UpdateDefault.(func() uint32)
	// readannouncementDescDeletedAt is the schema descriptor for deleted_at field.
	readannouncementDescDeletedAt := readannouncementMixinFields0[2].Descriptor()
	// readannouncement.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	readannouncement.DefaultDeletedAt = readannouncementDescDeletedAt.Default.(func() uint32)
	// readannouncementDescAppID is the schema descriptor for app_id field.
	readannouncementDescAppID := readannouncementFields[1].Descriptor()
	// readannouncement.DefaultAppID holds the default value on creation for the app_id field.
	readannouncement.DefaultAppID = readannouncementDescAppID.Default.(func() uuid.UUID)
	// readannouncementDescUserID is the schema descriptor for user_id field.
	readannouncementDescUserID := readannouncementFields[2].Descriptor()
	// readannouncement.DefaultUserID holds the default value on creation for the user_id field.
	readannouncement.DefaultUserID = readannouncementDescUserID.Default.(func() uuid.UUID)
	// readannouncementDescAnnouncementID is the schema descriptor for announcement_id field.
	readannouncementDescAnnouncementID := readannouncementFields[3].Descriptor()
	// readannouncement.DefaultAnnouncementID holds the default value on creation for the announcement_id field.
	readannouncement.DefaultAnnouncementID = readannouncementDescAnnouncementID.Default.(func() uuid.UUID)
	// readannouncementDescID is the schema descriptor for id field.
	readannouncementDescID := readannouncementFields[0].Descriptor()
	// readannouncement.DefaultID holds the default value on creation for the id field.
	readannouncement.DefaultID = readannouncementDescID.Default.(func() uuid.UUID)
	sendannouncementMixin := schema.SendAnnouncement{}.Mixin()
	sendannouncement.Policy = privacy.NewPolicies(sendannouncementMixin[0], schema.SendAnnouncement{})
	sendannouncement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := sendannouncement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	sendannouncementMixinFields0 := sendannouncementMixin[0].Fields()
	_ = sendannouncementMixinFields0
	sendannouncementFields := schema.SendAnnouncement{}.Fields()
	_ = sendannouncementFields
	// sendannouncementDescCreatedAt is the schema descriptor for created_at field.
	sendannouncementDescCreatedAt := sendannouncementMixinFields0[0].Descriptor()
	// sendannouncement.DefaultCreatedAt holds the default value on creation for the created_at field.
	sendannouncement.DefaultCreatedAt = sendannouncementDescCreatedAt.Default.(func() uint32)
	// sendannouncementDescUpdatedAt is the schema descriptor for updated_at field.
	sendannouncementDescUpdatedAt := sendannouncementMixinFields0[1].Descriptor()
	// sendannouncement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sendannouncement.DefaultUpdatedAt = sendannouncementDescUpdatedAt.Default.(func() uint32)
	// sendannouncement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sendannouncement.UpdateDefaultUpdatedAt = sendannouncementDescUpdatedAt.UpdateDefault.(func() uint32)
	// sendannouncementDescDeletedAt is the schema descriptor for deleted_at field.
	sendannouncementDescDeletedAt := sendannouncementMixinFields0[2].Descriptor()
	// sendannouncement.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	sendannouncement.DefaultDeletedAt = sendannouncementDescDeletedAt.Default.(func() uint32)
	// sendannouncementDescAppID is the schema descriptor for app_id field.
	sendannouncementDescAppID := sendannouncementFields[1].Descriptor()
	// sendannouncement.DefaultAppID holds the default value on creation for the app_id field.
	sendannouncement.DefaultAppID = sendannouncementDescAppID.Default.(func() uuid.UUID)
	// sendannouncementDescUserID is the schema descriptor for user_id field.
	sendannouncementDescUserID := sendannouncementFields[2].Descriptor()
	// sendannouncement.DefaultUserID holds the default value on creation for the user_id field.
	sendannouncement.DefaultUserID = sendannouncementDescUserID.Default.(func() uuid.UUID)
	// sendannouncementDescAnnouncementID is the schema descriptor for announcement_id field.
	sendannouncementDescAnnouncementID := sendannouncementFields[3].Descriptor()
	// sendannouncement.DefaultAnnouncementID holds the default value on creation for the announcement_id field.
	sendannouncement.DefaultAnnouncementID = sendannouncementDescAnnouncementID.Default.(func() uuid.UUID)
	// sendannouncementDescChannel is the schema descriptor for channel field.
	sendannouncementDescChannel := sendannouncementFields[4].Descriptor()
	// sendannouncement.DefaultChannel holds the default value on creation for the channel field.
	sendannouncement.DefaultChannel = sendannouncementDescChannel.Default.(string)
	// sendannouncementDescID is the schema descriptor for id field.
	sendannouncementDescID := sendannouncementFields[0].Descriptor()
	// sendannouncement.DefaultID holds the default value on creation for the id field.
	sendannouncement.DefaultID = sendannouncementDescID.Default.(func() uuid.UUID)
	txnotifstateMixin := schema.TxNotifState{}.Mixin()
	txnotifstate.Policy = privacy.NewPolicies(txnotifstateMixin[0], schema.TxNotifState{})
	txnotifstate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := txnotifstate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	txnotifstateMixinFields0 := txnotifstateMixin[0].Fields()
	_ = txnotifstateMixinFields0
	txnotifstateFields := schema.TxNotifState{}.Fields()
	_ = txnotifstateFields
	// txnotifstateDescCreatedAt is the schema descriptor for created_at field.
	txnotifstateDescCreatedAt := txnotifstateMixinFields0[0].Descriptor()
	// txnotifstate.DefaultCreatedAt holds the default value on creation for the created_at field.
	txnotifstate.DefaultCreatedAt = txnotifstateDescCreatedAt.Default.(func() uint32)
	// txnotifstateDescUpdatedAt is the schema descriptor for updated_at field.
	txnotifstateDescUpdatedAt := txnotifstateMixinFields0[1].Descriptor()
	// txnotifstate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	txnotifstate.DefaultUpdatedAt = txnotifstateDescUpdatedAt.Default.(func() uint32)
	// txnotifstate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	txnotifstate.UpdateDefaultUpdatedAt = txnotifstateDescUpdatedAt.UpdateDefault.(func() uint32)
	// txnotifstateDescDeletedAt is the schema descriptor for deleted_at field.
	txnotifstateDescDeletedAt := txnotifstateMixinFields0[2].Descriptor()
	// txnotifstate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	txnotifstate.DefaultDeletedAt = txnotifstateDescDeletedAt.Default.(func() uint32)
	// txnotifstateDescTxID is the schema descriptor for tx_id field.
	txnotifstateDescTxID := txnotifstateFields[1].Descriptor()
	// txnotifstate.DefaultTxID holds the default value on creation for the tx_id field.
	txnotifstate.DefaultTxID = txnotifstateDescTxID.Default.(func() uuid.UUID)
	// txnotifstateDescNotifState is the schema descriptor for notif_state field.
	txnotifstateDescNotifState := txnotifstateFields[2].Descriptor()
	// txnotifstate.DefaultNotifState holds the default value on creation for the notif_state field.
	txnotifstate.DefaultNotifState = txnotifstateDescNotifState.Default.(string)
	// txnotifstateDescNotifType is the schema descriptor for notif_type field.
	txnotifstateDescNotifType := txnotifstateFields[3].Descriptor()
	// txnotifstate.DefaultNotifType holds the default value on creation for the notif_type field.
	txnotifstate.DefaultNotifType = txnotifstateDescNotifType.Default.(string)
	// txnotifstateDescID is the schema descriptor for id field.
	txnotifstateDescID := txnotifstateFields[0].Descriptor()
	// txnotifstate.DefaultID holds the default value on creation for the id field.
	txnotifstate.DefaultID = txnotifstateDescID.Default.(func() uuid.UUID)
	userannouncementMixin := schema.UserAnnouncement{}.Mixin()
	userannouncement.Policy = privacy.NewPolicies(userannouncementMixin[0], schema.UserAnnouncement{})
	userannouncement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := userannouncement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userannouncementMixinFields0 := userannouncementMixin[0].Fields()
	_ = userannouncementMixinFields0
	userannouncementFields := schema.UserAnnouncement{}.Fields()
	_ = userannouncementFields
	// userannouncementDescCreatedAt is the schema descriptor for created_at field.
	userannouncementDescCreatedAt := userannouncementMixinFields0[0].Descriptor()
	// userannouncement.DefaultCreatedAt holds the default value on creation for the created_at field.
	userannouncement.DefaultCreatedAt = userannouncementDescCreatedAt.Default.(func() uint32)
	// userannouncementDescUpdatedAt is the schema descriptor for updated_at field.
	userannouncementDescUpdatedAt := userannouncementMixinFields0[1].Descriptor()
	// userannouncement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userannouncement.DefaultUpdatedAt = userannouncementDescUpdatedAt.Default.(func() uint32)
	// userannouncement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userannouncement.UpdateDefaultUpdatedAt = userannouncementDescUpdatedAt.UpdateDefault.(func() uint32)
	// userannouncementDescDeletedAt is the schema descriptor for deleted_at field.
	userannouncementDescDeletedAt := userannouncementMixinFields0[2].Descriptor()
	// userannouncement.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	userannouncement.DefaultDeletedAt = userannouncementDescDeletedAt.Default.(func() uint32)
	// userannouncementDescAppID is the schema descriptor for app_id field.
	userannouncementDescAppID := userannouncementFields[1].Descriptor()
	// userannouncement.DefaultAppID holds the default value on creation for the app_id field.
	userannouncement.DefaultAppID = userannouncementDescAppID.Default.(func() uuid.UUID)
	// userannouncementDescUserID is the schema descriptor for user_id field.
	userannouncementDescUserID := userannouncementFields[2].Descriptor()
	// userannouncement.DefaultUserID holds the default value on creation for the user_id field.
	userannouncement.DefaultUserID = userannouncementDescUserID.Default.(func() uuid.UUID)
	// userannouncementDescAnnouncementID is the schema descriptor for announcement_id field.
	userannouncementDescAnnouncementID := userannouncementFields[3].Descriptor()
	// userannouncement.DefaultAnnouncementID holds the default value on creation for the announcement_id field.
	userannouncement.DefaultAnnouncementID = userannouncementDescAnnouncementID.Default.(func() uuid.UUID)
	// userannouncementDescID is the schema descriptor for id field.
	userannouncementDescID := userannouncementFields[0].Descriptor()
	// userannouncement.DefaultID holds the default value on creation for the id field.
	userannouncement.DefaultID = userannouncementDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)

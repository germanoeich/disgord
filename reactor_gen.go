package disgord

// Code generated - This file has been automatically generated by generate/events/main.go - DO NOT EDIT.
// Warning: This file is overwritten at "go generate", instead adapt reactor.go and event/events.go and run `go generate`

//////////////////////////////////////////////////////
//
// Helper funcs
//
//////////////////////////////////////////////////////

func defineResource(evt string) (resource evtResource) {
	switch evt {

	case EvtChannelCreate:
		resource = &ChannelCreate{}
	case EvtChannelDelete:
		resource = &ChannelDelete{}
	case EvtChannelPinsUpdate:
		resource = &ChannelPinsUpdate{}
	case EvtChannelUpdate:
		resource = &ChannelUpdate{}
	case EvtGuildBanAdd:
		resource = &GuildBanAdd{}
	case EvtGuildBanRemove:
		resource = &GuildBanRemove{}
	case EvtGuildCreate:
		resource = &GuildCreate{}
	case EvtGuildDelete:
		resource = &GuildDelete{}
	case EvtGuildEmojisUpdate:
		resource = &GuildEmojisUpdate{}
	case EvtGuildIntegrationsUpdate:
		resource = &GuildIntegrationsUpdate{}
	case EvtGuildMemberAdd:
		resource = &GuildMemberAdd{}
	case EvtGuildMemberRemove:
		resource = &GuildMemberRemove{}
	case EvtGuildMemberUpdate:
		resource = &GuildMemberUpdate{}
	case EvtGuildMembersChunk:
		resource = &GuildMembersChunk{}
	case EvtGuildRoleCreate:
		resource = &GuildRoleCreate{}
	case EvtGuildRoleDelete:
		resource = &GuildRoleDelete{}
	case EvtGuildRoleUpdate:
		resource = &GuildRoleUpdate{}
	case EvtGuildUpdate:
		resource = &GuildUpdate{}
	case EvtInviteDelete:
		resource = &InviteDelete{}
	case EvtMessageCreate:
		resource = &MessageCreate{}
	case EvtMessageDelete:
		resource = &MessageDelete{}
	case EvtMessageDeleteBulk:
		resource = &MessageDeleteBulk{}
	case EvtMessageReactionAdd:
		resource = &MessageReactionAdd{}
	case EvtMessageReactionRemove:
		resource = &MessageReactionRemove{}
	case EvtMessageReactionRemoveAll:
		resource = &MessageReactionRemoveAll{}
	case EvtMessageUpdate:
		resource = &MessageUpdate{}
	case EvtPresenceUpdate:
		resource = &PresenceUpdate{}
	case EvtReady:
		resource = &Ready{}
	case EvtResumed:
		resource = &Resumed{}
	case EvtTypingStart:
		resource = &TypingStart{}
	case EvtUserUpdate:
		resource = &UserUpdate{}
	case EvtVoiceServerUpdate:
		resource = &VoiceServerUpdate{}
	case EvtVoiceStateUpdate:
		resource = &VoiceStateUpdate{}
	case EvtWebhooksUpdate:
		resource = &WebhooksUpdate{}
	}

	return resource
}

func isHandler(h Handler) (ok bool) {
	switch h.(type) {
	case SimpleHandler:
		ok = true
	case SimplestHandler:
		ok = true
	case chan interface{}:
		ok = true
	case ChannelCreateHandler:
		ok = true
	case chan *ChannelCreate:
		ok = true
	case ChannelDeleteHandler:
		ok = true
	case chan *ChannelDelete:
		ok = true
	case ChannelPinsUpdateHandler:
		ok = true
	case chan *ChannelPinsUpdate:
		ok = true
	case ChannelUpdateHandler:
		ok = true
	case chan *ChannelUpdate:
		ok = true
	case GuildBanAddHandler:
		ok = true
	case chan *GuildBanAdd:
		ok = true
	case GuildBanRemoveHandler:
		ok = true
	case chan *GuildBanRemove:
		ok = true
	case GuildCreateHandler:
		ok = true
	case chan *GuildCreate:
		ok = true
	case GuildDeleteHandler:
		ok = true
	case chan *GuildDelete:
		ok = true
	case GuildEmojisUpdateHandler:
		ok = true
	case chan *GuildEmojisUpdate:
		ok = true
	case GuildIntegrationsUpdateHandler:
		ok = true
	case chan *GuildIntegrationsUpdate:
		ok = true
	case GuildMemberAddHandler:
		ok = true
	case chan *GuildMemberAdd:
		ok = true
	case GuildMemberRemoveHandler:
		ok = true
	case chan *GuildMemberRemove:
		ok = true
	case GuildMemberUpdateHandler:
		ok = true
	case chan *GuildMemberUpdate:
		ok = true
	case GuildMembersChunkHandler:
		ok = true
	case chan *GuildMembersChunk:
		ok = true
	case GuildRoleCreateHandler:
		ok = true
	case chan *GuildRoleCreate:
		ok = true
	case GuildRoleDeleteHandler:
		ok = true
	case chan *GuildRoleDelete:
		ok = true
	case GuildRoleUpdateHandler:
		ok = true
	case chan *GuildRoleUpdate:
		ok = true
	case GuildUpdateHandler:
		ok = true
	case chan *GuildUpdate:
		ok = true
	case InviteDeleteHandler:
		ok = true
	case chan *InviteDelete:
		ok = true
	case MessageCreateHandler:
		ok = true
	case chan *MessageCreate:
		ok = true
	case MessageDeleteHandler:
		ok = true
	case chan *MessageDelete:
		ok = true
	case MessageDeleteBulkHandler:
		ok = true
	case chan *MessageDeleteBulk:
		ok = true
	case MessageReactionAddHandler:
		ok = true
	case chan *MessageReactionAdd:
		ok = true
	case MessageReactionRemoveHandler:
		ok = true
	case chan *MessageReactionRemove:
		ok = true
	case MessageReactionRemoveAllHandler:
		ok = true
	case chan *MessageReactionRemoveAll:
		ok = true
	case MessageUpdateHandler:
		ok = true
	case chan *MessageUpdate:
		ok = true
	case PresenceUpdateHandler:
		ok = true
	case chan *PresenceUpdate:
		ok = true
	case ReadyHandler:
		ok = true
	case chan *Ready:
		ok = true
	case ResumedHandler:
		ok = true
	case chan *Resumed:
		ok = true
	case TypingStartHandler:
		ok = true
	case chan *TypingStart:
		ok = true
	case UserUpdateHandler:
		ok = true
	case chan *UserUpdate:
		ok = true
	case VoiceServerUpdateHandler:
		ok = true
	case chan *VoiceServerUpdate:
		ok = true
	case VoiceStateUpdateHandler:
		ok = true
	case chan *VoiceStateUpdate:
		ok = true
	case WebhooksUpdateHandler:
		ok = true
	case chan *WebhooksUpdate:
		ok = true
	}
	return ok
}

func closeChannel(channel interface{}) {
	switch t := channel.(type) {
	case chan interface{}:
		close(t)
	case chan *ChannelCreate:
		close(t)
	case chan *ChannelDelete:
		close(t)
	case chan *ChannelPinsUpdate:
		close(t)
	case chan *ChannelUpdate:
		close(t)
	case chan *GuildBanAdd:
		close(t)
	case chan *GuildBanRemove:
		close(t)
	case chan *GuildCreate:
		close(t)
	case chan *GuildDelete:
		close(t)
	case chan *GuildEmojisUpdate:
		close(t)
	case chan *GuildIntegrationsUpdate:
		close(t)
	case chan *GuildMemberAdd:
		close(t)
	case chan *GuildMemberRemove:
		close(t)
	case chan *GuildMemberUpdate:
		close(t)
	case chan *GuildMembersChunk:
		close(t)
	case chan *GuildRoleCreate:
		close(t)
	case chan *GuildRoleDelete:
		close(t)
	case chan *GuildRoleUpdate:
		close(t)
	case chan *GuildUpdate:
		close(t)
	case chan *InviteDelete:
		close(t)
	case chan *MessageCreate:
		close(t)
	case chan *MessageDelete:
		close(t)
	case chan *MessageDeleteBulk:
		close(t)
	case chan *MessageReactionAdd:
		close(t)
	case chan *MessageReactionRemove:
		close(t)
	case chan *MessageReactionRemoveAll:
		close(t)
	case chan *MessageUpdate:
		close(t)
	case chan *PresenceUpdate:
		close(t)
	case chan *Ready:
		close(t)
	case chan *Resumed:
		close(t)
	case chan *TypingStart:
		close(t)
	case chan *UserUpdate:
		close(t)
	case chan *VoiceServerUpdate:
		close(t)
	case chan *VoiceStateUpdate:
		close(t)
	case chan *WebhooksUpdate:
		close(t)
	}
}

//////////////////////////////////////////////////////
//
// Dispatcher: contructor + repetitive methods
//
//////////////////////////////////////////////////////

// newDispatcher construct a Dispatch object for reacting to web socket events
// from discord
func newDispatcher() *dispatcher {
	d := &dispatcher{
		handlerSpecs: make(map[string][]*handlerSpec),
		shutdown:     make(chan struct{}),
	}

	return d
}

func (d *dispatcher) trigger(h Handler, evt resource) {
	switch t := h.(type) {
	case SimpleHandler:
		t(d.session)
	case SimplestHandler:
		t()
	case chan interface{}:
		t <- evt
	case chan<- interface{}:
		t <- evt
	case ChannelCreateHandler:
		t(d.session, evt.(*ChannelCreate))
	case chan *ChannelCreate:
		t <- evt.(*ChannelCreate)
	case chan<- *ChannelCreate:
		t <- evt.(*ChannelCreate)
	case ChannelDeleteHandler:
		t(d.session, evt.(*ChannelDelete))
	case chan *ChannelDelete:
		t <- evt.(*ChannelDelete)
	case chan<- *ChannelDelete:
		t <- evt.(*ChannelDelete)
	case ChannelPinsUpdateHandler:
		t(d.session, evt.(*ChannelPinsUpdate))
	case chan *ChannelPinsUpdate:
		t <- evt.(*ChannelPinsUpdate)
	case chan<- *ChannelPinsUpdate:
		t <- evt.(*ChannelPinsUpdate)
	case ChannelUpdateHandler:
		t(d.session, evt.(*ChannelUpdate))
	case chan *ChannelUpdate:
		t <- evt.(*ChannelUpdate)
	case chan<- *ChannelUpdate:
		t <- evt.(*ChannelUpdate)
	case GuildBanAddHandler:
		t(d.session, evt.(*GuildBanAdd))
	case chan *GuildBanAdd:
		t <- evt.(*GuildBanAdd)
	case chan<- *GuildBanAdd:
		t <- evt.(*GuildBanAdd)
	case GuildBanRemoveHandler:
		t(d.session, evt.(*GuildBanRemove))
	case chan *GuildBanRemove:
		t <- evt.(*GuildBanRemove)
	case chan<- *GuildBanRemove:
		t <- evt.(*GuildBanRemove)
	case GuildCreateHandler:
		t(d.session, evt.(*GuildCreate))
	case chan *GuildCreate:
		t <- evt.(*GuildCreate)
	case chan<- *GuildCreate:
		t <- evt.(*GuildCreate)
	case GuildDeleteHandler:
		t(d.session, evt.(*GuildDelete))
	case chan *GuildDelete:
		t <- evt.(*GuildDelete)
	case chan<- *GuildDelete:
		t <- evt.(*GuildDelete)
	case GuildEmojisUpdateHandler:
		t(d.session, evt.(*GuildEmojisUpdate))
	case chan *GuildEmojisUpdate:
		t <- evt.(*GuildEmojisUpdate)
	case chan<- *GuildEmojisUpdate:
		t <- evt.(*GuildEmojisUpdate)
	case GuildIntegrationsUpdateHandler:
		t(d.session, evt.(*GuildIntegrationsUpdate))
	case chan *GuildIntegrationsUpdate:
		t <- evt.(*GuildIntegrationsUpdate)
	case chan<- *GuildIntegrationsUpdate:
		t <- evt.(*GuildIntegrationsUpdate)
	case GuildMemberAddHandler:
		t(d.session, evt.(*GuildMemberAdd))
	case chan *GuildMemberAdd:
		t <- evt.(*GuildMemberAdd)
	case chan<- *GuildMemberAdd:
		t <- evt.(*GuildMemberAdd)
	case GuildMemberRemoveHandler:
		t(d.session, evt.(*GuildMemberRemove))
	case chan *GuildMemberRemove:
		t <- evt.(*GuildMemberRemove)
	case chan<- *GuildMemberRemove:
		t <- evt.(*GuildMemberRemove)
	case GuildMemberUpdateHandler:
		t(d.session, evt.(*GuildMemberUpdate))
	case chan *GuildMemberUpdate:
		t <- evt.(*GuildMemberUpdate)
	case chan<- *GuildMemberUpdate:
		t <- evt.(*GuildMemberUpdate)
	case GuildMembersChunkHandler:
		t(d.session, evt.(*GuildMembersChunk))
	case chan *GuildMembersChunk:
		t <- evt.(*GuildMembersChunk)
	case chan<- *GuildMembersChunk:
		t <- evt.(*GuildMembersChunk)
	case GuildRoleCreateHandler:
		t(d.session, evt.(*GuildRoleCreate))
	case chan *GuildRoleCreate:
		t <- evt.(*GuildRoleCreate)
	case chan<- *GuildRoleCreate:
		t <- evt.(*GuildRoleCreate)
	case GuildRoleDeleteHandler:
		t(d.session, evt.(*GuildRoleDelete))
	case chan *GuildRoleDelete:
		t <- evt.(*GuildRoleDelete)
	case chan<- *GuildRoleDelete:
		t <- evt.(*GuildRoleDelete)
	case GuildRoleUpdateHandler:
		t(d.session, evt.(*GuildRoleUpdate))
	case chan *GuildRoleUpdate:
		t <- evt.(*GuildRoleUpdate)
	case chan<- *GuildRoleUpdate:
		t <- evt.(*GuildRoleUpdate)
	case GuildUpdateHandler:
		t(d.session, evt.(*GuildUpdate))
	case chan *GuildUpdate:
		t <- evt.(*GuildUpdate)
	case chan<- *GuildUpdate:
		t <- evt.(*GuildUpdate)
	case InviteDeleteHandler:
		t(d.session, evt.(*InviteDelete))
	case chan *InviteDelete:
		t <- evt.(*InviteDelete)
	case chan<- *InviteDelete:
		t <- evt.(*InviteDelete)
	case MessageCreateHandler:
		t(d.session, evt.(*MessageCreate))
	case chan *MessageCreate:
		t <- evt.(*MessageCreate)
	case chan<- *MessageCreate:
		t <- evt.(*MessageCreate)
	case MessageDeleteHandler:
		t(d.session, evt.(*MessageDelete))
	case chan *MessageDelete:
		t <- evt.(*MessageDelete)
	case chan<- *MessageDelete:
		t <- evt.(*MessageDelete)
	case MessageDeleteBulkHandler:
		t(d.session, evt.(*MessageDeleteBulk))
	case chan *MessageDeleteBulk:
		t <- evt.(*MessageDeleteBulk)
	case chan<- *MessageDeleteBulk:
		t <- evt.(*MessageDeleteBulk)
	case MessageReactionAddHandler:
		t(d.session, evt.(*MessageReactionAdd))
	case chan *MessageReactionAdd:
		t <- evt.(*MessageReactionAdd)
	case chan<- *MessageReactionAdd:
		t <- evt.(*MessageReactionAdd)
	case MessageReactionRemoveHandler:
		t(d.session, evt.(*MessageReactionRemove))
	case chan *MessageReactionRemove:
		t <- evt.(*MessageReactionRemove)
	case chan<- *MessageReactionRemove:
		t <- evt.(*MessageReactionRemove)
	case MessageReactionRemoveAllHandler:
		t(d.session, evt.(*MessageReactionRemoveAll))
	case chan *MessageReactionRemoveAll:
		t <- evt.(*MessageReactionRemoveAll)
	case chan<- *MessageReactionRemoveAll:
		t <- evt.(*MessageReactionRemoveAll)
	case MessageUpdateHandler:
		t(d.session, evt.(*MessageUpdate))
	case chan *MessageUpdate:
		t <- evt.(*MessageUpdate)
	case chan<- *MessageUpdate:
		t <- evt.(*MessageUpdate)
	case PresenceUpdateHandler:
		t(d.session, evt.(*PresenceUpdate))
	case chan *PresenceUpdate:
		t <- evt.(*PresenceUpdate)
	case chan<- *PresenceUpdate:
		t <- evt.(*PresenceUpdate)
	case ReadyHandler:
		t(d.session, evt.(*Ready))
	case chan *Ready:
		t <- evt.(*Ready)
	case chan<- *Ready:
		t <- evt.(*Ready)
	case ResumedHandler:
		t(d.session, evt.(*Resumed))
	case chan *Resumed:
		t <- evt.(*Resumed)
	case chan<- *Resumed:
		t <- evt.(*Resumed)
	case TypingStartHandler:
		t(d.session, evt.(*TypingStart))
	case chan *TypingStart:
		t <- evt.(*TypingStart)
	case chan<- *TypingStart:
		t <- evt.(*TypingStart)
	case UserUpdateHandler:
		t(d.session, evt.(*UserUpdate))
	case chan *UserUpdate:
		t <- evt.(*UserUpdate)
	case chan<- *UserUpdate:
		t <- evt.(*UserUpdate)
	case VoiceServerUpdateHandler:
		t(d.session, evt.(*VoiceServerUpdate))
	case chan *VoiceServerUpdate:
		t <- evt.(*VoiceServerUpdate)
	case chan<- *VoiceServerUpdate:
		t <- evt.(*VoiceServerUpdate)
	case VoiceStateUpdateHandler:
		t(d.session, evt.(*VoiceStateUpdate))
	case chan *VoiceStateUpdate:
		t <- evt.(*VoiceStateUpdate)
	case chan<- *VoiceStateUpdate:
		t <- evt.(*VoiceStateUpdate)
	case WebhooksUpdateHandler:
		t(d.session, evt.(*WebhooksUpdate))
	case chan *WebhooksUpdate:
		t <- evt.(*WebhooksUpdate)
	case chan<- *WebhooksUpdate:
		t <- evt.(*WebhooksUpdate)
	}
}

//////////////////////////////////////////////////////
//
// Handler Signatures
//
//////////////////////////////////////////////////////

// ChannelCreateHandler is triggered in ChannelCreate events
type ChannelCreateHandler = func(s Session, h *ChannelCreate)

// ChannelDeleteHandler is triggered in ChannelDelete events
type ChannelDeleteHandler = func(s Session, h *ChannelDelete)

// ChannelPinsUpdateHandler is triggered in ChannelPinsUpdate events
type ChannelPinsUpdateHandler = func(s Session, h *ChannelPinsUpdate)

// ChannelUpdateHandler is triggered in ChannelUpdate events
type ChannelUpdateHandler = func(s Session, h *ChannelUpdate)

// GuildBanAddHandler is triggered in GuildBanAdd events
type GuildBanAddHandler = func(s Session, h *GuildBanAdd)

// GuildBanRemoveHandler is triggered in GuildBanRemove events
type GuildBanRemoveHandler = func(s Session, h *GuildBanRemove)

// GuildCreateHandler is triggered in GuildCreate events
type GuildCreateHandler = func(s Session, h *GuildCreate)

// GuildDeleteHandler is triggered in GuildDelete events
type GuildDeleteHandler = func(s Session, h *GuildDelete)

// GuildEmojisUpdateHandler is triggered in GuildEmojisUpdate events
type GuildEmojisUpdateHandler = func(s Session, h *GuildEmojisUpdate)

// GuildIntegrationsUpdateHandler is triggered in GuildIntegrationsUpdate events
type GuildIntegrationsUpdateHandler = func(s Session, h *GuildIntegrationsUpdate)

// GuildMemberAddHandler is triggered in GuildMemberAdd events
type GuildMemberAddHandler = func(s Session, h *GuildMemberAdd)

// GuildMemberRemoveHandler is triggered in GuildMemberRemove events
type GuildMemberRemoveHandler = func(s Session, h *GuildMemberRemove)

// GuildMemberUpdateHandler is triggered in GuildMemberUpdate events
type GuildMemberUpdateHandler = func(s Session, h *GuildMemberUpdate)

// GuildMembersChunkHandler is triggered in GuildMembersChunk events
type GuildMembersChunkHandler = func(s Session, h *GuildMembersChunk)

// GuildRoleCreateHandler is triggered in GuildRoleCreate events
type GuildRoleCreateHandler = func(s Session, h *GuildRoleCreate)

// GuildRoleDeleteHandler is triggered in GuildRoleDelete events
type GuildRoleDeleteHandler = func(s Session, h *GuildRoleDelete)

// GuildRoleUpdateHandler is triggered in GuildRoleUpdate events
type GuildRoleUpdateHandler = func(s Session, h *GuildRoleUpdate)

// GuildUpdateHandler is triggered in GuildUpdate events
type GuildUpdateHandler = func(s Session, h *GuildUpdate)

// InviteDeleteHandler is triggered in InviteDelete events
type InviteDeleteHandler = func(s Session, h *InviteDelete)

// MessageCreateHandler is triggered in MessageCreate events
type MessageCreateHandler = func(s Session, h *MessageCreate)

// MessageDeleteHandler is triggered in MessageDelete events
type MessageDeleteHandler = func(s Session, h *MessageDelete)

// MessageDeleteBulkHandler is triggered in MessageDeleteBulk events
type MessageDeleteBulkHandler = func(s Session, h *MessageDeleteBulk)

// MessageReactionAddHandler is triggered in MessageReactionAdd events
type MessageReactionAddHandler = func(s Session, h *MessageReactionAdd)

// MessageReactionRemoveHandler is triggered in MessageReactionRemove events
type MessageReactionRemoveHandler = func(s Session, h *MessageReactionRemove)

// MessageReactionRemoveAllHandler is triggered in MessageReactionRemoveAll events
type MessageReactionRemoveAllHandler = func(s Session, h *MessageReactionRemoveAll)

// MessageUpdateHandler is triggered in MessageUpdate events
type MessageUpdateHandler = func(s Session, h *MessageUpdate)

// PresenceUpdateHandler is triggered in PresenceUpdate events
type PresenceUpdateHandler = func(s Session, h *PresenceUpdate)

// ReadyHandler is triggered in Ready events
type ReadyHandler = func(s Session, h *Ready)

// ResumedHandler is triggered in Resumed events
type ResumedHandler = func(s Session, h *Resumed)

// TypingStartHandler is triggered in TypingStart events
type TypingStartHandler = func(s Session, h *TypingStart)

// UserUpdateHandler is triggered in UserUpdate events
type UserUpdateHandler = func(s Session, h *UserUpdate)

// VoiceServerUpdateHandler is triggered in VoiceServerUpdate events
type VoiceServerUpdateHandler = func(s Session, h *VoiceServerUpdate)

// VoiceStateUpdateHandler is triggered in VoiceStateUpdate events
type VoiceStateUpdateHandler = func(s Session, h *VoiceStateUpdate)

// WebhooksUpdateHandler is triggered in WebhooksUpdate events
type WebhooksUpdateHandler = func(s Session, h *WebhooksUpdate)

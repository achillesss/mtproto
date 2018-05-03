package mtproto

import "fmt"

func (m *MTProto) ChannelsInviteToChannel(channel TL, users []TL) (*TL, error) {
	tl, err := m.InvokeSync(TL_channels_inviteToChannel{
		Channel: channel,
		Users:   users,
	})

	fmt.Printf("tl: %s\n", (*tl).encode())
	return tl, err
}

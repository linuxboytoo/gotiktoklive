package tests

import (
	"testing"

	"github.com/linuxboytoo/gotiktoklive"
)

func TestUserInfo(t *testing.T) {
	tiktok := gotiktoklive.NewTikTok()

	info, err := tiktok.GetUserInfo(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Got user info: %+v", info)
}

func TestRoomInfo(t *testing.T) {
	tiktok := gotiktoklive.NewTikTok()

	info, err := tiktok.GetRoomInfo(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Got user info: %+v", info)
}

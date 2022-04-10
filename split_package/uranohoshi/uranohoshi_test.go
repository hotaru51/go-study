package uranohoshi

import (
    "testing"
)

func TestChikaVoiceActor(t *testing.T) {
    expect := "Anju Inami"
    actual := ChikaVoiceActor()

    if expect != actual {
        t.Errorf("%s != %s", expect, actual)
    }
}

func TestYouVoiceActor(t *testing.T) {
    expect := "Shuka Saito"
    actual := YouVoiceActor()

    if expect != actual {
        t.Errorf("%s != %s", expect, actual)
    }
}

func TestRikoVoiceActor(t *testing.T) {
    expect := "Rikako Aida"
    actual := RikoVoiceActor()

    if expect != actual {
        t.Errorf("%s != %s", expect, actual)
    }
}

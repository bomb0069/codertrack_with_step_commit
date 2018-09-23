package codetrack

import (
	"testing"
)

func Test_like_action_will_get_score_from_count(t *testing.T) {
	expected := 4

	socialData := NewSocialData()

	socialData.action("BartoszMilewski", "jay_soo", "like", 4)

	actual := socialData.score("jay_soo")

	if expected != actual {
		t.Errorf("%d != %d", expected, actual)
	}
}

func Test_reply_action_will_get_double_score(t *testing.T) {
	expected := 27

	socialData := NewSocialData()

	socialData.action("BartoszMilewski", "jay_soo", "retweet", 9)

	actual := socialData.score("jay_soo")

	if expected != actual {
		t.Errorf("%d != %d", expected, actual)
	}
}

func Test_2_action_score_must_added(t *testing.T) {
	expected := 31
	socialData := NewSocialData()

	socialData.action("BartoszMilewski", "jay_soo", "like", 4)
	socialData.action("BartoszMilewski", "jay_soo", "retweet", 9)

	actual := socialData.score("jay_soo")

	if expected != actual {
		t.Errorf("%d != %d", expected, actual)
	}
}

func Test_string_action(t *testing.T) {
	expected := 4
	socialData := NewSocialData()

	socialData.actionString("BartoszMilewski jay_soo like 4")

	actual := socialData.score("jay_soo")

	if expected != actual {
		t.Errorf("%d != %d", expected, actual)
	}
}

func Test_string_action_more_than_one(t *testing.T) {
	expected := 31
	socialData := NewSocialData()

	socialData.actionString("BartoszMilewski jay_soo like 4")
	socialData.actionString("BartoszMilewski jay_soo retweet 9")

	actual := socialData.score("jay_soo")

	if expected != actual {
		t.Errorf("%d != %d", expected, actual)
	}
}

func Test_read_file(t *testing.T) {
	expected := 31
	socialData := NewSocialData()

	socialData.read("social_data.txt")

	actual := socialData.score("jay_soo")

	if expected != actual {
		t.Errorf("%d != %d", expected, actual)
	}
}

func Test_read_file_and_get_top(t *testing.T) {
	expected := "jay_soo"
	socialData := NewSocialData()

	socialData.read("social_data.txt")

	actual := socialData.topInfluencer(1)

	if expected != actual {
		t.Errorf("%s != %s", expected, actual)
	}
}

func Test_string_many_action_and_get_top(t *testing.T) {
	expected := "jay_soo"
	socialData := NewSocialData()

	socialData.actionString("BartoszMilewski bomb0069 like 5")
	socialData.actionString("BartoszMilewski jay_soo like 4")
	socialData.actionString("BartoszMilewski jay_soo retweet 9")

	actual := socialData.topInfluencer(1)

	if expected != actual {
		t.Errorf("%s != %s", expected, actual)
	}
}

func Test_string_many_action_and_get_top_2(t *testing.T) {
	expected := "martinfowler\nBartoszMilewski"
	socialData := NewSocialData()

	socialData.actionString("BartoszMilewski jay_soo like 4")
	socialData.actionString("BartoszMilewski jay_soo retweet 9")
	socialData.actionString("BartoszMilewski KentBeck reply 20")
	socialData.actionString("KentBeck BartoszMilewski reply 11")
	socialData.actionString("KentBeck BartoszMilewski retweet 22")
	socialData.actionString("BartoszMilewski martinfowler like 45")
	socialData.actionString("BartoszMilewski martinfowler reply 30")

	actual := socialData.topInfluencer(2)

	if expected != actual {
		t.Errorf("%s != %s", expected, actual)
	}
}

func Test_get_answer(t *testing.T) {
	expected := "dan_abramov\nmbrandonw\nmath3ma\nyouyuxi\nlambda_conf"
	socialData := NewSocialData()

	socialData.read("social_input_data.txt")

	actual := socialData.topInfluencer(5)

	if expected != actual {
		t.Errorf("%s != %s", expected, actual)
	}
}

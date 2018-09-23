package codetrack

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// SocialData is struct that contains all of influencer
type SocialData struct {
	scores map[string]Influencer
}

// Influencer is struct that contains name and score
type Influencer struct {
	name  string
	score int
}

var scoresRate = map[string]int{
	"like":    1,
	"reply":   2,
	"retweet": 3,
}

type ByScore []Influencer

func (a ByScore) Len() int { return len(a) }

func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].score < a[j].score }

// NewSocialData is function for create new influencer
func NewSocialData() SocialData {
	inf := SocialData{
		scores: make(map[string]Influencer),
	}
	return inf
}

func (inf *SocialData) actionString(action string) {
	actionArray := strings.Split(action, " ")
	actor := actionArray[0]
	actee := actionArray[1]
	action = actionArray[2]
	count, _ := strconv.Atoi(actionArray[3])
	inf.action(actor, actee, action, count)
}

func (inf *SocialData) action(actor string, actee string, action string, count int) {
	thisScore := count * scoresRate[action]
	influencer, ok := inf.scores[actee]
	if !ok {
		influencer = Influencer{name: actee, score: 0}
	}
	influencer.score = influencer.score + thisScore
	inf.scores[actee] = influencer
}

func (inf *SocialData) score(actee string) int {
	return inf.scores[actee].score
}

func (inf *SocialData) read(fileName string) {
	data, _ := ioutil.ReadFile(fileName)
	splittedLine := strings.Split(string(data), "\r\n")
	for _, line := range splittedLine {
		inf.actionString(line)
	}
}
func (inf *SocialData) topInfluencer(of int) string {
	returnString := ""
	influencers := make([]Influencer, 0)
	for _, influencer := range inf.scores {
		influencers = append(influencers, influencer)
	}
	sort.Sort(sort.Reverse(ByScore(influencers)))
	lastInfluencer := of - 1
	influencer := Influencer{}
	for index := 0; index < of; index++ {
		influencer = influencers[index]
		returnString = returnString + influencer.name
		if index != lastInfluencer {
			returnString = returnString + "\n"
		}
	}

	return returnString

}

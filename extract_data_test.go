package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_getAgentScore100(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(100)
	agentPrev := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)

	act_score := agentPrev.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreStarFive(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(98)
	agentPrev := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 5, Score: 100}
	getAgentScore(agentPrev, agentMap)

	act_score := agentPrev.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreSolicited(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(103)
	agentPrev := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: true,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)

	act_score := agentPrev.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreNumOfWords(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(99.5)
	agentPrev := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 105, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)

	act_score := agentPrev.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreMulti(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(100)
	timePrev := time.Now().Add(-(time.Hour * 2))
	agentPrev := &Agent{ReviewTime: timePrev, Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)
	agentMap[agentPrev.Name] = agentPrev

	agentNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDX", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentNext, agentMap)

	act_score := agentNext.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreTag(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(70)
	timePrev := time.Now().Add(-(time.Hour * 2))
	agentPrev := &Agent{ReviewTime: timePrev, Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)
	agentMap[agentPrev.Name] = agentPrev

	agentNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentNext, agentMap)

	act_score := agentNext.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreStarsAvg(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(92)
	timePrev := time.Now().Add(-(time.Hour * 2))
	agentPrev := &Agent{ReviewTime: timePrev, Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)
	agentMap[agentPrev.Name] = agentPrev

	agentNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDX", NumOfWords: 75, Stars: 1, Score: 100}
	getAgentScore(agentNext, agentMap)

	act_score := agentNext.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreMin(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(60)
	timePrev := time.Now().Add(-(time.Second * 2))
	agentPrev := &Agent{ReviewTime: timePrev, Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)
	agentMap[agentPrev.Name] = agentPrev

	agentNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDX", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentNext, agentMap)

	act_score := agentNext.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getAgentScoreHour(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_score := float32(80)
	timePrev := time.Now().Add(-(time.Minute * 2))
	agentPrev := &Agent{ReviewTime: timePrev, Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentPrev, agentMap)
	agentMap[agentPrev.Name] = agentPrev

	agentNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDX", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentNext, agentMap)

	act_score := agentNext.Score
	t.Logf("Agent %s has score %g\n", agentPrev.Name, act_score)
	if act_score != exp_score {
		t.Errorf("Failed with expected score:%g and actual score:%g\n", exp_score, act_score)
	}
}

func Test_getScoreReview(t *testing.T) {
	exp_review := fmt.Sprint("Info: Steve has a trusted review score of 100")
	agentPrev := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}
	act_review := getScoreReview(agentPrev)

	t.Logf("Agent %s has score review:\n %s\n", agentPrev.Name, act_review)
	if act_review != exp_review {
		t.Errorf("Failed with expected review:\n%s\n and actual review:\n%s\n", exp_review, act_review)
	}
}

func Test_getScoreReviewWarn(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_review := fmt.Sprint("Warning: Steve has a trusted review score of 60")
	agentPrev := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}

	getAgentScore(agentPrev, agentMap)
	agentMap[agentPrev.Name] = agentPrev

	agentNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDX", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentNext, agentMap)

	act_review := getScoreReview(agentNext)

	t.Logf("Agent %s has score review:\n %s\n", agentPrev.Name, act_review)
	if act_review != exp_review {
		t.Errorf("Failed with expected review:\n%s\n and actual review:\n%s\n", exp_review, act_review)
	}
}

func Test_getScoreReviewDeact(t *testing.T) {
	agentMap := make(map[string]*Agent)
	exp_review := fmt.Sprint("Alert: Steve has been de-activated due to a low trusted review score")
	agentPrev := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDF", NumOfWords: 75, Stars: 4, Score: 100}

	getAgentScore(agentPrev, agentMap)
	agentMap[agentPrev.Name] = agentPrev

	agentNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDX", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentNext, agentMap)
	agentMap[agentNext.Name] = agentNext

	agentNextNext := &Agent{ReviewTime: time.Now(), Name: "Steve", IsSolicited: false,
		Tag: "IF1-DDX", NumOfWords: 75, Stars: 4, Score: 100}
	getAgentScore(agentNextNext, agentMap)

	act_review := getScoreReview(agentNextNext)

	t.Logf("Agent %s has score review:\n %s\n", agentPrev.Name, act_review)
	if act_review != exp_review {
		t.Errorf("Failed with expected review:\n%s\n and actual review:\n%s\n", exp_review, act_review)
	}
}

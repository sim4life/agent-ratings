package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func readLine(filePath string) (func() (string, error), *bufio.Scanner, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineReader := func() (string, error) {
		scanner.Scan()
		return scanner.Text(), scanner.Err()
	}
	return lineReader, scanner, nil
}

func processFromCSVFile(filePath string) {
	agentMap := make(map[string]*Agent)
	file, err := os.Open(filePath)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scan := scanner.Text()
		fmt.Println(scan)
		agent, err := parseToStruct(strings.Split(scan, ","))
		if err != nil {
			fmt.Println("Could not read review summary data")
		} else {
			getAgentScore(agent, agentMap)
			agentMap[agent.Name] = agent

			fmt.Println(getScoreReview(agent))
		}
	}
}

func parseToStruct(rawData []string) (*Agent, error) {
	if len(rawData) < 5 {
		return nil, errors.New("Not enough fields")
	}
	agent := &Agent{Score: 100, NumOfRecords: 1}
	revTimeStr, name, isSolStr, tag, numOfWordsStr, starsStr :=
		rawData[0], strings.TrimSpace(rawData[1]), strings.TrimSpace(rawData[2]),
		strings.TrimSpace(rawData[3]), strings.TrimSpace(rawData[4]),
		strings.TrimSpace(rawData[5])
	timeValidator := &TimeValidator{revTimeStr}
	nameValidator := &NameValidator{name}
	isSolicitedValidator := &IsSolicitedValidator{isSolStr}
	tagValidator := &TagValidator{tag}
	numOfWordsValidator := &NumOfWordsValidator{numOfWordsStr}
	starsValidator := &StarsValidator{starsStr}
	if validateData(timeValidator) && validateData(nameValidator) &&
		validateData(isSolicitedValidator) && validateData(tagValidator) &&
		validateData(numOfWordsValidator) && validateData(starsValidator) {
		revTime, _ := parseReviewTime(revTimeStr)
		isSolicited := parseIsSolicited(isSolStr)
		numOfWords := parseNumOfWords(numOfWordsStr)
		stars := float32(len(starsStr))
		agent.ReviewTime, agent.Name, agent.IsSolicited, agent.Tag, agent.NumOfWords, agent.Stars =
			revTime, name, isSolicited, tag, numOfWords, stars
	} else {
		return nil, errors.New("Parse error")
	}
	return agent, nil
}

func getScoreReview(agent *Agent) string {
	var agentReview string
	if agent.Score < 50 {
		agentReview = fmt.Sprintf("Alert: %s has been de-activated due to a low trusted review score", agent.Name)
	} else if agent.Score < 70 {
		agentReview = fmt.Sprintf("Warning: %s has a trusted review score of %g", agent.Name, agent.Score)
	} else {
		agentReview = fmt.Sprintf("Info: %s has a trusted review score of %g", agent.Name, agent.Score)
	}
	return agentReview
}

func getAgentScore(agent *Agent, agentMap map[string]*Agent) {
	val, ok := agentMap[agent.Name]
	if ok {
		agent.Score = val.Score
		agent.NumOfRecords = val.NumOfRecords + 1
		if agent.Tag == val.Tag {
			agent.Score = val.Score - (0.3 * val.Score)
		}
		if (((float32(val.NumOfRecords) * val.Stars) + agent.Stars) / float32(agent.NumOfRecords)) < 3.5 {
			agent.Score = val.Score - ((0.02 * 4) * val.Score)
		}
		if agent.ReviewTime.Sub(val.ReviewTime) < time.Minute {
			agent.Score = val.Score - (0.4 * val.Score)
		} else if agent.ReviewTime.Sub(val.ReviewTime) < time.Hour {
			agent.Score = val.Score - (0.2 * val.Score)
		}
	}
	if agent.NumOfWords > 100 {
		agent.Score -= (0.005 * agent.Score)
	}
	if agent.Stars == 5 {
		agent.Score -= (0.02 * agent.Score)
	}
	if agent.IsSolicited {
		agent.Score += (0.03 * agent.Score)
	}

}

func parseReviewTime(revTimeStr string) (time.Time, error) {
	revTime, err := time.Parse(TimeFormat1, revTimeStr)
	if err != nil {
		revTime, err = time.Parse(TimeFormat2, revTimeStr)
		if err != nil {
			revTime, err = time.Parse(TimeFormat3, revTimeStr)
			if err != nil {
				revTime, err = time.Parse(TimeFormat4, revTimeStr)
			}
		}
	}
	if err == nil {
		addYear := time.Now().Year()
		revTime = revTime.AddDate(addYear, 0, 0)
	}
	return revTime, err
}

func parseIsSolicited(isSolicitedStr string) bool {
	var isSolicited bool
	if isSolicitedStr == "solicited" {
		isSolicited = true
	} else if isSolicitedStr == "unsolicited" {
		isSolicited = false
	}
	return isSolicited
}

func parseNumOfWords(numOfWordsStr string) int {
	re := regexp.MustCompile("[^0-9]")
	numStr := re.ReplaceAllString(numOfWordsStr, "")
	num, _ := strconv.ParseInt(numStr, 10, 32)
	return int(num)
}

func checkReadErr(e error) {
	if e == io.EOF {
		log.Fatal("No data found in file")
	}
	if e != nil {
		log.Fatalf("Error reading file: %s", e)
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatalf("Error: %s", e)
	}
}

package main

import (
	"regexp"
	"time"
)

const (
	TimeFormat1 = "_2th January 15:04"
	TimeFormat2 = "_2st January 15:04"
	TimeFormat3 = "_2nd January 15:04"
	TimeFormat4 = "_2rd January 15:04"
)

/* Match for tag, num_of_words and stars regexp */
var TagRegex = regexp.MustCompile(`^([[:upper:]]{2}\d)-([[:upper:]]{3})$`)
var NumOfWordsRegex = regexp.MustCompile(`^([\d]+)\swords$`)
var StarsRegex = regexp.MustCompile(`^([*]{1,5})$`)

type Validator interface {
	Validate() bool
}

type TimeValidator struct {
	time string
}

func (tmv *TimeValidator) Validate() bool {
	_, err := time.Parse(TimeFormat1, tmv.time)
	if err != nil {
		_, err = time.Parse(TimeFormat2, tmv.time)
		if err != nil {
			_, err = time.Parse(TimeFormat3, tmv.time)
			if err != nil {
				_, err = time.Parse(TimeFormat4, tmv.time)
			}
		}
	}
	return err == nil
}

type NameValidator struct {
	name string
}

func (nv *NameValidator) Validate() bool {
	return len(nv.name) > 0
}

type IsSolicitedValidator struct {
	isSolicited string
}

func (isv *IsSolicitedValidator) Validate() bool {
	return isv.isSolicited == "solicited" || isv.isSolicited == "unsolicited"
}

type TagValidator struct {
	tag string
}

func (tgv *TagValidator) Validate() bool {
	return TagRegex.MatchString(tgv.tag)
}

type NumOfWordsValidator struct {
	numOfWords string
}

func (nwv *NumOfWordsValidator) Validate() bool {
	return NumOfWordsRegex.MatchString(nwv.numOfWords)
}

type StarsValidator struct {
	stars string
}

func (sv *StarsValidator) Validate() bool {
	return StarsRegex.MatchString(sv.stars)
}

func validateData(validator Validator) bool {
	return validator.Validate()
}

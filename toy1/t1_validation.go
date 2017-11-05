package toy1

import (
	"net/http"
	"regexp"
	"strconv"
)

func validateString(r *http.Request) bool {
	if len(r.Form["username"][0]) == 0 {
		return false
	}
	return true
}

func validateInt(r *http.Request) bool {
	strage := r.Form.Get("age")
	intage, err := strconv.Atoi(strage)
	if err != nil {
		return false
	}
	return intage < 100 && intage > 0
}

func validateByRegexp(r *http.Request) bool {
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		return false
	}
	return true
}

func validateHan(r *http.Request) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
		return false
	}
	return true
}

func validateEnglish(r *http.Request) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("realname")); !m {
		return false
	}
	return true
}

func validateSelect(r *http.Request) bool {
	selectSlice := []string{"apple", "pear", "banana"}
	fruit := r.Form.Get("fruit")
	for _, item := range selectSlice {
		if item == fruit {
			return true
		}
	}
	return false
}

func validateRadio(r *http.Request) bool {
	slice := []int{1, 2, 3}
	gender, err := strconv.Atoi(r.Form.Get("gender"))
	if err != nil {
		return false
	}
	for _, v := range slice {
		if gender == v {
			return true
		}
	}
	return false
}

func inSlice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func sliceDiff(s1, s2 []interface{}) (diffSlice []interface{}) {
	for _, v := range s1 {
		if !inSlice(v, s2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

func validateCheckbox(r *http.Request) bool {
	rightAnswers := []string{"a", "b"}
	answers := r.Form["answers"]

	rightAnswersInt := make([]interface{}, len(rightAnswers))
	answersInt := make([]interface{}, len(answers))
	for i, v := range rightAnswers {
		rightAnswersInt[i] = interface{}(v)
	}
	for i, v := range answers {
		answersInt[i] = interface{}(v)
	}

	a := sliceDiff(rightAnswersInt, answersInt)
	if a == nil {
		return true
	}
	return false
}

package rules

type Rule func(message string) string

func GetAllRules() []Rule {
	return []Rule{
		CheckLowercase,
		CheckEnglish,
	}
}

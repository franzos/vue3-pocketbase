package main

import (
    "log"
    "strings"
)

func spamKeywordScore(text string) int {
    spamKeywords := []string{"viagra", "drugs", "cheap"} // Add more spam keywords as needed
    score := 0

    for _, keyword := range spamKeywords {
        count := strings.Count(strings.ToLower(text), keyword)
        score += count * 10 // Increase the score based on the number of occurrences of the keyword
    }

    return score
}

func wordRepetitionScore(text string) int {
    words := strings.Fields(strings.ToLower(text))
    wordCounts := make(map[string]int)
    maxCount := 0

    for _, word := range words {
        wordCounts[word]++
        if wordCounts[word] > maxCount {
            maxCount = wordCounts[word]
        }
    }

    return maxCount * 5 // Increase the score based on the maximum count of a single word
}

func sameValueFieldsScore(data map[string]interface{}) int {
    valueCounts := make(map[string]int)

    for _, value := range data {
        stringValue, ok := value.(string)
        if ok {
            valueCounts[stringValue]++
        }
    }

    sameValueCount := 0
    for _, count := range valueCounts {
        if count > 1 {
            sameValueCount += count
        }
    }

    return sameValueCount * 10 // Increase the score based on the number of fields with the same value
}

func linkCountScore(text string) int {
    linkCount := 0
    linkCount += strings.Count(strings.ToLower(text), "http://")
    linkCount += strings.Count(strings.ToLower(text), "https://")
    linkCount += strings.Count(strings.ToLower(text), "www.")

    if linkCount > 3 {
        return (linkCount - 3) * 10 // Increase the score for each link above 3
    }
    return 0
}

func shortLinkScore(text string) int {
    shortLinkDomains := []string{"bit.ly", "tinyurl.com"} // Add more link shortener domains as needed
    score := 0

    for _, domain := range shortLinkDomains {
        count := strings.Count(strings.ToLower(text), domain)
        score += count * 15 // Increase the score based on the number of occurrences of the short link domain
    }

    return score
}

func calculateSpamScore(data map[string]interface{}) int {
    totalScore := 0

    for _, value := range data {
        if stringValue, ok := value.(string); ok {
            totalScore += spamKeywordScore(stringValue)
            totalScore += wordRepetitionScore(stringValue)
            totalScore += linkCountScore(stringValue)
            totalScore += shortLinkScore(stringValue)
        }
    }

    totalScore += sameValueFieldsScore(data)

    // Clamp the score between 0 and 100
    if totalScore > 100 {
        totalScore = 100
    }

    log.Printf("Spam score: %d\n", totalScore)

    return totalScore
}
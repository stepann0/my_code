package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Allowed symbols
const (
	ascii_letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits        = "0123456789"
	specials      = `!#$%&()*+-/:<;=?>@[\]^_{}~`
)

type Config struct {
	length                               int
	with_digits, with_specials, readable bool
}

var WEAK_PWD = Config{
	length:        6,
	with_digits:   false,
	with_specials: false,
	readable:      true,
}

var MIDDLE_PWD = Config{
	length:        15,
	with_digits:   true,
	with_specials: false,
	readable:      true,
}

var STRONG_PWD = Config{
	length:        18,
	with_digits:   true,
	with_specials: true,
	readable:      true,
}

var SUPER_STRONG_PWD = Config{
	length:        30,
	with_digits:   true,
	with_specials: true,
	readable:      true,
}

func checkpwd(pwd string, config Config) bool {
	if len(pwd) != config.length || len(pwd) < 6 {
		return false
	}
	if config.with_digits && !strings.ContainsAny(pwd, digits) {
		return false
	}
	if config.with_specials && !strings.ContainsAny(pwd, specials) {
		return false
	}
	return true
}

func GeneratePassword(config Config) string {
	// Starts with just latin letters
	alphabet := ascii_letters
	if config.with_digits {
		alphabet += digits
	}
	if config.with_specials {
		alphabet += specials
	}
	if config.readable {
		unreadable := []string{"l", "I", "O", "б"}
		for _, s := range unreadable {
			alphabet = strings.ReplaceAll(alphabet, s, "")
		}
	}

	// Password generation
	var buff []byte
	l := len(alphabet)
	for i := 0; i < config.length; i++ {
		buff = append(buff, alphabet[rand.Intn(l)])
	}

	// Shuffle to increase security
	rand.Shuffle(config.length, func(i, j int) {
		buff[i], buff[j] = buff[j], buff[i]
	})

	password := string(buff)
	if !checkpwd(password, config) {
		return GeneratePassword(config)
	}
	return password
}

func RepeatPassword(conf Config, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(GeneratePassword(conf))
	}
}

func hash(pwd []byte) []byte {
	h := sha256.New()
	h.Write(pwd)
	return h.Sum(nil)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	RepeatPassword(STRONG_PWD, 5)
	RepeatPassword(MIDDLE_PWD, 5)
}

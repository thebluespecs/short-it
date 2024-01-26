package services

import (
	"short-it/internal/db"
	"short-it/internal/logger"
	"strconv"
	"time"
)

// takes the url and returns the encoded url
func Encode (url string, expiresAt time.Duration, dataStore db.DB) (string, error) {
    id, err := dataStore.Save(url, expiresAt)
    if err != nil {
        return "", err
    }
    return encode(uint64(id)), nil
}

// takes the code and returns the decoded url
func Decode (code string, dataStore db.DB) (string, error) {
    id := decode(code)
    logger.Info("decoded id: " + strconv.Itoa(int(id)))
    url, err := dataStore.Find(int(id))
    if err != nil {
        return "", err
    }
    return url, nil
}

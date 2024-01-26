package models

import (
    "time"
    "math/rand"
)

type ShortUrl struct {
    Id int `bson:"_id"`
    Url string `bson:"url"`
    Visits int `bson:"visits"`
    ExpireAt time.Duration `bson:"expireAt"`
    CreatedAt time.Time `bson:"createdAt"`
    UpdatedAt time.Time `bson:"updatedAt"`
}

func NewShortUrl(url string, expiry time.Duration) *ShortUrl {
    return &ShortUrl{
        Id: rand.Intn(100000),
        Url: url,
        Visits: 0,
        ExpireAt: expiry,
    }
}


package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type UserRequest struct {
	IP    string
	Count int
	Time  time.Time
}

var ipBucket = make(map[string]UserRequest)

func RateLimiter(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userReq := UserRequest{}
		start := time.Now()
		IP := r.RemoteAddr

		_, ok := ipBucket[IP]

		if !ok {
			userReq.IP = IP
			userReq.Count = 1
			userReq.Time = start
			ipBucket[IP] = userReq

		} else {
			userReq.IP = IP
			userReq.Count = ipBucket[IP].Count + 1
			userReq.Time = ipBucket[IP].Time
			ipBucket[IP] = userReq

		}

		log.Printf("count:%d, IP:%s, Time:%s", userReq.Count, userReq.IP, userReq.Time)
		end := time.Now().Unix()
		fmt.Printf("start:%d and end:%d\n", userReq.Time.Unix(), end)
		duration := end - userReq.Time.Unix()
		// log.Println("seconds:", duration)
		// to limit the rate per second, set the value of ur.count >= "Your desired value". Here Rate Limiter is 2 requests/second
		if userReq.Count >= 2 && r.RemoteAddr == userReq.IP {
			if duration < 1 {
				http.Error(w, "exceeded limit", http.StatusForbidden)
				fmt.Println("limit exceeded")
				// userReq.Count = 0
				return
			} else {
				// fmt.Println("delete key from ipBucket is hit")
				delete(ipBucket, IP)
			}

		}

		next.ServeHTTP(w, r)
	})

}

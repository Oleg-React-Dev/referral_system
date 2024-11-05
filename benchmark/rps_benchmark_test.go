package benchmark

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func BenchmarkRPS(b *testing.B) {
	dsn := os.Getenv("BENCH_DSN")
	query := os.Getenv("BENCH_QUERY")
	durationMs := 5000

	if dsn == "" || query == "" {
		b.Fatal("requires to setup BENCH_DSN and BENCH_QUERY in env")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		b.Fatalf("connection err: %v", err)
	}
	defer db.Close()

	measureRPS(b, db, query, time.Duration(durationMs)*time.Millisecond)
}

func measureRPS(b *testing.B, db *sql.DB, query string, duration time.Duration) {
	var wg sync.WaitGroup
	stopChan := make(chan struct{})
	var requestCount int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stopChan:
					return
				default:
					_, err := db.Exec(query)
					if err != nil {
						b.Logf("query err: %v", err)
					}
					requestCount++
				}
			}
		}()
	}

	time.Sleep(duration)
	close(stopChan)
	wg.Wait()

	rps := float64(requestCount) / duration.Seconds()
	fmt.Printf("RPS: %.2f\n", rps)
}

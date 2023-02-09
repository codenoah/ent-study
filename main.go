package main

import (
	"context"
	"ent-study/ent"
	"ent-study/ent/migrate"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=ent sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// 자동 마이그레이션 툴 실행
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.TODO()
	member, err := client.Debug().Member.Create().
		SetNickname("test").
		SetAccountID("test").
		SetAccountType("test").
		Save(ctx)

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		_, err = client.Debug().Post.Create().
			SetTitle("test").
			SetContent("test").
			SetMember(member).
			Save(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

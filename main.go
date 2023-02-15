package main

import (
	"context"
	"ent-study/application"
	"ent-study/ent"
	"ent-study/infrastructure/persistence"
	"ent-study/interfaces/handler/web"
	"entgo.io/ent/dialect/sql/schema"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// ent init
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=ent password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run migration
	if err = client.Schema.Create(context.Background(), schema.WithDropIndex(true), schema.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// repository init
	memberRepository := persistence.NewMemberRepository(client)

	// service init
	memberService := application.NewMemberApplication(memberRepository)

	// web run
	web.NewWebHandler().Run(memberService)
}

package main

import (
	"ent-study/application"
	"ent-study/ent"
	"ent-study/infrastructure/persistence"
	"ent-study/interfaces/handler/web"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// ent init
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=ent password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// repository init
	memberRepository := persistence.NewMemberRepository(client)

	// service init
	memberService := application.NewMemberApplication(memberRepository)

	// web run
	web.NewWebHandler().Run(memberService)
}

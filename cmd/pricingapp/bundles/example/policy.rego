package example.authz
import future.keywords.if

default allow := false
allow if input.open == data.example.bigSesame

default award_value := 0.0
award_value := 20.0 if input.open == data.example.hugeSesame
award_value := 10.0 if input.open == data.example.bigSesame
award_value := 5.0 if input.open == data.example.smallSesame

package siuyin.time

test_time {
  t := time.parse_ns("2006-01-02T15:04:05Z-0700","2020-12-25T12:34:56Z+0800")
  s := format(t)
  trace(sprintf("%v",[s]))
  s  == "2020-12-25T04:34:56"
}

test_replace_year {
  t := time.parse_ns("2006-01-02T15:04:05Z-0700","2020-12-25T12:34:56Z+0800")
  u := replace_year(1962,t)
  s := format(u)
  s == "1962-12-25T00:00:00"
}

test_year {
  t := time.parse_ns("2006-01-02T15:04:05Z-0700","2020-12-25T12:34:56Z+0800")
  y := year(t)
  y == 2020
}

test_age {
  t := time.parse_ns("2006-01-02","2020-06-23")
  bt := time.parse_ns("2006-01-02","1962-06-28")
  a := age(t,bt)
  trace(sprintf("age = %v",[a]))
  a == 57
}

test_double {
  d := double(3)
  d == 6
}

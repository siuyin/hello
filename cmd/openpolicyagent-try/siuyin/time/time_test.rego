package siuyin.time

test_time {
  t := time.parse_ns("2006-01-02T15:04:05Z-0700","2020-12-25T12:34:56Z+0800")
  s := format(t)
  trace(sprintf("%v",[s]))
  s  == "2020-12-25T04:34:56"
}

test_double {
  d := double(3)
  d == 6
}

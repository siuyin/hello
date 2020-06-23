package hdb.new_flat

test_age_eligible {
  age_eligible with input as {"applicants":[{"applicant":{"age":21,"name":"Alpha"}}]}
}

test_age_not_eligible {
  not age_eligible with input as {"applicants":[{"applicant":{"age":19,"name":"Beta"}}]}
}

# bd is the birth date in format 2006-01-02.
age(now,bd) = x {
  n := ymd(now)
  b := ymd(replace_year(n[0],bd))
  x := b[0]
}
ymd(d) = x {
  x := time.date(date(d))
}
date(d) = x {
 x := time.parse_ns("2006-01-02",d)
}
replace_year(y,d) = x {
  a := ymd(d)
  x := sprintf("%04d-%02d-%02d",[y,a[1],a[2]])
}
today(dummy) = x {
  d := time.date(time.now_ns())
  x := sprintf("%04d-%02d-%02d",[d[0],d[1],d[2]])
}

test_age {
  #a := age("2020-01-01","1962-06-28")
  a := age(today({}),"1962-06-28")
  trace(sprintf("%v",[a]))
  a == 2020
}

import data.sytime as st

test_format {
  t := time.parse_ns("2006-01-02T15:04:05Z-0700","2020-12-25T12:34:56Z+0800")
  s := st.format(t)
  trace(sprintf("s = %v",[s]))
  s == "2020-12-25T04:34:57"
}

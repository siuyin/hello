# Package time provides time related functions written in rego.
# Time in rego is represented as the nanoseconds elapsed since unix epoch (1970-01-01 UTC).
package siuyin.time


double(a) = x {
  x := a*2
}


# format takes t, the number of nanoseconds elapsed since unix epoch and returns x, a string representation of the time.
format(t) = x {
  x := sprintf("%04d-%02d-%02dT%02d:%02d:%02d",array.concat(time.date(t),time.clock(t)))
}

# replace_year replaces the year of time t with y and returns new time x.
replace_year(y,t) = x {
  ymd := time.date(t)
  x := time.parse_ns("2006-01-02",sprintf("%04d-%02d-%02d",[y,ymd[1],ymd[2]]))
}

# year extract the year from time t and return year x as an integer.
year(t) = x {
  ymd := time.date(t)
  x := ymd[0]
}

# age returns the age in years,x as an integer as at time t for a person with birth time bt.
age(t,bt) = x {
  cbt := replace_year(year(t),bt) # current year's birth time
  t >= cbt
  x := year(t) - year(bt)
}
age(t,bt) = x {
  cbt := replace_year(year(t),bt) # current year's birth time
  t < cbt
  x := year(t) - year(bt) - 1
}

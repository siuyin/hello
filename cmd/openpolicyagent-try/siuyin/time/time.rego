package siuyin.time

double(a) = x {
  x := a*2
}


# format takes t, the number of nanoseconds elapsed since unix epoch and returns x, a string representation of the time.
format(t) = x {
  x := sprintf("%04d-%02d-%02dT%02d:%02d:%02d",array.concat(time.date(t),time.clock(t)))
}

package hdb.new_flat

default age_eligible = false
age_eligible {
  input.age >= data.minimum_age
}

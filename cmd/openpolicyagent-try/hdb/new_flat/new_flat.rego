package hdb.new_flat

default age_eligible = false
age_eligible {
  input.applicants[_].applicant.age >= data.minimum_age
}

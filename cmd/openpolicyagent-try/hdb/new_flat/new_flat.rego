package hdb.new_flat

# parse_date parses date string d and returns time x represented in nanoseconds elapsed since unix epoch.
parse_date(d) = x {
  x := time.parse_ns("2006-01-02",d)
}

import data.siuyin.time as st

default age_eligible = false
age_eligible {
  bd = parse_date(input.applicants[_].applicant.date_of_birth)
  st.age(parse_date(input.date_of_application),bd) >= data.minimum_age
}

default at_least_one_singapore_citizen = false
at_least_one_singapore_citizen {
  input.applicants[_].applicant.resident_status_in_singapore == "citizen"
}

default two_singapore_citizens_or_a_citizen_and_a_pr = false
two_singapore_citizens_or_a_citizen_and_a_pr {
  some i,j
  input.applicants[i].applicant.resident_status_in_singapore == "citizen"
  input.applicants[j].applicant.resident_status_in_singapore == "citizen"
  i != j
}
two_singapore_citizens_or_a_citizen_and_a_pr {
  some i,j
  input.applicants[i].applicant.resident_status_in_singapore == "citizen"
  input.applicants[j].applicant.resident_status_in_singapore == "pr"
  i != j
}

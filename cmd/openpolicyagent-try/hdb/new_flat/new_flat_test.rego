package hdb.new_flat

test_age_eligible {
  age_eligible with input as {"applicants":[{"applicant":{"age":21,"name":"Alpha"}}]}
}

test_age_not_eligible {
  not age_eligible with input as {"applicants":[{"applicant":{"age":19,"name":"Beta"}}]}
}


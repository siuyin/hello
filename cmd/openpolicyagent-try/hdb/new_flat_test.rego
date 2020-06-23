package hdb.new_flat

test_age_eligible {
  age_eligible with input as {"age":21}
}

test_age_not_eligible {
  not age_eligible with input as {"age":20}
}


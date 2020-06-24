package hdb.new_flat

test_age_eligible {
  age_eligible with input as { "date_of_application":"2020-06-23", "applicants":[ {"applicant":{"date_of_birth":"1999-06-23"}} ] }
  age_eligible with input as { "date_of_application":"2020-06-23", "applicants":[ {"applicant":{"date_of_birth":"1999-06-22"}} ] }

  not age_eligible with input as { "date_of_application":"2020-06-23", "applicants":[ {"applicant":{"date_of_birth":"1999-06-24"}} ] }
  not age_eligible with input as { "date_of_application":"2020-06-23", "applicants":[ {"applicant":{"date_of_birth":"2000-06-23"}} ] }
}

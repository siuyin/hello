package hdb.new_flat

test_age_eligible {
  age_eligible with input as
  {"date_of_application":"2020-06-23","applicants":[
    {"applicant":{"date_of_birth":"1999-06-23"}},
  ]}

  age_eligible with input as
  {"date_of_application":"2020-06-23","applicants":[
    {"applicant":{"date_of_birth":"1999-06-22"}}
  ]}

  age_eligible with input as
  {"date_of_application":"2020-06-23","applicants":[
    {"applicant":{"date_of_birth":"1999-06-22"}},
    {"applicant":{"date_of_birth":"2002-06-22"}}
  ]}

  not age_eligible with input as
  {"date_of_application":"2020-06-23","applicants":[
    {"applicant":{"date_of_birth":"1999-06-24"}}
  ]}

  not age_eligible with input as
  {"date_of_application":"2020-06-23","applicants":[
    {"applicant":{"date_of_birth":"2000-06-23"}},
    {"applicant":{"date_of_birth":"2000-06-23"}}
  ]}
}

test_at_least_one_singapore_citizen{
  at_least_one_singapore_citizen with input as 
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"citizen"}},
    {"applicant":{"resident_status_in_singapore":"other"}}
  ]}

  at_least_one_singapore_citizen with input as 
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"citizen"}},
    {"applicant":{"resident_status_in_singapore":"citizen"}}
  ]}

  not at_least_one_singapore_citizen with input as 
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"pr"}},
    {"applicant":{"resident_status_in_singapore":"other"}}
  ]}

  not at_least_one_singapore_citizen with input as 
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"other"}},
    {"applicant":{"resident_status_in_singapore":"other"}}
  ]}
}

test_two_singapore_citizens_or_a_citizen_and_a_pr {
  # two citizens
  two_singapore_citizens_or_a_citizen_and_a_pr with input as
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"citizen"}},
    {"applicant":{"resident_status_in_singapore":"citizen"}}
  ]}

  # citizen and pr
  two_singapore_citizens_or_a_citizen_and_a_pr with input as
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"citizen"}},
    {"applicant":{"resident_status_in_singapore":"pr"}}
  ]}

  # no citizen or pr
  not two_singapore_citizens_or_a_citizen_and_a_pr with input as
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"other"}},
    {"applicant":{"resident_status_in_singapore":"other"}}
  ]}

  # one pr 
  not two_singapore_citizens_or_a_citizen_and_a_pr with input as
  {"applicants":[
    {"applicant":{"resident_status_in_singapore":"pr"}},
    {"applicant":{"resident_status_in_singapore":"other"}}
  ]}
}

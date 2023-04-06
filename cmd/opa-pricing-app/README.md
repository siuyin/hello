# Policy based pricing example

## Requirements
Given I am a company manager having a store with goods to sell,
When I provide a quotation to a prospective customer
Then I should see the pricing on that quotation reflect my company's pricing policy.

## Simple implementation
1. Have a table of SKUs (Stock Keeping Units) to track store inventory and base pricing.
1. Have a table of customer types (eg. prospect, high-value prospect, existing customer, gold-level customer etc.)
1. Have a way of classifying the prospect given the above customer types.
1. Have a pricing policy covering SKU pricing and customer type.
1. Compute the quotation (list of SKUs, qty and final price).

[![](https://mermaid.ink/img/pako:eNqVkk1rg0AQhv_KsJdW0CCFXDx4CDk0pKUt0puXyTomS3TX7EeDhPz37kaTSj8OFQR3nfd9Z57dE-OqIpYxQwdHktNS4FZjW0rkVmlYyY_nrb6uXlUjeD_ZKLAhU0pN3ILebvA-jR_m8ziN01k6j0oJ_hkskmi5yOCRNAH61-4IivW7AZQVbNAQdFrw4BUky0WSRIMug5d1DKu7poGGLPTKwV6qI4g6eIxuwhhHZjaIb00m0ZtTlorejMHCx0F3-Q2uq9DSKLnW-dSb-n_BJKsphhHCFMMFVZLnv_aklemCONBA0OEsjIXaEz6E8m9d5nlg-aTU3nWB4t8Qp3EozZH0D6uvAq7azlmCo7A7ENKSltiMvL5zyvPLPOMQAc5WDRhYzFrSLYrK36pT0JXMA2upZJn_rKhG19iSlfLsS9FZVfSSs8xqRzEbjmW8hCyrsTF0_gS_lOK3?type=png)](https://mermaid.live/edit#pako:eNqVkk1rg0AQhv_KsJdW0CCFXDx4CDk0pKUt0puXyTomS3TX7EeDhPz37kaTSj8OFQR3nfd9Z57dE-OqIpYxQwdHktNS4FZjW0rkVmlYyY_nrb6uXlUjeD_ZKLAhU0pN3ILebvA-jR_m8ziN01k6j0oJ_hkskmi5yOCRNAH61-4IivW7AZQVbNAQdFrw4BUky0WSRIMug5d1DKu7poGGLPTKwV6qI4g6eIxuwhhHZjaIb00m0ZtTlorejMHCx0F3-Q2uq9DSKLnW-dSb-n_BJKsphhHCFMMFVZLnv_aklemCONBA0OEsjIXaEz6E8m9d5nlg-aTU3nWB4t8Qp3EozZH0D6uvAq7azlmCo7A7ENKSltiMvL5zyvPLPOMQAc5WDRhYzFrSLYrK36pT0JXMA2upZJn_rKhG19iSlfLsS9FZVfSSs8xqRzEbjmW8hCyrsTF0_gS_lOK3)

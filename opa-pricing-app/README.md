# Policy based pricing example

## Requirements
Given I am a company manager having a store with goods to sell,  
When I provide a quotation to a prospective customer  
Then I should see the pricing on that quotation reflect my company's pricing policy.

## Implementation
1. Have a table of SKUs (Stock Keeping Units) to track store inventory and base pricing.
1. Have a table of customer types (eg. prospect, high-value prospect, existing customer, gold-level customer etc.)
1. Have a way of classifying the prospect given the above customer types.
1. Have a pricing policy covering SKU pricing and customer type.
1. Compute the quotation (list of SKUs, qty and final price).
[![](https://mermaid.ink/img/pako:eNqVkk1rg0AQhv_KsJc2oEEKuXjwEDw0pKUt0puXiU7MEt01-9EgIf-9u2oS-3WosOIu876-88yeWCFLYjHTdLAkCko5VgqbXGBhpIKV-Hiu1GX3KmtedJODDGvSuVBUGFDVBu-j4GGxCKIgmkeLWS7APS0qSJdg2xKNL_Zng204S5cxPJIiQLfMjiBbv2tAUcIGNUGreHGRpMswnA26GF7WAazu6hpqMtBJC3shj8C33mN041pb0vNB7B2H8F9zXBsKZ29WGso6PQbiLga0U8lodalzaa7q_wciUeaif93YjeSm7Hq-YZL8Gk5J3Xqxbw5B-QFqA1s3loMv_xY3STzsJyn3tvWY_6Y8_R0KfST1w-pWUMimtYbgyM0OuDCkBNYjuO_AkqTvZ2zCU6rkvMfAAtaQapCX7iqevC5njlxDOYvdZ0lbtLXJWS7OrhStkVknChYbZSlgw3zGm8viLdaazp8C-POk?type=png)](https://mermaid.live/edit#pako:eNqVkk1rg0AQhv_KsJc2oEEKuXjwEDw0pKUt0puXiU7MEt01-9EgIf-9u2oS-3WosOIu876-88yeWCFLYjHTdLAkCko5VgqbXGBhpIKV-Hiu1GX3KmtedJODDGvSuVBUGFDVBu-j4GGxCKIgmkeLWS7APS0qSJdg2xKNL_Zng204S5cxPJIiQLfMjiBbv2tAUcIGNUGreHGRpMswnA26GF7WAazu6hpqMtBJC3shj8C33mN041pb0vNB7B2H8F9zXBsKZ29WGso6PQbiLga0U8lodalzaa7q_wciUeaif93YjeSm7Hq-YZL8Gk5J3Xqxbw5B-QFqA1s3loMv_xY3STzsJyn3tvWY_6Y8_R0KfST1w-pWUMimtYbgyM0OuDCkBNYjuO_AkqTvZ2zCU6rkvMfAAtaQapCX7iqevC5njlxDOYvdZ0lbtLXJWS7OrhStkVknChYbZSlgw3zGm8viLdaazp8C-POk)

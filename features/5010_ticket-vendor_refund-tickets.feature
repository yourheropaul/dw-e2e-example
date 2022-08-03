Feature:
      As a Ticket vendor
      I want to Refund and invalidate existing tickets
      So that I can refund ticket sales and handle failed payment settlements
      But I know that once an NFT has been transferred to the customer, it's out of my control

  Background:
    Given I have access to the vendor backend

  Scenario: Refunds for un-redeemed tickets
    Given There is an NFT ticket that has not been transferred to the customer
    When I mark the ticket as refunded
    Then The NFT should be deleted

  Scenario: Can't refund redeemed tickets
    Given There is an NFT ticket that has been transferred to the customer
    When I mark the ticket as refunded
    Then I should get an error telling me that the NFT can't be deleted

Feature:
      As a Ticket vendor
      I want to Sell tickets for my events
      So that I can enjoy the benefits of conventional and NFT ticketing
      But I will have to handle all financial transactions myself
      And I will have to build some technical infrastructure to handle notifications from the Garizon API

  Background:
    Given I have configured the Garizon API to notify my API about NFT events

  Scenario: Receive notifications about minted tickets
    When A new ticket NFT is minted
    Then My API should receive an "NFT minted" notification
    And The notification should include all details of the NFT

  Scenario: Receive notifications about refunded tickets
    When An NFT ticket is deleted before being transferred to the customer
    Then My API should receive an "NFT deleted" notification
    And The notification should include all details of the NFT

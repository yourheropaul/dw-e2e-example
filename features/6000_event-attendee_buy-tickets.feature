Feature:
      As an Event attendee
      I want to buy tickets for an event from an official vendor
      So that I can attend an event
      But I accept that, for NFT tickets, the NFT will not be transferred to me until I check into the event

  Scenario: Buy a ticket and see the NFT
    When I buy an NFT ticket from a vendor
    Then The NFT representing the ticket should be minted
    And I should be able to see the NFT in the block explorer


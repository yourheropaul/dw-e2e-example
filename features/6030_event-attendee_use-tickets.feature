Feature:
      As an Event attendee
      I want to Use my tickets to get into events
      So that I can enjoy the event

  Scenario: Get into event with a Loud and Live-issued ticket
    Given I have purchased an NFT ticket for an event
    When I check in at the event
    Then The ownership of my NFT ticket should be transferred to me
    And I should be able to see the NFT in my wallet

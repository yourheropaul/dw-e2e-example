Feature:
      As an Account holder
      I want to see the balances of my account
      So that I can make financial decisions

  Background:
    Given there's a network
    And I have an account with funds

  Scenario: See balance in block explorer
    When I view the block explorer for my account
    Then I should see the balance of my account in GAR
    And I should see the balance of my account in USD

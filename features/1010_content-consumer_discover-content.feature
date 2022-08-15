Feature:
  As a Content consumer
  I want to discover and read Daily Wire news articles
  So that I can keep up with general events and opion

  Background:
    Given there a several articles already posted

  Scenario: See top story
    Given an article has been promoted to the Top Story
    When I visit the home page
    Then I should see the Top Story article preview displayed prominently



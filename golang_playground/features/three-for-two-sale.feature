Feature: three for two stock clearance sale
  In order to clear the stock of old items
  As a store manager
  I want customer to get a promotion of three items for the price of two

  Scenario: buyer gets third sale item for free
    Given there is on sale bicycle "Kangaroo 3000" which costs 299 pounds
    When I add three "Kangaroo 3000" bicycles into the basket
    Then I should get third bicycle for free
    And the basket total should be 598 pounds

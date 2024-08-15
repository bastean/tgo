Feature: Update a user account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Username with update
    * I fill the Currency list with usd
    * I fill the Coins list with monero
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see Account created notification

  Scenario: Update a valid existing account
    Given I am on / page
    Then I click the Update tab
    * I fill the Username with update
    * I fill the Currency list with eur
    * I fill the Coins list with bitcoin
    * I click the Update button
    And I see Account updated notification

  Scenario: Update a valid non existing account
    Given I am on / page
    Then I click the Update tab
    * I fill the Username with unknown
    * I fill the Currency list with eur
    * I fill the Coins list with bitcoin
    * I click the Update button
    But I see unknown not found notification

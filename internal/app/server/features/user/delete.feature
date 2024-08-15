Feature: Delete a user account

  Scenario: Create a valid non existing account
    Given I am on / page
    Then I click the Create tab
    * I fill the Username with delete
    * I fill the Currency list with usd
    * I fill the Coins list with monero
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see Account created notification

 Scenario: Delete a valid existing User
    Given I am on / page
    Then I click the Delete tab
    * I fill the Username with delete
    * I click the Delete button
    And I see Account deleted notification

  Scenario: Delete a valid non existing User
    Given I am on / page
    Then I click the Delete tab
    * I fill the Username with unknown
    * I click the Delete button
    But I see unknown not found notification

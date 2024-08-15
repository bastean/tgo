Feature: Create a new User

  Scenario: Create a valid non existing User
    Given I am on / page
    Then I click the Create tab
    * I fill the Username with create
    * I fill the Currency list with usd
    * I fill the Coins list with monero
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see Account created notification

  Scenario: Create already existing User
    Given I am on / page
    Then I click the Create tab
    * I fill the Username with create
    * I fill the Currency list with usd
    * I fill the Coins list with monero
    * I check the I agree to the terms and conditions
    * I click the Create button
    But I see create already registered notification

Feature: Read a User

  Scenario: Create a valid non existing User
    Given I am on / page
    Then I click the Create tab
    * I fill the Username with read
    * I fill the Currency list with usd
    * I fill the Coins list with monero
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see Account created notification

  Scenario: Read a valid existing User
    Given I am on / page
    Then I click the Read tab
    * I fill the Username with read
    * I click the Read button
    And I see Found notification

  Scenario: Read a valid non existing User
    Given I am on / page
    Then I click the Read tab
    * I fill the Username with unknown
    * I click the Read button
    But I see unknown not found notification

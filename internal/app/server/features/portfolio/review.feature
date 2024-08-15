Feature: Review a User Portfolio

  Scenario: Create a valid non existing User
    Given I am on / page
    Then I click the Create tab
    * I fill the Username with portfolio
    * I fill the Currency list with usd
    * I fill the Coins list with monero
    * I check the I agree to the terms and conditions
    * I click the Create button
    And I see Account created notification

  Scenario: Review a valid existing User Portfolio
    Given I am on / page
    Then I click the Portfolio tab
    * I fill the Username with portfolio
    * I click the Review button
    And I see Result notification

  Scenario: Review a valid non existing User Portfolio
    Given I am on / page
    Then I click the Portfolio tab
    * I fill the Username with unknown
    * I click the Review button
    But I see unknown not found notification

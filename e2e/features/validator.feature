Feature: Joining network as a validator
  As a user
  I want to be able to join validators set

  Background:
    Given I have the following accounts:
      | account | password | tokens | funds | validator |
      | A       | test     | 20     | 10    | true      |
      | B       | test     | 10     | 10    | false     |

  Scenario: Start validator
    When I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    Then the deposit of A should be 15 mTokens

  Scenario: Stop mining
    Given I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    And the deposit of A should be 15 mTokens
    When I withdraw my node from validation
    Then there should be 5 mTokens available to me after 5 days

  Scenario: Re-Start mining
    Given I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    And the deposit of A should be 15 mTokens
    When I withdraw my node from validation
    Then there should be 5 mTokens available to me after 5 days
    And I start validator with 5 mTokens deposit
    And I wait for my node to be synced
    And the deposit of A should be 10 mTokens

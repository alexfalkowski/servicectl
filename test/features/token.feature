Feature: Token

  The ability to rotate keys.

  Scenario: Successfully rotate token key
    When we rotate the token key
    Then we should have a succesful rotation of token key
    And I should see a log entry of "token: successfully rotated key" in the file "reports/token.log"
